package schemas

import (
	"github.com/project-eria/xaal-go/device"

	logger "github.com/project-eria/eria-logger"
)

// Lamp : Simple switch lamp
func Lamp(addr string) *device.Device {
	dev, _ := device.New("lamp.basic", addr)

	// -- Attributes --
	// State of the lamp
	dev.NewAttribute("light", nil)

	// -- Methods --
	dev.AddMethod("on", defaultOn)
	dev.AddMethod("off", defaultOff)

	return dev
}

func defaultOn(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Switch on the lamp"""
	logger.Module("xaal:schema-lamp").Debug("defaultOn()")
	return nil
}

func defaultOff(d *device.Device, args map[string]interface{}) map[string]interface{} {
	// """Switch off the lamp"""
	logger.Module("xaal:schema-lamp").Debug("defaultOff()")
	return nil
}
