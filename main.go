package main

import (
	"github.com/Balonbal/usb-monitor-websocket/handlers"
	"net/http"
	"github.com/google/gousb"
)

func main() {
	usbCtx := gousb.NewContext()

	http.HandleFunc("/listUsbs", func(w http.ResponseWriter, r *http.Request) { handlers.ListUsb(w, r, usbCtx) });

	panic(http.ListenAndServe(":19123", nil))
}
