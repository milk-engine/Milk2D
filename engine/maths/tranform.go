package maths

import "math"

type Transform struct {
	position Vector2
	scale    Vector2
	rotation float32
}

func NewTransform() *Transform {
	return &Transform{
		position: Vector2{0, 0},
		scale:    Vector2{1, 1},
		rotation: 0,
	}
}

func (t *Transform) SetPosition(x, y float32) {
	t.position.Set(x, y)
}

func (t *Transform) SetScale(sx, sy float32) {
	t.scale.Set(sx, sy)
}

func (t *Transform) SetRotation(deg float32) {
	t.rotation = deg
}

func (t *Transform) Translate(destination Vector2) {
	t.position.Add(destination)
}

func (t *Transform) Rotate(deltaDeg float32) {
	t.rotation += deltaDeg

	if t.rotation >= 360 {
		t.rotation -= 360
	} else if t.rotation < 0 {
		t.rotation += 360
	}
}

func degreesToRadians(deg float32) float32 {
	return (deg * math.Pi) / 180
}

func (t *Transform) RotationRadians() float32 {
	return degreesToRadians(t.rotation)
}

func (t *Transform) Apply(v Vector2) Vector2 {
	x := v.x * t.scale.x
	y := v.y * t.scale.y

	rad := degreesToRadians(t.rotation)
	cos := float32(math.Cos(float64(rad)))
	sin := float32(math.Sin(float64(rad)))

	rx := x*cos - y*sin
	ry := x*sin + y*cos

	rx += t.position.x
	ry += t.position.y

	return Vector2{rx, ry}
}

func (t *Transform) Reset() {
	t.position.Set(0, 0)
	t.scale.Set(1, 1)
	t.rotation = 0
}
