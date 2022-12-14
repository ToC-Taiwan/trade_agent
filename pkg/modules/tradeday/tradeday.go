// Package tradeday package tradeday
package tradeday

import (
	"time"

	"trade_agent/global"
	"trade_agent/pkg/cache"
	"trade_agent/pkg/config"
	"trade_agent/pkg/dbagent"
	"trade_agent/pkg/log"
)

const (
	openTime time.Duration = 9 * time.Hour
)

// InitTradeDay InitTradeDay
func InitTradeDay() {
	log.Get().Info("Initial TradeDay")

	// save calendar to db and cache
	err := ImportCalendar()
	if err != nil {
		log.Get().Panic(err)
	}

	// get trade config
	tradeConf := config.GetTradeConfig()
	// save trade day map to cache
	tradeDayArr, err := dbagent.Get().GetAllTradeDayDate()
	if err != nil {
		log.Get().Panic(err)
	}
	tmp := make(map[time.Time]bool)
	for _, v := range tradeDayArr {
		if v.IsTradeDay {
			tmp[v.Date] = true
			cache.GetCache().SetCalendarID(v)
		}
	}
	cache.GetCache().SetCalendar(tmp)

	// get trade day and save to cache
	tradeDay, err := GetTradeDay()
	if err != nil {
		log.Get().Panic(err)
	}
	cache.GetCache().SetTradeDay(tradeDay)

	// trade out time
	tradeOutEndTime := tradeDay.Add(openTime).Add(time.Duration(tradeConf.TradeOutEndTime) * time.Hour)
	cache.GetCache().SetTradeDayTradeOutEndTime(tradeOutEndTime)

	// trade day end time
	tradeDayEndTime := tradeDay.Add(openTime).Add(time.Duration(tradeConf.TotalOpenTime) * time.Hour)
	cache.GetCache().SetTradeDayOpenEndTime(tradeDayEndTime)

	log.Get().WithFields(map[string]interface{}{
		"Date":            tradeDay.Format(global.ShortTimeLayout),
		"TradeOutEndTime": tradeOutEndTime.Format(global.LongTimeLayout),
	}).Info("TradeDay")

	// every 10 seconds to check if now is open time
	go func() {
		for range time.Tick(10 * time.Second) {
			allowTrade := checkIsAllowTrade(tradeDay, tradeConf.HoldTimeFromOpen, tradeConf.TradeInEndTime)
			cache.GetCache().SetIsAllowTrade(allowTrade)
		}
	}()

	// get config
	closeRange := GetLastNTradeDayByDate(config.GetTradeConfig().HistoryClosePeriod, tradeDay)
	tickRange := GetLastNTradeDayByDate(config.GetTradeConfig().HistoryTickPeriod, tradeDay)
	kbarRange := GetLastNTradeDayByDate(config.GetTradeConfig().HistoryKbarPeriod, tradeDay)

	// save to cahce
	cache.GetCache().SetHistroyCloseRange(closeRange)
	cache.GetCache().SetHistroyTickRange(tickRange)
	cache.GetCache().SetHistroyKbarRange(kbarRange)
}

func checkIsAllowTrade(tradeDay time.Time, holdTimeFromOpen, tradInEndTime float64) bool {
	starTime := tradeDay.Add(openTime).Add(time.Duration(holdTimeFromOpen) * time.Minute)
	if time.Now().After(starTime) && time.Now().Before(starTime.Add(time.Duration(tradInEndTime)*time.Hour)) {
		return true
	}
	return false
}
