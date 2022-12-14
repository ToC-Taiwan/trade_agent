// Package tickprocess package tickprocess
package tickprocess

import (
	"time"

	"trade_agent/pkg/cache"
	"trade_agent/pkg/config"
	"trade_agent/pkg/dbagent"
	"trade_agent/pkg/eventbus"
	"trade_agent/pkg/log"
	"trade_agent/pkg/sinopacapi"
)

var (
	simTradeRealTimeTickChannel   = make(chan int)
	simTradeRealTimeBidAskChannel = make(chan int)
)

// InitTickProcess InitTickProcess
func InitTickProcess() {
	log.Get().Info("Initial TickProcess")

	go simTradeRealTimeTickCollector()
	go simTradeRealTimeBidAskCollector()

	// sub targets to sub mq history tick, kbar, realtime tick, bidask
	eventbus.Get().SubscribeTargets(targetsBusCallback)

	err := subHistroyTick()
	if err != nil {
		log.Get().Panic(err)
	}
	err = subHistroyKbar()
	if err != nil {
		log.Get().Panic(err)
	}
}

func targetsBusCallback(targetArr []*dbagent.Target) {
	tmp := []*dbagent.Target{}
	for _, v := range targetArr {
		if v.Subscribe {
			tmp = append(tmp, v)

			cache.GetCache().SetRealTimeTickChannel(v.Stock.Number, make(chan *dbagent.RealTimeTick))
			go realTimeTickProcessor(v.Stock.Number)

			cache.GetCache().SetRealTimeBidAskChannel(v.Stock.Number, make(chan *dbagent.RealTimeBidAsk))
			go realTimeBidAskProcessor(v.Stock.Number)
		}
	}

	err := subRealTimeTick()
	if err != nil {
		log.Get().Panic(err)
	}
	err = subRealTimeBidAsk()
	if err != nil {
		log.Get().Panic(err)
	}
	if len(tmp) != 0 {
		eventbus.Get().PublishSubscribeTargets(tmp)
	}
}

func realTimeTickProcessor(stockNum string) {
	analyzeConf := config.GetAnalyzeConfig()
	ch := cache.GetCache().GetRealTimeTickChannel(stockNum)

	var tickArr dbagent.RealTimeTickArr
	var lastPeriodEndTime time.Time
	var firstArrive bool
	for {
		tick := <-ch
		tickArr = append(tickArr, tick)
		if !firstArrive {
			firstArrive = true
			cache.GetCache().SetStockHistoryOpen(stockNum, tick.Close, cache.GetCache().GetTradeDay())
		}

		// save realtime tick close to cache
		cache.GetCache().SetRealTimeTickClose(stockNum, tick.Close)
		if getForwardRestCount(stockNum)+getReverseRestCount(stockNum) == 0 {
			if tick.PctChg < analyzeConf.CloseChangeRatioLow || tick.PctChg > analyzeConf.CloseChangeRatioHigh {
				continue
			}
		}

		if lastPeriodEndTime.IsZero() {
			lastPeriodEndTime = checkFirstTick(tick, analyzeConf.OpenCloseChangeRatioLow, analyzeConf.OpenCloseChangeRatioHigh)
			continue
		}

		lastPeriodArr := tickArr.GetLastNSecondArr(analyzeConf.TickAnalyzeMinPeriod)
		if tick.TickTime.Before(lastPeriodEndTime.Add(time.Duration(analyzeConf.TickAnalyzeMinPeriod) * time.Second)) {
			continue
		} else {
			lastPeriodEndTime = tick.TickTime
		}

		var action sinopacapi.OrderAction
		realTimeBalancePct, emerAction := getRealTimeBalancePct(stockNum, tick.Close)
		if realTimeBalancePct >= analyzeConf.MaxLoss {
			action = emerAction
		} else {
			if action = realTimeTickArrActionGenerator(tickArr, lastPeriodArr, analyzeConf); action == 0 {
				continue
			}
		}

		order := &sinopacapi.Order{
			StockNum:  stockNum,
			Price:     tick.Close,
			Action:    action,
			TradeTime: tick.TickTime,
		}

		// send order event
		eventbus.Get().PublishStockOrder(order)
	}
}

func realTimeBidAskProcessor(stockNum string) {
	ch := cache.GetCache().GetRealTimeBidAskChannel(stockNum)
	var bidAskArr []*dbagent.RealTimeBidAsk
	for {
		bidAsk := <-ch
		bidAskArr = append(bidAskArr, bidAsk)
		status := realTimeBidAskArrStatusGenerator(bidAsk, bidAskArr)
		cache.GetCache().SetRealTimeBidAskStatus(stockNum, status)
	}
}

func simTradeRealTimeTickCollector() {
	printMinute := time.Now().Minute()
	var count int
	for {
		simTrade := <-simTradeRealTimeTickChannel
		count += simTrade
		if time.Now().Minute() != printMinute {
			printMinute = time.Now().Minute()
			log.Get().WithFields(map[string]interface{}{
				"Count": count,
			}).Info("SimTradeTick")
		}
	}
}

func simTradeRealTimeBidAskCollector() {
	printMinute := time.Now().Minute()
	var count int
	for {
		simTrade := <-simTradeRealTimeBidAskChannel
		count += simTrade
		if time.Now().Minute() != printMinute {
			printMinute = time.Now().Minute()
			log.Get().WithFields(map[string]interface{}{
				"Count": count,
			}).Info("SimTradeBidAsk")
		}
	}
}

func getRealTimeBalancePct(stockNum string, close float64) (float64, sinopacapi.OrderAction) {
	historyOrderBuy := cache.GetCache().GetOrderBuy(stockNum)
	historyOrderSell := cache.GetCache().GetOrderSell(stockNum)
	restOrderCount := len(historyOrderBuy) - len(historyOrderSell)
	if restOrderCount != 0 {
		for i := len(historyOrderSell); i <= len(historyOrderSell)-1+restOrderCount; i++ {
			lastOrder := historyOrderBuy[i]
			buyCost := sinopacapi.GetStockBuyCost(lastOrder.Price, lastOrder.Quantity)
			sellCost := sinopacapi.GetStockSellCost(close, lastOrder.Quantity)
			if buyCost > sellCost {
				return 100 * float64(buyCost-sellCost) / float64(buyCost), sinopacapi.ActionSell
			}
		}
	}

	historyOrderSellFirst := cache.GetCache().GetOrderSellFirst(stockNum)
	historyOrderBuyLater := cache.GetCache().GetOrderBuyLater(stockNum)
	restOrderCount = len(historyOrderSellFirst) - len(historyOrderBuyLater)
	if restOrderCount != 0 {
		for i := len(historyOrderBuyLater); i <= len(historyOrderBuyLater)-1+restOrderCount; i++ {
			lastOrder := historyOrderSellFirst[i]
			sellFirstCost := sinopacapi.GetStockSellCost(lastOrder.Price, lastOrder.Quantity)
			buyLaterCost := sinopacapi.GetStockBuyCost(close, lastOrder.Quantity)
			if sellFirstCost < buyLaterCost {
				return 100 * float64(buyLaterCost-sellFirstCost) / float64(sellFirstCost), sinopacapi.ActionBuyLater
			}
		}
	}
	return 0, 0
}

func getForwardRestCount(stockNum string) int {
	historyOrderBuy := cache.GetCache().GetOrderBuy(stockNum)
	historyOrderSell := cache.GetCache().GetOrderSell(stockNum)
	return len(historyOrderBuy) - len(historyOrderSell)
}

func getReverseRestCount(stockNum string) int {
	historyOrderSellFirst := cache.GetCache().GetOrderSellFirst(stockNum)
	historyOrderBuyLater := cache.GetCache().GetOrderBuyLater(stockNum)
	return len(historyOrderSellFirst) - len(historyOrderBuyLater)
}

func checkFirstTick(tick *dbagent.RealTimeTick, low, high float64) time.Time {
	if tick.PctChg < low || tick.PctChg > high {
		target := cache.GetCache().GetTargetByStockNum(tick.Stock.Number)
		eventbus.Get().PublishUnSubscribeTargets(target)
	}
	return tick.TickTime
}
