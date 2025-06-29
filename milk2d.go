package milk2d

import (
	"fmt"

	"github.com/milk-engine/Milk2D/core"
	assets "github.com/milk-engine/Milk2D/core/asset"
	"github.com/milk-engine/Milk2D/core/renderer"
	"github.com/milk-engine/Milk2D/core/vector"
)

type Milk2DEngine struct {
	SceneManager *core.ScenesManager
	Renderer     *renderer.Renderer
	Window       *renderer.Window
	AssetLoader  *assets.AssetsLoader
}

func New() *Milk2DEngine {
	// initialize all systems
	return &Milk2DEngine{
		SceneManager: core.SceneManager,
		Renderer:     &renderer.Renderer{SceneManager: core.SceneManager},
		Window:       &renderer.Window{},
		AssetLoader:  &assets.AssetLoader,
	}
}

func (engine *Milk2DEngine) Run() {
	fmt.Println("Running Milk2D engine")
	engine.Renderer.Run()
}

type Vector2 = vector.Vector2
type Transform = vector.Transform
type Component = core.Component
type GameObject = core.GameObject

func NewGameObject(name string, startPosition *Vector2) *GameObject {
	return core.NewGameObject(name, startPosition)
}

type Sprite = assets.Sprite
type SpriteDestination = assets.SpriteDestination

func NewSprite(spriteName string, destination *assets.SpriteDestination) *Sprite {
	return assets.NewSprite(spriteName, destination)
}
