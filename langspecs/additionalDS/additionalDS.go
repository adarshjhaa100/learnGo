package additionalDS

import (
	"fmt"
	"reflect"
	"strings"
)

// declaring a struct type with couple of integers
type Vertex struct {
	X int
	Y int
}

func TestPkg() {

	i := []int{12, 2, 2}
	p := &i

	fmt.Println("hello test")
	fmt.Println(reflect.TypeOf(p), *p) // using reflect to get type of variables

}

func PointerType() {
	intvar := 65
	floatvar := float64(21)
	stringvar := "hello"
	testvar := string(intvar) // ocnvert int to string

	var ptr2 *float64 // declare a ptr to float64
	ptr2 = &floatvar  // initialize ptr

	ptr := &intvar // declare and init pointer on same line

	var ptr3 *int

	// fmt.Println(intvar)
	// fmt.Println(floatvar)
	fmt.Println(stringvar)
	fmt.Println(testvar)

	fmt.Println(*ptr) //dereference
	fmt.Println(*ptr2 + 1.0)

	fmt.Printf("%p %v\n", ptr3, ptr3) //out: 0x0 nil // dereferencing will give a segmentation fault due to nil pointer reference
}

func StructTest() {
	v1 := Vertex{10, 20}
	v2 := Vertex{X: 10} // Y = 0
	v3 := Vertex{}      // X, Y both 0

	pstr := &v1 // pointer to struct
	pstr.X = 56 // we can write (*pstr).X, but language permits using direct variable names

	fmt.Println(Vertex{10, 20}) // struct literal
	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(pstr)

}

func ArrayTypes() {
	// arrays are declared with type of the form []T
	arr := [10]int{12, 3, 4, 45, 5, 5, 5, 5} // other indices will have default type values, here: 0
	arr[8] = 10
	arr[9] = 56

	arr2d := [10][10]int{{1, 2}} // 2d array declaration and initialization, everyting except arr2d[0][0] and arr2d[0][1] will be default value, here: 0

	fmt.Println(arr)

	fmt.Println("2d array: ", arr2d)

	fmt.Println("Access 2d array element: ", arr2d[0][1])

}

func Slices() {
	sliceInit := [10]int{} // declare an empty slice

	fmt.Println(sliceInit)

	ptrsliceInit := &sliceInit

	fmt.Printf("example ofptr to slice: %p %v\n", ptrsliceInit, ptrsliceInit[1])
	
	// slices are declared using []T (without specifying size), their size can be dynamically increased upto a specific capacity
	// arr[start:end] (from start to end-1) (can omit either of start/end)
	slice1 := sliceInit[1:8]
	slice1[2] = 34 // slices are references to the array items 


	fmt.Println(sliceInit)
	fmt.Println(slice1)

	// slice has both length and capacity (len() and cap() can be used oh other types with theire respective meanings )
	fmt.Println(len(slice1), cap(slice1))

	// Slice lengths can be modified upto the capacity

	slice1 = slice1[:2] // cpmpressing size: len(slice1) = 2

	fmt.Println(len(slice1), cap(slice1))

	slice1 = slice1[:8] // expanding: len = 8

	fmt.Println(len(slice1), cap(slice1))

	slice1 = slice1[2:] // this will decrease the capacity and drop first 2 elements

	fmt.Println(len(slice1), cap(slice1)) 


	var slice2 []int // empty slice(value == nil is true)

	fmt.Println(len(slice2), cap(slice2), slice2, slice2 == nil)

	//create slice with make 
	// make allocates memory for maps, slice an channels
	// first argument is a type
	
	slice3 := make([]string, 5, 10) // make string slice with len = 5, cap = 10
	slice3[4] = "dskfoldks"
	slice3[1] = "23"

	fmt.Printf("%v \n", slice3) //%v to print slice value

	slice2dfake := [][] int {{}}; // slice with 0 elements, useless
	slice2d := [][] int {
		[]int {2, 33},
		[]int {1},
		[]int {3},
	}; 
	

	fmt.Println("fake 2d slice", slice2dfake, len(slice2dfake), cap(slice2dfake))

	fmt.Println("realest 2d slice", slice2d, len(slice2d), cap(slice2d), slice2d[1])



	// append to slice
	fmt.Println("Append to slice : ")
	slExpandable := []int32 {}
	for i := int32(0); i<10; i++ {
		slExpandable = append(slExpandable, i) // append to slice, return resliced result if sufficient size, else assign a new array with expanded size
		fmt.Println( slExpandable, len(slExpandable), cap(slExpandable) )
	}

	// the capacity of slice increases exponentially while appending if size is not sufficient to accomodate a new element(multiple of 2)
	/*
	[0] 1 2
	[0 1] 2 2
	[0 1 2] 3 4
	[0 1 2 3] 4 4
	[0 1 2 3 4] 5 8
	[0 1 2 3 4 5] 6 8
	[0 1 2 3 4 5 6] 7 8
	[0 1 2 3 4 5 6 7] 8 8
	[0 1 2 3 4 5 6 7 8] 9 16
	[0 1 2 3 4 5 6 7 8 9] 10 16
	*/

	// slice of custom struct type
	sliceStr := []Vertex {}
	for i := 1; i<=5; i++ {
		sliceStr = append(sliceStr, Vertex{i*2, i*3})
	}

	//  we can append more than one element at a time
	sliceStr = append(sliceStr, Vertex{}, Vertex{}, Vertex{})

	fmt.Println(sliceStr)


}

func SliceExercise() {
// 	package main

// import "golang.org/x/tour/pic"

// func Pic(dx, dy int) [][]uint8 {
// 	img := make([]uint8, dx, dx)
// 	var result [][]uint8
	
// 	for i := 0; i<dy; i++ {
// 		for j := 0; j<dx; j++ {
// 			img[j] = uint8((i*j)/2)
// 		}
// 		result = append(result, img)		
// 	}
// 	return result
// }

// func main() {
// 	pic.Show(Pic)
// }
}


func RangeTest() {
	// for item in arr
	arr := [5]int {1,2,3,45,5}

	for index, val := range arr {
		fmt.Println(index, val)
	}

	for _, val := range arr {
		fmt.Println(val)
	}

	// either index or val can be skipped using _
}



type VertexPt struct {
	Lat, Long float64
}

// redeclaration of types (similar to typedef in c)
type Int64 int


func MapTest() {
	
	// Initializing a map of key type string and val type VertexPt, make will return an initialized map
	mp := make(map[string]VertexPt) 

	// map literals (need to add key and values) 
	// (Here, we didn't need to add types to vertex as they've already been given in the var type)
	var mp3 = map[string]VertexPt{
		"A": {1,2,},
		"B": {3,4,},
	}

	//map of map (the parenthesis not required in the type, just for readability)
	var mpOmp = make(map[string](map[string]VertexPt))
	mpOmp["loopOloop"] = mp3


	// or (this is a nil map which has no key)
	var mp1 map[string]Vertex;
	
	// this assignment is illegal as we're assigning to a nil map
	// mp1["Ru"] = Vertex{
	// 	1,
	// 	1,
	// }


	mp["Delhi"] = VertexPt {
		2.3, 
		4.5, 
	}
	mp["mumbai"] = VertexPt{
		1.1,
		1.2,
	}


	fmt.Println(mp)
	fmt.Println(mp["Delhi"])
	fmt.Println("Trying to access unassigned key",mp["fsdifhjdsf"]) // accessing a non assigned key will return the default value of value type, here {0 0}
	fmt.Println("NIL map: ", mp1)
	fmt.Println("Size of map mp: ", len(mp))
	fmt.Println("Map initialized using a literal", mp3)
	fmt.Println("map of map: ", mpOmp, mpOmp["loopOloop"])


	// mutating maps
	// adding new key
	mp["Kanpur"] = VertexPt{2,4,} 

	//modifying a key
	mp["Delhi"] = VertexPt{6,7,}

	// removing a key (only if key is present, else its a no op)
	delete(mp, "Kanpur")

	// searching for a key, ok is false if nothing returned (if only one var on lhs, type is the value type else type is "value type, bool")
	element, ok := mp["cawnpore"]

	fmt.Println(mp)
	fmt.Println(element, ok)

}

// Problem: Calculate Word Count in a paragraph using maps
func WordCountMap() {
	var strQuestion string =  "hello we friend, things have been great since we last met my friend. Hopefully we get Better."

	

	// warning, this would look like a single string but its not, its actually split
	// split by any no of whitespace chars in between
	// fmt.Println(strings.Fields(strQuestion))

	// format specifier (# ensures that output is human readable and not like [1 1 3 4] which can be confusing in case of string)
	fmt.Printf("%#v\n", strings.Fields(strQuestion))
	// fmt.Printf("%#v\n", Vertex{1,2,})

	mpAns := make(map[string]int)
	var wordList []string = strings.Fields(strQuestion)
	
	for _,val := range wordList {
		if _, keyPresent := mpAns[val]; !keyPresent {
			mpAns[val] = 1
		} else {
			mpAns[val]+=1
		}	
	}

	fmt.Printf("Word count Map: %#v\n", mpAns);

}

// functions can be passes to another function
func PassFunction() {

	funcFirst := func (a int, b float32) float64{	
			return float64(float64(a)+float64(b))
	}
	/*
		f func (int, float32) float64 :
		
		f is the name of function, followed by the type "(int, float32) float64"
	*/
	funcSecond := func ( add func (int, float32) float64, n []int ) float64 {
		sum := float64(0)
		for index, val := range n {
			sum += add(index, float32(val))
		}
		return sum
	}	

	fmt.Printf("First function: %+v\n", funcFirst(1,2))	

	fmt.Printf("Second calling first: %+v\n", funcSecond(funcFirst, []int{10,20,30,40,}))	
	
}

// closures: binding function to variables declared outside scope
func ClosureTest() {
	bindedVar := 0
	calculateSumArr := func( list []int ) {
		for _, val := range list {
			bindedVar += val 
		}

		// return bindedVar
	}

	calculateSumArr([]int{1,2,3,4,})
	// the bindedVar was modified by closureTest
	fmt.Println(bindedVar)
	

}
