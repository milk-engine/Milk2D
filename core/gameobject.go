package core

import (
	"log"
	"reflect"

	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/milk-engine/Milk2D/core/vector"
)

type IGameObject interface {
	Init()
	Update()
	Render(screen *ebiten.Image)
	SetParent(parent *GameObject)
	AddComponent(component IComponent)
	RemoveComponent(component IComponent)
	RemoveComponents(component IComponent)
	GetComponent(component IComponent)
	GetComponents(component IComponent)
	AddToScene(sceneName string)
	Destroy()
}

type GameObject struct {
	IGameObject
	Id         uuid.UUID
	Parent     *GameObject
	Name       string
	Transform  *vector.Transform
	Active     bool
	Components []IComponent
	SortIndex  uint8
}

func NewGameObject(name string, startPosition *vector.Vector2) *GameObject {
	randomId, err := uuid.NewRandom()

	if err != nil {
		log.Fatal(err)
	}

	transform := vector.Transform{
		Position: *startPosition,
		Scale: *vector.One(),
		Rotation: 0.0,
	}
	return &GameObject{
		Id:         randomId,
		Parent:     nil,
		Name:       name,
		Transform:  &transform,
		Active:     true,
		Components: make([]IComponent, 0),
		SortIndex:  1,
	}
}

func (g *GameObject) Init() {
	for _, c := range g.Components {
		c.Init()
	}
}

func (g *GameObject) Update() {
	if g.Active {
		for _, c := range g.Components {
			c.Update()
		}
	}
}

func (g *GameObject) Render(screen *ebiten.Image) {
	if g.Active {
		for _, c := range g.Components {
			c.Render(screen)
		}
	}
}

func (g *GameObject) AddComponent(c IComponent) {
	c.SetParent(g)
	g.Components = append(g.Components, c)
}

func (g *GameObject) RemoveComponent(c IComponent) {
	for i, comp := range g.Components {
		if reflect.TypeOf(comp) == reflect.TypeOf(c) {
			g.Components = append(g.Components[:i], g.Components[i+1:]...)
			return
		}
	}
}
func (g *GameObject) RemoveComponents(c IComponent) {
	for i, comp := range g.Components {
		if reflect.TypeOf(comp) == reflect.TypeOf(c) {
			g.Components = append(g.Components[:i], g.Components[i+1:]...)
		}
	}
}

func (g *GameObject) GetComponent(c IComponent) IComponent {
	for _, comp := range g.Components {
		if reflect.TypeOf(comp) == reflect.TypeOf(c) {
			return comp
		}
	}
	return nil
}

func (g *GameObject) GetComponents(c IComponent) []IComponent {
	var result []IComponent
	for _, comp := range g.Components {
		if reflect.TypeOf(comp) == reflect.TypeOf(c) {
			result = append(result, comp)
		}
	}
	return result
}

func (g *GameObject) Destroy() {
	g.Active = false
}

func (g *GameObject) AddToScene(sceneName string) {
	SceneManager.AddGameObject(sceneName, g)
}
