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

func SendToScreen(img []byte) {
	// Initialize a new USB context
	ctx := gousb.NewContext()
	defer ctx.Close()

	// Open the USB device by its VendorID and ProductID
	dev, err := ctx.OpenDeviceWithVIDPID(VendorID, ProductID)
	if err != nil {
		log.Fatalf("Could not open a device: %v", err)
	}
	defer dev.Close()

	// Claim an interface on the USB device
	intf, done, err := dev.DefaultInterface()
	if err != nil {
		log.Fatalf("Error claiming interface: %v", err)
	}
	defer done()

	// Wake up the screen
	_, err = dev.Control(0x40, 0xB0, 0, 0, []byte{0x00, 0x29})
	if err != nil {
		log.Fatalf("Error waking up the screen: %v", err)
	}

	// Call writeStart to prepare the screen for frame data
	_, err = dev.Control(0x40, 0xB0, 0, 0, []byte{0x00, 0x2C, 0x00, 0xB8, 0x0B, 0x00})
	if err != nil {
		log.Fatalf("Error in writeStart: %v", err)
	}

	// Send pixel data to the USB display using blit
	err = blit(intf, img)
	if err != nil {
		log.Fatalf("Error in blit: %v", err)
	}

	log.Println("Pixel data sent successfully to the USB display!")
}
