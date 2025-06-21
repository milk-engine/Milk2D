package main

import milk2d "github.com/milk-engine/Milk2D"

func main() {
	milk2D := milk2d.New()
	milk2D.Window.SetTitle("Milk2D demo")
	milk2D.Window.SetSize(1280, 720)
	milk2D.Run()
}
