package maths

import (
	"fmt"
	"math"
)

type Vector2 struct {
	x, y float32
}

func NewVector2(x, y float32) *Vector2 {
	return &Vector2{
		x: x,
		y: y,
	}
}

func (v *Vector2) Set(x, y float32) {
	v.x = x
	v.y = y
}

func (v *Vector2) Add(nv Vector2) {
	v.x += nv.x
	v.y += nv.y
}

func (v *Vector2) Subtract(nv Vector2) {
	v.x -= nv.x
	v.y -= nv.y
}

func (v *Vector2) Multiply(scalar float32) {
	v.x *= scalar
	v.y *= scalar
}

func (v *Vector2) Divide(scalar float32) {
	v.x /= scalar
	v.y /= scalar
}

func (v *Vector2) Negate() {
	v.x = -v.x
	v.y = -v.y
}

func (v *Vector2) Normalize() {
	length := float32(math.Hypot(float64(v.x), float64(v.y)))
	if length != 0 {
		v.x /= length
		v.y /= length
	}
}

func (v *Vector2) Dot(other Vector2) float32 {
	return v.x*other.x + v.y*other.y
}

func (v *Vector2) Clamp(max float32) {
	lengthSq := v.LengthSq()
	maxSq := max * max
	if lengthSq > maxSq {
		length := float32(math.Sqrt(float64(lengthSq)))
		v.x = (v.x / length) * max
		v.y = (v.y / length) * max
	}
}

func (v *Vector2) Lerp(target Vector2, t float32) {
	v.x += (target.x - v.x) * t
	v.y += (target.y - v.y) * t
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
	dx := v.x - other.x
	dy := v.y - other.y
	return float32(math.Hypot(float64(dx), float64(dy)))
}

func (v *Vector2) DistanceToSq(other Vector2) float32 {
	dx := v.x - other.x
	dy := v.y - other.y
	return dx*dx + dy*dy
}

func (v *Vector2) IsZero() bool {
	return v.x == 0 && v.y == 0
}

func (v *Vector2) Equals(other Vector2, epsilon float32) bool {
	dx := v.x - other.x
	dy := v.y - other.y
	return (dx*dx + dy*dy) < epsilon*epsilon
}

func (v *Vector2) SetLength(length float32) {
	v.Normalize()
	v.Multiply(length)
}

func (v *Vector2) Length() float32 {
	return float32(math.Hypot(float64(v.x), float64(v.y)))
}

func (v *Vector2) LengthSq() float32 {
	return v.x*v.x + v.y*v.y
}

func (v *Vector2) String() string {
	return fmt.Sprintf("(%f %f)", v.x, v.y)
}
