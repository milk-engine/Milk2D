package milk2d

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk-engine/Milk2D/core"
	assets "github.com/milk-engine/Milk2D/core/asset"
	"github.com/milk-engine/Milk2D/core/renderer"
	"github.com/milk-engine/Milk2D/core/vector"
)

type Milk2DEngine struct {
	Renderer    *renderer.Renderer
	Window      *renderer.Window
	AssetLoader *assets.AssetsLoader
}

func New() *Milk2DEngine {
	// initialize all systems
	return &Milk2DEngine{
		Renderer:    &renderer.Renderer{},
		Window:      &renderer.Window{},
		AssetLoader: &assets.AssetLoader,
	}
}

func (engine *Milk2DEngine) Run() {
	fmt.Println("Running Milk2D engine")
	engine.Renderer.Run()
}

type Vector2 = vector.Vector2

func NewVector2(x, y float32) *Vector2 {
	return vector.NewVector2(x, y)
}

type Transform = vector.Transform

func NewTransform() *Transform {
	return vector.NewTransform()
}

type Component = core.Component
type GameObject = core.GameObject

func NewGameObject(name string, start Transform) *GameObject {
	return core.NewGameObject(name, start)
}

type Sprite = assets.Sprite

func NewSprite(spriteName string, dest *Vector2, options *ebiten.DrawImageOptions) *Sprite {
	return assets.NewSprite(spriteName, dest, options)
}
