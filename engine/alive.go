// Package engine : Alive messages
package engine

import (
	"time"

	"github.com/Eria-Project/xaal-go/device"
	"github.com/Eria-Project/xaal-go/messagefactory"

	"github.com/Eria-Project/logger"
)

var _tickerAlive *time.Ticker

//TODO		self.__alives = []                       # list of alive devices

// SendAlive : Send a Alive message for a given device
func SendAlive(dev *device.Device) {
	timeout := dev.GetTimeout()
	msg, err := messagefactory.BuildAliveFor(dev, timeout)
	if err != nil {
		logger.Module("engine").WithError(err).Error("Cannot build alive message")
	} else {
		logger.Module("engine").WithField("from", dev.Address).Debug("Sending alive message")
		_queueMsgTx <- msg
	}
}

func sendAlives() {
	for _, dev := range _devices {
		SendAlive(dev)
	}
}

// SendIsAlive : Send a isAlive message, w/ devTypes filtering
func SendIsAlive(dev *device.Device, devTypes string) {
	body := make(map[string]interface{})
	body["devTypes"] = devTypes
	msg, err := messagefactory.BuildMsg(dev, []string{}, "request", "isAlive", body)
	if err != nil {
		logger.Module("engine").WithError(err).Error("Cannot build isAlive message")
	} else {
		_queueMsgTx <- msg
	}
}

// processAlives : Periodic sending alive messages
func processAlives(aliveTimer uint16) {
	_tickerAlive = time.NewTicker(time.Duration(aliveTimer) * time.Second)
	go func() {
		logger.Module("engine").Debug("Send initial alive messages")
		sendAlives()
		for range _tickerAlive.C {
			logger.Module("engine").Debug("Send alive messages")
			sendAlives()
		}
	}()
}
