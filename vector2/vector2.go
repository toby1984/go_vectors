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

func (v *Vector2) Add(v2 *Vector2) *Vector2 {
	return &Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}
