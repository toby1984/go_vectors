package vector3

import (
	"fmt"
	"github.com/toby1984/go_vectors/vector2"
	"math"
)

type Vector3 struct {
	vector2.Vector2
	Z float32
}

func (v Vector3) Len() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v Vector3) String() string {
	return fmt.Sprintf("(%f,%f,%f)", v.X, v.Y, v.Z)
}

func (v Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{
		Vector2: vector2.Vector2{
			X: v.X + v2.X,
			Y: v.Y + v2.Y,
		},
		Z: v.Z + v2.Z,
	}
}
