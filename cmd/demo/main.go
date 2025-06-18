package main

import (
	"github.com/milk-engine/Milk2D/engine"
)

func main() {
	milk2D := engine.New()
	milk2D.Renderer.WindowSetTitle("Milk2D demo")
	milk2D.Renderer.WindowSetSize(1280, 720)
	milk2D.Run()
}
