package handlers

import "net/http"
import "fmt"
import "encoding/json"
import "github.com/google/gousb"
import "github.com/google/gousb/usbid"
import "github.com/Balonbal/usb-monitor-websocket/usb"

type UsbDesc struct {
	*gousb.DeviceDesc
	Description string
}

func ListUsb(w http.ResponseWriter, r *http.Request, ctx *gousb.Context) {
	usbs := usb.ListUsbs(ctx)
	var arr []UsbDesc
	for _, d := range usbs {
		arr = append(arr, UsbDesc{
			&d,
			usbid.Describe(&d),
		})
	}
	mess, _ := json.Marshal(arr)
	i, err := w.Write([]byte(mess))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}
