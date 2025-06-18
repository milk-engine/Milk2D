package engine

import (
	"fmt"

	"github.com/milk-engine/Milk2D/engine/renderer"
)

type Engine struct {
	Renderer *renderer.Renderer
	// PhysicsWorld  *PhysicsWorld
	// SceneManager  *SceneManager
	// Input         *Input
	// EventBus      *EventBus
	// Audio         *Audio
	// AssetLoader   *AssetLoader
	// DebugTools    *DebugTools
	// GameObjectMgr *GameObjectManager
}

func New() *Engine {
	return &Engine{}
}

func (e *Engine) Run() {
	fmt.Println("Running engine")
	e.Renderer.Run()
}
