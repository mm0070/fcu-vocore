package render

import (
	"bytes"
	"image"
	_ "image/png" // Import the PNG package to decode PNG images
	"time"

	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func GetDriver() (driver selenium.WebDriver, service *selenium.Service, err error) {
	// initialize a Chrome browser instance on port 4444
	service, err = selenium.NewChromeDriverService("./render/chromedriver", 4444)

	if err != nil {
		log.Fatal("Error:", err)
	}

	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless", // comment out this line for testing
		"window-size=800,480",
	}})

	// create a new remote client with the specified options
	driver, err = selenium.NewRemote(caps, "")
	if err != nil {
		log.Printf("Error instantiating Selenium driver: %v", err)
		return
	}

	return
}

func RenderBitmap(url string, driver selenium.WebDriver) (screenshot []byte, err error) {
	// visit the target page
	startTime := time.Now()
	err = driver.Get(url)
	endTime := time.Since(startTime)
	log.Printf("Page visit took %s", endTime)
	if err != nil {
		log.Printf("GET request failed: %v", err)
		return
	}

	// Capture a screenshot of the whole page
	startTime = time.Now()
	screenshot, err = driver.Screenshot()
	endTime = time.Since(startTime)
	log.Printf("Screenshot capture took %s", endTime)
	if err != nil {
		log.Printf("Failed to capture screenshot: %v", err)
		return
	}

	// Convert to RGB565
	startTime = time.Now()
	imgData, _, err := image.Decode(bytes.NewReader(screenshot))
	endTime = time.Since(startTime)
	log.Printf("Decoding image took %s", endTime)
	if err != nil {
		log.Printf("Failed to decode image format: %v", err)
		return
	}

	startTime = time.Now()
	rgb565Data := convertToRGB565(imgData)
	endTime = time.Since(startTime)
	log.Printf("RGB565 conversion took %s", endTime)
	// TODO return and handle an error here

	return rgb565Data, nil
}
