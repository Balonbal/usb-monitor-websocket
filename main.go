package main

import (
	"github.com/Balonbal/usb-monitor-websocket/handlers"
	"net/http"
	"github.com/google/gousb"
	"flag"
)

func main() {
	flag.Parse()
	usbCtx := gousb.NewContext()
	usbCh := make(chan interface{})
	clientCh := handlers.InitBroadcaster(usbCh)

	http.HandleFunc("/listUsbs", func(w http.ResponseWriter, r *http.Request) { handlers.ListUsb(w, r, usbCtx) })
	http.HandleFunc("/monitor", func(w http.ResponseWriter, r *http.Request) {
		handlers.RegisterMonitor(w, r, clientCh)
	})

	defer usbCtx.Close()
	panic(http.ListenAndServe(":19123", nil))
}
