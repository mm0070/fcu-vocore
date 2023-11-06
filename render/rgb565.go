package render

import (
	"bytes"
	"encoding/binary"
	"image"
	_ "image/png" // Import the PNG package to decode PNG images
	"log"
)

func ConvertToRGB565(bitmapData []byte) []byte {
	img, _, err := image.Decode(bytes.NewReader(bitmapData))
	if err != nil {
		log.Fatalf("Failed to decode image: %v", err)
	}

	width := img.Bounds().Dy()  // Swap width and height
	height := img.Bounds().Dx() // Swap width and height

	// Prepare a byte slice for the RGB565 data
	rgb565Data := make([]byte, width*height*2) // 2 bytes for each pixel in RGB565 format

	index := 0
	for y := height - 1; y >= 0; y-- { // Start from the last row and move upwards
		for x := 0; x < width; x++ {
			// Get the color of the current pixel
			r, g, b, _ := img.At(y, x).RGBA() // Swap x and y here

			// Convert to RGB565 format
			rgb565 := toRGB565(int(r>>8), int(g>>8), int(b>>8))
			binary.LittleEndian.PutUint16(rgb565Data[index:], uint16(rgb565))
			index += 2
		}
	}

	return rgb565Data
}

func toRGB565(r, g, b int) int {
	r = r >> 3
	g = g >> 2
	b = b >> 3
	return (r << 11) | (g << 5) | b
}
