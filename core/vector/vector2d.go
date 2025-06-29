package vector

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X, Y float32
}

func Zero() *Vector2{
	return &Vector2{
		X: 0,
		Y: 0,
	}
}

func One() *Vector2{
	return &Vector2{
		X: 1,
		Y: 1,
	}
}

func (v *Vector2) Set(x, y float32) {
	v.X = x
	v.Y = y
}

func (v *Vector2) Add(nv Vector2) {
	v.X += nv.X
	v.Y += nv.Y
}

func (v *Vector2) Subtract(nv Vector2) {
	v.X -= nv.X
	v.Y -= nv.Y
}

func (v *Vector2) Multiply(scalar float32) {
	v.X *= scalar
	v.Y *= scalar
}

func (v *Vector2) Divide(scalar float32) {
	v.X /= scalar
	v.Y /= scalar
}

func (v *Vector2) Negate() {
	v.X = -v.X
	v.Y = -v.Y
}

func (v *Vector2) Normalize() {
	length := float32(math.Hypot(float64(v.X), float64(v.Y)))
	if length != 0 {
		v.X /= length
		v.Y /= length
	}
}

func (v *Vector2) Dot(other Vector2) float32 {
	return v.X*other.X + v.Y*other.Y
}

func (v *Vector2) Clamp(max float32) {
	lengthSq := v.LengthSq()
	maxSq := max * max
	if lengthSq > maxSq {
		length := float32(math.Sqrt(float64(lengthSq)))
		v.X = (v.X / length) * max
		v.Y = (v.Y / length) * max
	}
}

func (v *Vector2) Lerp(target Vector2, t float32) {
	v.X += (target.X - v.X) * t
	v.Y += (target.Y - v.Y) * t
}
func (v *Vector2) AngleTo(other Vector2) float32 {
	dot := v.Dot(other)
	lenProduct := v.Length() * other.Length()
	if lenProduct == 0 {
		return 0
	}

	cosAngle := dot / lenProduct
	if cosAngle > 1 {
		cosAngle = 1
	} else if cosAngle < -1 {
		cosAngle = -1
	}
	return float32(math.Acos(float64(cosAngle)))
}

func (v *Vector2) DistanceTo(other Vector2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return float32(math.Hypot(float64(dx), float64(dy)))
}

func (v *Vector2) DistanceToSq(other Vector2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return dx*dx + dy*dy
}

func (v *Vector2) IsZero() bool {
	return v.X == 0 && v.Y == 0
}

func (v *Vector2) Equals(other Vector2, epsilon float32) bool {
	dx := v.X - other.X
	dy := v.Y - other.Y
	return (dx*dx + dy*dy) < epsilon*epsilon
}

func (v *Vector2) SetLength(length float32) {
	v.Normalize()
	v.Multiply(length)
}

func (v *Vector2) Length() float32 {
	return float32(math.Hypot(float64(v.X), float64(v.Y)))
}

func (v *Vector2) LengthSq() float32 {
	return v.X*v.X + v.Y*v.Y
}

func (v *Vector2) String() string {
	return fmt.Sprintf("(%f %f)", v.X, v.Y)
}
