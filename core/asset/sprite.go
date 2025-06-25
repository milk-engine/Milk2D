package assets

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk-engine/Milk2D/core"
	"github.com/milk-engine/Milk2D/core/vector"
)

type Sprite struct {
	IAsset
	core.IComponent
	Image       *ebiten.Image
	Options     *ebiten.DrawImageOptions
	Transform   vector.Transform
	Destination vector.Vector2
}

func NewSprite(spriteName string, dest *vector.Vector2, options *ebiten.DrawImageOptions) *Sprite {
	src := AssetLoader.GetAsset(spriteName).(*ebiten.Image)

	return &Sprite{
		Image:       src,
		Options:     options,
		Transform:   *vector.NewTransform(),
		Destination: *dest,
	}
}

func (s *Sprite) Init() {
	//
}

func (s *Sprite) Update() {
	//
}
func (s *Sprite) Render(screen *ebiten.Image) {
	screen.DrawImage(s.Image, s.Options)
}
