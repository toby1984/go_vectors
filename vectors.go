package vectors

import (
	"github.com/toby1984/go_vectors/vector2"
	"github.com/toby1984/go_vectors/vector3"
)

type Vectors interface {
	vector2.Vector2 | vector3.Vector3
}

type Vector[T Vectors] interface {
	Len() float32
	Add(v2 T) T
	String() string
}
