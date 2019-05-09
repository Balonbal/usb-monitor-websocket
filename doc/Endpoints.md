## Endpoints

This usb monitor is interacted with via a http connections, and the main monitor endpoint uses WebSockets to continously push events to the client. Since the program should be configured via the http interface, there is a few endpoints to help with configuration.

#### listUsb
endpoint: **/listUsb**

_Returns a list of the currently connected devices to the computer the program is running on. Response format is as specified in gousb.DeviceDesc, with the added human description field Description, from usbid._

Example response:

```
[
{
	Bus: 1,
	Address: 1,
	Speed: 1,
	Port: 0,
	Spec: 512,
	Device 1045,
	Vendor: 7531,
	Product: 2,
	Class: 9,
	SubClass: 0,
	Protocol: 1,
	MaxControlPacketSize: 64,
	Configs: [
	{
		Number: 1,
		SelfPowered: true,
		RemoteWakeup: true,
		MaxPower: 0,
		Interfaces: [
		{
			Number: 0,
			AltSettings: [
			{
				Number: 0,
				Alternate: 0,
				Class: 9,
				SubClass: 0,
				Protocol: 0,
				Endpoints: {
					129: {
						Address: 129,
						Number: 1,
						Direction: true,
						MaxPacketSize: 4,
						TransferType: 3,
						PollInterval: 256000000,
						IsoSyncType: 0,
						UsageType: 0
					}
				}
			}
		}
		]
	}
	],
	Description: "3.0 root hub (Linux Foundation)"
},
	...
]
```

#### Monitor
Endpoint: **/monitor**

_Continously pushes data to every client connected._
