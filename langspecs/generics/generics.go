package generics

/*
	Useful for to make functions work on multiple types

*/

import (
	"fmt"
	"reflect"
)

type Vertex2d struct {
	x, y float64
}


// Func Index which has a generic type T that has taken comparable type(useful to compare multiple values)
// after defining T, directly use T
func Index[T comparable](arr []T, x T) int{

	fmt.Printf("Type of x is : %v\n", reflect.TypeOf(x))
	
	for i,v := range arr {
		// v and x are of type T
		if v == x {
			return i //return index of the element for which search matcdhes
		}
	}

	return -1
	
}

func GenericsUsage() {
	// generics passed to a function
	var (
		arr1  = []int{10, 20, 30, 40, 50}
		val1  = 30

		arr2  = []string{"This", "is", "a", "string", "array"}
		val2  = "string"

		arr3  = []Vertex2d {
			{2,3},
			{3,4},
			{4,5},
			{5,6},
		} 
		val3  = Vertex2d{4,5}
	)

	// call generic function for int arr
	fmt.Printf("Index of %v in %#v : %v\n", val1, arr1, Index(arr1, val1))
	// call generic function for string arr
	fmt.Printf("Index of %v in %#v : %v\n", val2, arr2, Index(arr2, val2))
	// call generic function for struct arr
	fmt.Printf("Index of %v in %#v : %v\n", val3, arr3, Index(arr3, val3))
	
}

/*
	In addition to functions, generics can also work on multiple types
*/
// Singly Linked List with element of type any

type List[T any] struct {
	val T
	next *List[T]
}

//Methods cannot have generic type parameters
func Insert(head *List[any], value any) {
	// If list is null, make the new value as head
	newValue := List[any] {value, nil}
	current := head
	
	for {
		// fmt.Printf("Type: %#v\n", current)
		
		if(current == nil){
			println("Reached End of Linked List")
			break
		}	

		// Assumptions: Null linked list has head {-1, null}
		if(*current == List[any]{}) {
			current.val = value
			break
		}	
		
		// Next is 
		if(current.next == nil) {
			current.next = &newValue
			break
		}

		current = current.next
		
	}
}

// Traverse linked list:
func Traverse(head *List[any]) {
	count := 0
	ptr := head

	for {
		
		fmt.Printf("Node at posn %v : %#v\n", count, ptr)
		if(ptr == nil){
			break
		}
		count+=1
		ptr = ptr.next
	}
}


func UtilizeLinkedList() {
	// Instantiate the head list (A generic type must be instantiated), 
	// a type must be specified as well. Also, any is alias for interface{}
	// head := List[any]{-1, nil}

	var head List[any] // the value of null head will be List[any]{}
	// megahead = List[any]{-1, nil}
	// head = List[any]{1, nil}
	Insert(&head, 23)
	Insert(&head, 24)
	Insert(&head, 25)
	

	Traverse(&head)
	fmt.Printf("Head Node: %#v", head )

}




