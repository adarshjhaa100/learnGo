package methods

import (
	"fmt"
	"math"
)


type Vertex3D struct {
	x float64
	y float64
	z float64
}


type INT64 int64


func SquaredDiff(v1 Vertex3D, v2 Vertex3D) float64{

	return (v1.x-v2.x)*(v1.x-v2.x) + (v1.y-v2.y)*(v1.y-v2.y) + (v1.z-v2.z)*(v1.z-v2.z)
}

// define a method for the type specified (parameter specified before the function name)
// method is just a function with a receiver argument
func (v Vertex3D) SquaredDistMethod(v1 Vertex3D) float64 {
	return (v1.x-v.x)*(v1.x-v.x) + (v1.y-v.y)*(v1.y-v.y) + (v1.z-v.z)*(v1.z-v.z)
}

// methods can also be declared on non struct types
type FLOT float64
func (f FLOT) Add(f1 float64) FLOT{
	return f + FLOT(f1)
}

// pointer receivers (modify the actual object instead or returning a copy)
func (f *FLOT) Scale(s float64) {
	*f = *f * FLOT(s)
} 

// pointer receiver: In case of struct vars, no need to explicitly dereference
func (v *Vertex3D) Scale(s float64) {
	// Golang implicitly converts the below statements to: (*v).x = (*)v.x * s
	v.x = v.x * s
	v.y = v.y * s
	v.z = v.z * s
	
}


func TestVertex() {
	v1 := Vertex3D {1, 2, 3}
	v2 := Vertex3D {2, 3, 4}
	
	result := float32(math.Abs(math.Sqrt(SquaredDiff(v1, v2))))

	resultMethod := float32(v1.SquaredDistMethod(v2))



	resultFloat := FLOT(2.4)
	resultFloat  = resultFloat.Add(3.5)

	resultFloatPtr := FLOT(2)
	resultFloatPtr.Scale(5)

	vertex1 := Vertex3D {3, 4, 5}
	// this explicitly changes to (&vertex1).Scale(10)
	vertex1.Scale(2)

	// pointer indirection: if we call a method on reference of a variable, it automatically dereferences
	ps := &vertex1
	// the statement below is implicitly converted to (*ps).MethodName
	fmt.Println("Pointer indirection", ps.SquaredDistMethod(Vertex3D{0, 0, 0}))


	fmt.Println(result)
	fmt.Println("method call", resultMethod)
	fmt.Println("method on non struct :", resultFloat )
	fmt.Println("pointer receiver on non struct: ", resultFloatPtr)
	fmt.Printf("pointer receiver on struct: %#v\n", vertex1)
}



