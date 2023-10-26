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
	// Serve display
	r := gin.Default()
	r.LoadHTMLGlob("display/*.html")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "display.html", nil)
	})

	go func() {
		r.Run(":3000")
	}()
	start := time.Now()
	// render it and generate bitmap
	img := render.RenderBitmap("http://localhost:3000")
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)

	// send bitmap to vocore
	vocore.SendToScreen(img)

}
