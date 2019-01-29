package schemas

import (
	"github.com/Eria-Project/xaal-go/device"

	"github.com/Eria-Project/logger"
)

// Shutter returns a simple shutter
func Shutter(addr string) *device.Device {
	dev, _ := device.New("shutter.basic", addr)

	// -- Attributes --
	// Ongoing action of the shutter
	dev.NewAttribute("action", nil)

	// -- Methods --
	dev.AddMethod("up", defaultUp)
	dev.AddMethod("down", defaultDown)
	dev.AddMethod("stop", defaultStop)

	return dev
}

// ShutterPosition returns a shutter with a position managment
func ShutterPosition(addr string) *device.Device {
	dev, _ := device.New("shutter.position", addr)

	// -- Attributes --
	// Ongoing action of the shutter
	dev.NewAttribute("action", nil)
	// Level of aperture of the shutter
	dev.NewAttribute("position", nil)

	// -- Methods --
	dev.AddMethod("up", defaultUp)
	dev.AddMethod("down", defaultDown)
	dev.AddMethod("stop", defaultStop)
	dev.AddMethod("position", defaultPosition)

	return dev
}

func defaultUp(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Up the shutter"""
	logger.Debug("defaultUp()")
	return nil
}

func defaultDown(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Down the shutter"""
	logger.Debug("defaultDown()")
	return nil
}

func defaultStop(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Stop ongoing action of the shutter"""
	logger.Debug("defaultStop()")
	return nil
}

func defaultPosition(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Change the position of the shutter"""
	logger.WithField("target", args["target"]).Debug("defaultPosition()")
	return nil
}