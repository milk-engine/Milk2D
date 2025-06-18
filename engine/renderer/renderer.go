package renderer

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Renderer struct {
}

func (r *Renderer) Run() {
	if err := ebiten.RunGame(r); err != nil {
		log.Fatal(err)
	}
}

func (r *Renderer) Update() error {
	// TODO: handle arrays of renderable gameobjects
	return nil
}

func (r *Renderer) Draw(screen *ebiten.Image) {
	fmt.Println("Renderering onto screen")
}

func (r *Renderer) Layout(w, h int) (int, int) {
	return w, h
}
