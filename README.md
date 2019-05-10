# usb-monitor-websocket

## Linux
- Install libusb
- `go get github.com/Balonbal/usb-monitor-websocket`
- Run

## Windows

- Install [MSYS2](https://www.msys2.org/)
- `pacman -S mingw64/mingw-w64-x86_64-gcc mingw64/mingw-w64-x86_64-libusb mingw64/mingw-w64-x86_64-go pkg-config`
- Make sure you are running the mingw64 aplication (NOT the msys2)
- Do everything else described [here](https://github.com/google/gousb/blob/master/.appveyor/install.sh)

Now because windows is windows, and windows sucks, you will have to manually install drivers to be able to read the raw data from an usb port. (Does not apply to hid-devices). Zadig is an useful tool for this.
