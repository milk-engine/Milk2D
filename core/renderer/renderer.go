package renderer

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk-engine/Milk2D/core"
)

type Renderer struct {
	ebiten.Game
	SceneManager *core.ScenesManager
}

func (r *Renderer) Run() {
	r.SceneManager.InitScene()
	if err := ebiten.RunGame(r); err != nil {
		log.Fatal(err)
	}
}

func (r *Renderer) Update() error {
	r.SceneManager.UpdateScene()
	return nil
}

func (r *Renderer) Draw(screen *ebiten.Image) {
	r.SceneManager.RenderScene(screen)
}

func (r *Renderer) Layout(w, h int) (int, int) {
	return w, h
}
