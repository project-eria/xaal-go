package schemas

import (
	"github.com/project-eria/xaal-go/device"
)

// Windgauge : Simple windgauge
func Windgauge(addr string) (*device.Device, error) {
	dev, err := Basic(addr)
	dev.SetDevType("windgauge.basic")

	// -- Attributes --
	// Strength of the wind
	dev.NewAttribute("windStrength", nil)
	// Direction of the wind
	dev.NewAttribute("windAngle", nil)
	// Strength of gusts
	dev.NewAttribute("gustStrength", nil)
	// Direction of gusts
	dev.NewAttribute("gustAngle", nil)

	return dev, err
}
