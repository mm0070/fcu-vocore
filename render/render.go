package render

import (
	_ "image/png" // Import the PNG package to decode PNG images
	"time"

	"log"

	"github.com/tebeka/selenium"
)

func CaptureHTML(url string, driver selenium.WebDriver) (bitmapData []byte, err error) {
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
	bitmapData, err = driver.Screenshot()
	endTime = time.Since(startTime)
	log.Printf("Screenshot capture took %s", endTime)
	if err != nil {
		log.Printf("Failed to capture screenshot: %v", err)
		return
	}
	return
}
