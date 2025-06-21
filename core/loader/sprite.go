package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/milk-engine/Milk2D/core/vector"
)

type Sprite struct {
	IAsset
	Image       *ebiten.Image
	Transform   vector.Transform
	Destination vector.Vector2
}

func (s *Sprite) Load(fileName string) {
	img, _, err := ebitenutil.NewImageFromFile(fileName)
	s.Image = img
	log.Fatal(err)
}
