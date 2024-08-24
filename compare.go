package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type ColorEntry struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func Compare(options Options) error {
	imgFile, err := os.Open(options.File)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	palette, err := LoadPalette(options.Palette)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}

	var colors []color.RGBA
	for _, entry := range palette {
		rgbaColor, err := HexToRGBA(entry.Value)
		if err != nil {
			return fmt.Errorf("error converting color %s: %v", entry.Name, err)
		}
		colors = append(colors, rgbaColor)
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
			pixelColor := ConvertToRGBA(img.At(x, y))

			if !ColorInPalette(pixelColor, colors) {
				if options.Verbose {
					fmt.Printf("Pixel at (%d, %d) is not in the palette\n", x, y)
				}
				diff.SetRGBA(x, y, color.RGBA{255, 0, 0, 255})
				count++
			}
		}
	}

	if count == 0 {
		fmt.Println("All the pixels matched!")
		return nil
	}

	outputFile, err := os.Create(options.Output)
	if err != nil {
		return fmt.Errorf("error creating output image: %w", err)
	}
	defer outputFile.Close()

	err = png.Encode(outputFile, diff)
	if err != nil {
		fmt.Println("Error encoding output image:", err)
		return fmt.Errorf("error encoding output image: %w", err)
	}

	fmt.Printf("Image processing complete, saved as %s\n", options.Output)
	return nil
}
