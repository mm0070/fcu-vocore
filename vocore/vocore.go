package vocore

import (
	"log"

	"github.com/google/gousb"
)

const (
	// Constants for device configuration
	Width      = 800
	Height     = 480
	PixelSize  = 2
	VendorID   = 0xC872
	ProductID  = 0x1004
	EndpointID = 0x2 // Endpoint address for data transfer
)

type VocoreScreen struct {
	dev  *gousb.Device
	intf *gousb.Interface
}

func InitializeScreen() (v *VocoreScreen, err error) {
	// Initialize a new USB context
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Open the USB device by its VendorID and ProductID
	dev, err := ctx.OpenDeviceWithVIDPID(VendorID, ProductID)
	if err != nil {
		log.Printf("Could not open a device: %v", err)
		return
	}

	// Claim an interface on the USB device
	intf, _, err := dev.DefaultInterface()
	if err != nil {
		log.Printf("Error claiming interface: %v", err)
		return
	}
	// TODO 2nd argument here is a done function, make sure to call this

	// Wake up the screen
	_, err = dev.Control(0x40, 0xB0, 0, 0, []byte{0x00, 0x29})
	if err != nil {
		log.Printf("Error waking up the screen: %v", err)
		return
	}

	v = &VocoreScreen{
		dev:  dev,
		intf: intf,
	}

	return
}

func (v *VocoreScreen) WriteToScreen(img []byte) (err error) {
	// Call writeStart to prepare the screen for frame data
	_, err = v.dev.Control(0x40, 0xB0, 0, 0, []byte{0x00, 0x2C, 0x00, 0xB8, 0x0B, 0x00})
	if err != nil {
		log.Printf("Error in writeStart: %v", err)
		return
	}

	// Send pixel data to the USB display
	err = sendPixelData(v.intf, img)
	if err != nil {
		log.Printf("Error in sendPixelData: %v", err)
		return
	}
	return
}

func (v *VocoreScreen) Close() {
	v.dev.Close()
	v.intf.Close()
}
