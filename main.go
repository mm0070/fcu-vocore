package main

import (
	"log"
	"time"

	"github.com/mm0070/fcu-vocore/render"
	"github.com/mm0070/fcu-vocore/vocore"
)

func main() {
	display, err := vocore.InitializeScreen()
	if err != nil {
		log.Fatalf("Failed to initialize screen: %v", err)
	}
	defer display.Close()

	driver, service, err := render.GetChromeDriver()
	if err != nil {
		log.Fatalf("Failed to initialize chrome driver: %v", err)
	}
	defer service.Stop()

	// Channels to pipeline the stages.
	captureDone := make(chan []byte, 1) // Buffer for one screenshot.
	convertDone := make(chan []byte, 1) // Buffer for one converted image.

	// Error channel to handle errors in goroutines.
	errChan := make(chan error)

	// Capture routine.
	go func() {
		for {
			bitmapData, err := render.CaptureHTML("file:///C:/Display/display.html", driver)
			if err != nil {
				errChan <- err
				return
			}
			captureDone <- bitmapData
		}
	}()

	// Convert routine.
	go func() {
		for {
			select {
			case bitmapData := <-captureDone:
				img := render.ConvertToRGB565(bitmapData)
				convertDone <- img
			case err := <-errChan:
				log.Printf("Error during capture: %v", err)
				return
			}
		}
	}()

	// Write routine.
	go func() {
		var lastFrameTime time.Time
		for {
			select {
			case img := <-convertDone:
				start := time.Now()

				err := display.WriteToScreen(img)
				if err != nil {
					errChan <- err
					return
				}

				// Calculate write duration and FPS.
				writeDuration := time.Since(start)
				log.Printf("Write to screen took: %v ms", writeDuration.Milliseconds())

				if !lastFrameTime.IsZero() {
					elapsed := time.Since(lastFrameTime)
					fps := float64(time.Second) / float64(elapsed)
					log.Printf("FPS: %.2f", fps)
				}
				lastFrameTime = time.Now()
			case err := <-errChan:
				log.Printf("Error during conversion or write: %v", err)
				return
			}
		}
	}()

	// Handle exit in main routine (e.g., on error or interrupt).
	for {
		select {
		case err := <-errChan:
			log.Printf("Exiting due to error: %v", err)
			return
		}
	}
}
