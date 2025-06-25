package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	milk2d "github.com/milk-engine/Milk2D"
)

func main() {
	milk2D := milk2d.New()
	milk2D.Window.SetTitle("Milk2D demo")
	milk2D.Window.SetSize(1280, 720)
	// this is how you can load images/audio into a project
	// and you can then reference them
	milk2D.AssetLoader.LoadAsset("test", "/examples/assets/cow.png")
	sprite := milk2d.NewSprite("test", milk2d.NewVector2(64, 64), &ebiten.DrawImageOptions{})

	cow := milk2d.NewGameObject("Cow", *milk2d.NewTransform())
	cow.AddComponent(sprite)
	cow.AddToScene("default scene")

	milk2D.Run()
}
