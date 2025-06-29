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
	sprite := milk2d.NewSprite("test", &milk2d.SpriteDestination{X: 0, Y: 0, SizeX: 256, SizeY: 256}, &ebiten.DrawImageOptions{})
	w, h := milk2D.Window.GetSize()
	centered := &milk2d.Vector2{X: (float32(w) / 2), Y: float32(h) / 2}

	cow := milk2d.NewGameObject("Cow", centered)
	cow.AddComponent(sprite)
	cow.AddToScene("default scene")

	milk2D.Run()
}
