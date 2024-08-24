package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Palette struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func LoadPalette(file string) ([]Palette, error) {
	paletteFile, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer paletteFile.Close()

	var palettes []Palette
	decoder := json.NewDecoder(paletteFile)
	if err := decoder.Decode(&palettes); err == nil {
		return palettes, nil
	}

	if _, err := paletteFile.Seek(0, 0); err != nil {
		return nil, fmt.Errorf("error seeking file: %w", err)
	}

	var fallback map[string]string
	if err := decoder.Decode(&fallback); err != nil {
		return nil, fmt.Errorf("error decoding JSON as map: %w", err)
	}

	for key, value := range fallback {
		palettes = append(palettes, Palette{Name: key, Value: value})
	}

	return palettes, nil
}
