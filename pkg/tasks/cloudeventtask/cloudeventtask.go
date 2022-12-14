// Package cloudeventtask package cloudeventtask
package cloudeventtask

import (
	"errors"
	"runtime/debug"
	"sync"
	"trade_agent/pkg/dbagent"
	"trade_agent/pkg/log"
)

var lock sync.Mutex

// Run Run
func Run() {
	lock.Lock()
	var err error
	defer func() {
		if r := recover(); r != nil {
			switch x := r.(type) {
			case string:
				err = errors.New(x)
			case error:
				err = x
			default:
				err = errors.New("unknown panic")
			}
			log.Get().Error(err.Error() + "\n" + string(debug.Stack()))
		}
	}()
	defer lock.Unlock()

	err = dbagent.Get().DeleteAllCloudEvent()
	if err != nil {
		log.Get().Panic(err)
	}
}
