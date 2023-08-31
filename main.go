package vectors

import (
	"fmt"
	"github.com/toby1984/go_vectors/vector2"
	"github.com/toby1984/go_vectors/vector3"
)

func main() {

	var v1 Vector[vector2.Vector2]
	var v2 Vector[vector3.Vector3]

	v1 = vector2.Vector2{X: 1, Y: 2}
	v2 = vector3.Vector3{Vector2: vector2.Vector2{X: 1, Y: 2}, Z: 3}
	fmt.Println(v1.String())
	fmt.Println(v1.Add(v1.(vector2.Vector2)).String())
	fmt.Println(v2.String())
}
