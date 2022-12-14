// Package cloudevent package cloudevent
package cloudevent

import (
	"time"
	"trade_agent/pkg/dbagent"
	"trade_agent/pkg/eventbus"
	"trade_agent/pkg/log"
	"trade_agent/pkg/mqhandler"
	"trade_agent/pkg/pb"
)

// InitCloudEvent InitCloudEvent
func InitCloudEvent() {
	log.Get().Info("Initial CloudEvent")

	err := subTradeEvent()
	if err != nil {
		log.Get().Panic(err)
	}
}

// subTradeEvent subTradeEvent
func subTradeEvent() error {
	handler := mqhandler.Get()
	err := handler.Sub(mqhandler.MQSubBody{
		MQTopic:  mqhandler.TopicTradeEvent(),
		Once:     false,
		Callback: tredeEventCallback,
	})
	if err != nil {
		return err
	}
	return nil
}

// tredeEventCallback tredeEventCallback
func tredeEventCallback(m mqhandler.MQMessage) {
	var err error
	body := pb.EventResponse{}
	err = body.UnmarshalProto(m.Payload())
	if err != nil {
		log.Get().Panic(err)
	}

	err = dbagent.Get().InsertCloudEvent(body.ToTradeEvent())
	if err != nil {
		log.Get().Panic(err)
	}

	if body.GetEventCode() != 16 {
		log.Get().WithFields(map[string]interface{}{
			"EventCode": body.GetEventCode(),
			"Event":     body.GetEvent(),
			"RespCode":  body.GetRespCode(),
			"Info":      body.GetInfo(),
		}).Warn("TradeEvent")
	}

	if body.GetEventCode() == 4 {
		eventbus.Get().PublishRestartSinopacMQSRV(time.Now())
	}
}
