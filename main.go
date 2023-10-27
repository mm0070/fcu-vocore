package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mm0070/fcu-vocore/render"
	"github.com/mm0070/fcu-vocore/vocore"
)

func main() {
	// Serve display HTML
	r := gin.Default()
	r.LoadHTMLGlob("display/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "display.html", nil)
	})

	go func() {
		r.Run(":3000")
	}()

	// Set up display
	display, err := vocore.InitializeScreen()
	if err != nil {
		log.Fatal("Failed to initialize screen: %v", err)
	}
	defer display.Close()

	// Initialize chrome driver
	driver, service, err := render.GetDriver()
	if err != nil {
		log.Fatal("Failed to initialize chrome driver: %v", err)
	}
	defer service.Stop()

	for true {
		frameRenderTimeStart := time.Now()
		img, _ := render.RenderBitmap("http://localhost:3000", driver)
		//TODO handle error

		_ = display.WriteToScreen(img)
		frameTime := time.Since(frameRenderTimeStart)
		log.Printf("Screen write took: %s", frameTime)
		fps := time.Second / frameTime
		log.Printf("FPS: %v", fps)
	}
}
