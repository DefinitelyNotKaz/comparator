package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strings"
)

type ColorEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func hexToRGBA(hex string) (color.RGBA, error) {
	var r, g, b uint8
	if _, err := fmt.Sscanf(strings.TrimPrefix(hex, "#"), "%02x%02x%02x", &r, &g, &b); err != nil {
		return color.RGBA{}, err
	}
	return color.RGBA{r, g, b, 255}, nil
}

func colorInPalette(color color.RGBA, palette []color.RGBA) bool {
	if color.A == 0 {
		return true
	}
	for _, val := range palette {
		if color == val {
			return true
		}
	}
	return false
}

func convertColorToRGBA(value color.Color) color.RGBA {
	r, g, b, a := value.RGBA()
	return color.RGBA{
		uint8(r >> 8),
		uint8(g >> 8),
		uint8(b >> 8),
		uint8(a >> 8),
	}
}

func compare(file string, paletteFile string, verbose bool) {
	imgFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	jsonFile, err := os.Open(paletteFile)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer jsonFile.Close()

	var colors []ColorEntry
	if err := json.NewDecoder(jsonFile).Decode(&colors); err != nil {
		fmt.Println("Error decoding JSON file: ", err)
		return
	}

	var palette []color.RGBA
	for _, entry := range colors {
		rgbaColor, err := hexToRGBA(entry.Value)
		if err != nil {
			fmt.Printf("Error converting color %s: %v\n", entry.Name, err)
			return
		}
		palette = append(palette, rgbaColor)
	}

	bounds := img.Bounds()

	diff := image.NewRGBA(bounds)
	draw.Draw(diff, bounds, img, image.Point{0, 0}, draw.Src)

	// Lower the opacity of the original art
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {

			originalColor := diff.At(x, y).(color.RGBA)

			newR := uint8(float64(originalColor.R) * 0.1)
			newG := uint8(float64(originalColor.G) * 0.1)
			newB := uint8(float64(originalColor.B) * 0.1)

			diff.SetRGBA(x, y, color.RGBA{newR, newG, newB, originalColor.A})
		}
	}

	count := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			pixelColor := convertColorToRGBA(img.At(x, y))

			if !colorInPalette(pixelColor, palette) {
				if verbose {
					fmt.Printf("Pixel at (%d, %d) is not in the palette\n", x, y)
				}
				diff.SetRGBA(x, y, color.RGBA{255, 0, 0, 255})
				count++
			}
		}
	}

	if count == 0 {
		fmt.Println("All the pixels matched!")
		os.Exit(1)
	}

	outputFile, err := os.Create("result.png")
	if err != nil {
		fmt.Println("Error creating output image file:", err)
		return
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, diff)
	if err != nil {
		fmt.Println("Error encoding output image:", err)
		return
	}

	fmt.Println("Image processing complete, saved as result.png")
}
