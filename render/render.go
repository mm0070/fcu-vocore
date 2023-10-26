package render

import (
	"bytes"
	"image"
	_ "image/png" // Import the PNG package to decode PNG images

	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func RenderBitmap(url string) (screenshot []byte) {
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./render/chromedriver", 4444)

	if err != nil {
		log.Fatal("Error:", err)
	}

	defer service.Stop()

	// configure the browser options

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless", // comment out this line for testing
		"window-size=800,480",
	}})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// visit the target page
	err = driver.Get(url)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Capture a screenshot of the whole page
	screenshot, err = driver.Screenshot()
	if err != nil {
		log.Fatal(err)
	}

	// Convert to RGB565
	imgData, _, err := image.Decode(bytes.NewReader(screenshot))
	if err != nil {
		log.Fatal(err)
	}

	rgb565Data := convertToRGB565(imgData)

	return rgb565Data
}
