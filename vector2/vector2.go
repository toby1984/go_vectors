package vector2

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X float32
	Y float32
}

func (v *Vector2) Len() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v *Vector2) ToString() string {
	return fmt.Sprintf("(%f,%f)", v.X, v.Y)
}

func (v *Vector2) Distance(x float32, y float32) float32 {
	return float32(math.Sqrt(float64(v.DistanceSqrt(x, y))))
}

func (v *Vector2) DistanceSqrt(x float32, y float32) float32 {
	dx := v.X - x
	dy := v.Y - y
	return dx*dx + dy*dy
}

func (v *Vector2) Add(v2 *Vector2) *Vector2 {
	return &Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}
