package core

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Iscene interface {
	Init()
	Update()
	Render(screen ebiten.Image)
	AddGameObject(gameObj *GameObject)
}

type Scene struct {
	Iscene
	Name        string
	GameObjects []*GameObject
	SortIndex   uint8
}

func NewScene(name string) *Scene {
	return &Scene{
		Name:        name,
		GameObjects: make([]*GameObject, 0),
		SortIndex:   1,
	}
}

func (s *Scene) Init() {
	for _, gameObj := range s.GameObjects {
		gameObj.Init()
	}
}
func (s *Scene) Update() {
	for _, gameObj := range s.GameObjects {
		if gameObj.Active {
			gameObj.Update()
		}
	}
}
func (s *Scene) Render(screen *ebiten.Image) {
	for _, gameObj := range s.GameObjects {
		if gameObj.Active {
			gameObj.Render(screen)
		}
	}
}

func (s *Scene) AddGameObject(gameObj *GameObject) {
	s.GameObjects = append(s.GameObjects, gameObj)
}
