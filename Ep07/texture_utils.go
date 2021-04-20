package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func drawTexture(
	tex *sdl.Texture,
	position vector,
	rotation float64,
	renderer *sdl.Renderer) error {

	_, _, width, height, err := tex.Query()
	if err != nil {
		return fmt.Errorf("querying texture: %v", err)
	}

	// Convert coordinates to the top left of the sprite
	position.x -= float64(width) / 2.0
	position.y -= float64(height) / 2.0

	return renderer.CopyEx(
		tex,
		&sdl.Rect{},
		&sdl.Rect{},
		rotation,
		&sdl.Point{width / 2, height / 2},
		sdl.FLIP_NONE)
}

func loadTextureFromBMP(filename string, renderer *sdl.Renderer) (*sdl.Texture, error) {
	img, err := sdl.LoadBMP(filename)
	if err != nil {
		return nil, fmt.Errorf("loading %v: %v", filename, err)
	}
	defer img.Free()
	tex, err := renderer.CreateTextureFromSurface(img)
	if err != nil {
		return nil, fmt.Errorf("creating texture from %v: %v", filename, err)
	}

	return tex, nil
}
