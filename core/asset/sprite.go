package assets

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk-engine/Milk2D/core"
	"github.com/milk-engine/Milk2D/core/vector"
)

type SpriteDestination struct {
	X     int
	Y     int
	SizeX int
	SizeY int
}

type Sprite struct {
	IAsset
	core.IComponent
	Parent      *core.GameObject
	Options     *ebiten.DrawImageOptions
	Transform   *vector.Transform
	Destination *SpriteDestination
	SubRect     *image.Rectangle
	SubImage    *ebiten.Image
}

func NewSprite(spriteName string, destination *SpriteDestination, options *ebiten.DrawImageOptions) *Sprite {
	src := AssetLoader.GetAsset(spriteName).(*ebiten.Image)
	rect := image.Rect(
		destination.X,
		destination.Y,
		destination.X+destination.SizeX,
		destination.Y+destination.SizeY,
	)
	subImg := src.SubImage(rect).(*ebiten.Image)

	fmt.Println(subImg)

	return &Sprite{
		Parent:    nil,
		Options:   options,
		Transform: &vector.Transform{},
		SubRect:   &rect,
		SubImage:  subImg,
	}
}

func (s *Sprite) SetParent(parent *core.GameObject) {
	s.Parent = parent
}

func (s *Sprite) Init() {
	if s.Parent != nil {
		s.Options.GeoM.Translate(float64(s.Parent.Transform.Position.X), float64(s.Parent.Transform.Position.Y))
	} else {
		s.Options.GeoM.Translate(float64(s.Transform.Position.X), float64(s.Transform.Position.Y))
	}
}

func (s *Sprite) Update() {
}

func (s *Sprite) Render(screen *ebiten.Image) {
	screen.DrawImage(s.SubImage, s.Options)
}
