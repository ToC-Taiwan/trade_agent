// Package sinopacapi package sinopacapi
package sinopacapi

const (
	urlFetchAllStockDetail                string = "/sinopac-mq-srv/basic/stock-detail"
	urlFetchAllSnapShot                   string = "/sinopac-mq-srv/real-time/all-snapshot"
	urlFetchSnapShots                     string = "/sinopac-mq-srv/real-time/snapshots"
	urlFetchTSESnapShot                   string = "/sinopac-mq-srv/real-time/snapshot/tse"
	urlFetchHistoryTickByStockAndDate     string = "/sinopac-mq-srv/history/tick"
	urlFetchHistoryKbarByDateRange        string = "/sinopac-mq-srv/history/kbar"
	urlFetchHistoryTSEKbarByDate          string = "/sinopac-mq-srv/history/kbar/tse"
	urlFetchHistoryCloseByStockDateArr    string = "/sinopac-mq-srv/history/close"
	urlFetchHistoryTSECloseByDate         string = "/sinopac-mq-srv/history/close/tse"
	urlFetchHistoryCloseByStockArrDateArr string = "/sinopac-mq-srv/history/close/multi-date"
	urlFetchVolumeRankByDate              string = "/sinopac-mq-srv/history/volumerank"
	urlSubBidAsk                          string = "/sinopac-mq-srv/subscribe/bid-ask"
	urlUnSubBidAskByStock                 string = "/sinopac-mq-srv/unsubscribe/bid-ask"
	urlUnSubscribeAllBidAsk               string = "/sinopac-mq-srv/unsubscribeall/bid-ask"
	urlSubRealTimeTick                    string = "/sinopac-mq-srv/subscribe/realtime-tick"
	urlUnSubRealTimeTickByStock           string = "/sinopac-mq-srv/unsubscribe/realtime-tick"
	urlUnSubscribeAllRealTimeTick         string = "/sinopac-mq-srv/unsubscribeall/realtime-tick"
	urlPlaceOrderBuy                      string = "/sinopac-mq-srv/trade/buy"
	urlPlaceOrderSell                     string = "/sinopac-mq-srv/trade/sell"
	urlPlaceOrderSellFirst                string = "/sinopac-mq-srv/trade/sell_first"
	urlCancelOrder                        string = "/sinopac-mq-srv/trade/cancel"
	urlFetchOrderStatus                   string = "/sinopac-mq-srv/trade/status"
	urlFetchOrderStatusByOrderID          string = "/sinopac-mq-srv/trade/order/status"
	urlFetchServerKey                     string = "/sinopac-mq-srv/system/healthcheck"
	urlAskSinpacMQSRVConnectMQ            string = "/sinopac-mq-srv/system/mq-connect"
	urlRestartSinopacSRV                  string = "/sinopac-mq-srv/system/restart"
	// urlFetchHistoryTSETickByDate          string = "/sinopac-mq-srv/history/tick/tse"
)
