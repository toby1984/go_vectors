package vector2

import (
	"fmt"
	"math"
)

type Vector2 struct {
	X float32
	Y float32
}

/*
 * Member functions.
 */

func (v Vector2) Len() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y)))
}

func (v Vector2) String() string {
	return fmt.Sprintf("(%f,%f)", v.X, v.Y)
}

func (v Vector2) DistanceTo(other Vector2) float32 {
	return float32(math.Sqrt(float64(v.DistanceToSqrt(other.X, other.Y))))
}

func (v Vector2) DistanceTo2(x float32, y float32) float32 {
	return float32(math.Sqrt(float64(v.DistanceToSqrt(x, y))))
}

func (v Vector2) DistanceToSqrt(x float32, y float32) float32 {
	dx := v.X - x
	dy := v.Y - y
	return dx*dx + dy*dy
}

func (v *Vector2) Set(x float32, y float32) {
	v.X = x
	v.Y = y
}

func (v Vector2) Add(v2 Vector2) Vector2 {
	return Vector2{
		X: v.X + v2.X,
		Y: v.Y + v2.Y,
	}
}

func Copy(v Vector2) Vector2 {
	return Vector2{v.X, v.Y}
}

/*
 * Global functions
 */

func NewVector2(x float32, y float32) Vector2 {
	return Vector2{x, y}
}

func NewVector2Copy(v Vector2) Vector2 {
	return Vector2{v.X, v.Y}
}

func Origin() Vector2 {
	return Vector2{0, 0}
}

func (v Vector2) Rotate90DegreesCCW() Vector2 {
	// The right-hand normal of vector (x, y) is (y, -x), and the left-hand normal is (-y, x)
	return Vector2{-v.Y, v.X}
}

func (v Vector2) Limit(maxValue float32) Vector2 {
	speed := v.Len()
	if speed < maxValue {
		return v
	}
	return v.Normalize().Multiply(maxValue)
}

func (v Vector2) Normalize() Vector2 {

	l := v.Len()
	if math.Abs(float64(l)) < 0.00001 {
		return v
	}
	return Vector2{v.X / l, v.Y / l}
}

func (v Vector2) Multiply(factor float32) Vector2 {
	return Vector2{factor * v.X, factor * v.Y}
}

func (v Vector2) WrapIfNecessary(maxValue float32) Vector2 {
	newX := v.X
	newY := v.Y

	if newX < 0 {
		for newX < 0 {
			newX += maxValue
		}
	} else if newX >= maxValue {
		for newX >= maxValue {
			newX -= maxValue
		}
	}

	if newY < 0 {
		for newY < 0 {
			newY += maxValue
		}
	} else if newY >= maxValue {
		for newY >= maxValue {
			newY -= maxValue
		}
	}

	if newX != v.X || newY != v.Y {
		return Vector2{newX, newY}
	}
	return v
}

func (v Vector2) Minus(other Vector2) Vector2 {
	return Vector2{v.X - other.X, v.Y - other.Y}
}

func (v Vector2) Minus2(x float32, y float32) Vector2 {
	return Vector2{v.X - x, v.Y - y}
}
