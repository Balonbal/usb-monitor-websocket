package handlers

import "net/http"
import "fmt"
import "encoding/json"
import "github.com/google/gousb"
import "github.com/google/gousb/usbid"
import "github.com/gorilla/websocket"
import "github.com/Balonbal/usb-monitor-websocket/usb"

type UsbDesc struct {
	*gousb.DeviceDesc
	Description string
}

var _clients = make(map[*websocket.Conn]bool)
var _upgrader = websocket.Upgrader{}

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
	w.Header().Set("Content-Type", "application/json")
	i, err := w.Write([]byte(mess))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(i)
}

func RegisterMonitor(w http.ResponseWriter, r *http.Request, clientCh chan<- *websocket.Conn) {
	socket, err := _upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
	}
	clientCh <- socket

}

func broadcast(eventCh <-chan interface{}, clientCh <-chan *websocket.Conn) {
	clients := make(map[*websocket.Conn]bool)
	for {
		select {
		case client := <-clientCh:
			clients[client] = true
		case event := <-eventCh:
			message, err := json.Marshal(event)
			if err != nil {
				fmt.Println(err)
			}
			for client := range clients {
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					fmt.Println(err)
					client.Close()
					delete(clients, client)
				}
			}
		}
	}
}

func InitBroadcaster(eventCh chan interface{}) (chan<- *websocket.Conn) {
	clientCh := make(chan *websocket.Conn)
	go broadcast(eventCh, clientCh)
	return clientCh
}
