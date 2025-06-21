package milk2d

import (
	"fmt"

	"github.com/milk-engine/Milk2D/core"
	"github.com/milk-engine/Milk2D/core/renderer"
	"github.com/milk-engine/Milk2D/core/vector"
)

type Vector2 = vector.Vector2
type Transform = vector.Transform
type Component = core.Component
type IComponent = core.IComponent
type GameObject = core.GameObject
type IGameObject = core.IGameObject

type Milk2DEngine struct {
	Renderer *renderer.Renderer
	Window   *renderer.Window
}

func New() *Milk2DEngine {
	// initialize all systems
	return &Milk2DEngine{
		Renderer: &renderer.Renderer{},
		Window:   &renderer.Window{},
	}
}

func (engine *Milk2DEngine) Run() {
	fmt.Println("Running Milk2D engine")
	engine.Renderer.Run()
}
