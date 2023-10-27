package vocore

import (
	"github.com/google/gousb"
)

func setFrame(intf *gousb.Interface, data []byte) error {
	// Out transfer to send frame data to the screen
	outEp, err := intf.OutEndpoint(EndpointID)
	if err != nil {
		return err
	}

	_, err = outEp.Write(data)
	return err
}

func sendPixelData(intf *gousb.Interface, pixelData []byte) error {
	// Call setFrame to send pixel data to the USB display
	err := setFrame(intf, pixelData)
	return err
}
