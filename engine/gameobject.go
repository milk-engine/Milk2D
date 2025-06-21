package engine

import (
	"log"

	"github.com/google/uuid"
	"github.com/milk-engine/Milk2D/engine/maths"
)

type GameObject struct {
	Id        uuid.UUID
	Name      string
	Parent    *GameObject
	Transform *maths.Transform
	Active    bool
}

func NewGameObject(name string, transform maths.Transform) *GameObject {
	randomId, err := uuid.NewRandom()

	if err != nil {
		log.Fatal(err)
	}

	return &GameObject{
		Id:        randomId,
		Parent:    nil,
		Name:      name,
		Transform: &transform,
		Active:    true,
	}
}
