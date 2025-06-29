package vector

import "math"

type Transform struct {
	Position Vector2
	Scale    Vector2
	Rotation float32
}

func (t *Transform) SetPosition(x, y float32) {
	t.Position.Set(x, y)
}

func (t *Transform) SetScale(sx, sy float32) {
	t.Scale.Set(sx, sy)
}

func (t *Transform) SetRotation(deg float32) {
	t.Rotation = deg
}

func (t *Transform) Translate(destination Vector2) {
	t.Position.Add(destination)
}

func (t *Transform) Rotate(deltaDeg float32) {
	t.Rotation += deltaDeg

	if t.Rotation >= 360 {
		t.Rotation -= 360
	} else if t.Rotation < 0 {
		t.Rotation += 360
	}
}

func degreesToRadians(deg float32) float32 {
	return (deg * math.Pi) / 180
}

func (t *Transform) RotationRadians() float32 {
	return degreesToRadians(t.Rotation)
}

func (t *Transform) Apply(v Vector2) Vector2 {
	x := v.X * t.Scale.X
	y := v.Y * t.Scale.Y

	rad := degreesToRadians(t.Rotation)
	cos := float32(math.Cos(float64(rad)))
	sin := float32(math.Sin(float64(rad)))

	rx := x*cos - y*sin
	ry := x*sin + y*cos

	rx += t.Position.X
	ry += t.Position.Y

	return Vector2{rx, ry}
}

func (t *Transform) Reset() {
	t.Position.Set(0, 0)
	t.Scale.Set(1, 1)
	t.Rotation = 0
}
