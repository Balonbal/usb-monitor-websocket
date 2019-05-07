package usb

import (
	"github.com/google/gousb"
)

func ListUsbs(ctx *gousb.Context) ([]gousb.DeviceDesc) {
	var response []gousb.DeviceDesc
	devs, _ := ctx.OpenDevices(func(desc *gousb.DeviceDesc) bool {
		response = append(response, *desc)

		return false
	})


	defer func() {
		for _, d := range devs {
			d.Close()
		}
	}()

	return response
}
