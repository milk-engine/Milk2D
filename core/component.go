package core

import "github.com/hajimehoshi/ebiten/v2"

type IComponent interface {
	Init()
	Render(screen *ebiten.Image)
	Update()
	SetParent(parent *GameObject)
}

type Component struct {
	IComponent
	Parent *GameObject
}

func (c *Component) Init() {
	//
}

func (c *Component) Render(screen *ebiten.Image) {
	//
}

func (c *Component) Update() {
	//
}

func (c *Component) SetParent(parent *GameObject) {
	c.Parent = parent
}
