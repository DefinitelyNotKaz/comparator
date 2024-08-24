package main

import (
	"fmt"
	"image/color"
	"strings"
)

func ConvertToRGBA(value color.Color) color.RGBA {
	r, g, b, a := value.RGBA()
	return color.RGBA{
		uint8(r >> 8),
		uint8(g >> 8),
		uint8(b >> 8),
		uint8(a >> 8),
	}
}

func HexToRGBA(hex string) (color.RGBA, error) {
	var r, g, b uint8
	if _, err := fmt.Sscanf(strings.TrimPrefix(hex, "#"), "%02x%02x%02x", &r, &g, &b); err != nil {
		return color.RGBA{}, err
	}
	return color.RGBA{r, g, b, 255}, nil
}

func ColorInPalette(color color.RGBA, palette []color.RGBA) bool {
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
