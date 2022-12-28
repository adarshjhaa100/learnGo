package generictype

import "fmt"

// List represents a singly linked list that holds value of anly type
type ListNode[T comparable] struct {
	val T // value of linked list
	next *ListNode[T] // pointer to next node
}

type List[T comparable] struct{
	head *ListNode[T] // head pointer
	length uint32 // number of nodes
}


func (lst *List[T]) Traverse() {
	
	tempPtr := lst.head
	fmt.Println()

	if(lst.head == nil) {
		fmt.Printf("\nEmpty\n")
	}

	for tempPtr!=nil {
		fmt.Printf("%+v->", tempPtr.val)
		tempPtr = tempPtr.next
	}		

	fmt.Println()

}


/*
	Insert value into the linked list at the specified posn
	posn = 0 // beginning
	posn =  length+1 // end

	error codes: 
		0: Success
		2: insufficient length
*/
func (lst *List[T]) Insert(value T, posn int32) uint8{
	
	// fmt.Println(uint32(posn), lst.length)

	if(posn > int32(lst.length)+1) {
		fmt.Println("Cannot insert due to length issue")
		return 2
	}

	tmpPtr  := lst.head
	newNode := ListNode[T]{value, nil}  
	
	// update posn if -1
	if(posn == -1) {
		posn = int32(lst.length)
	}
	
	// insert at head
	if(posn == 0 || tmpPtr == nil) {
		newNode.next = tmpPtr
		lst.head = &newNode
		lst.length++

		return 0
	}
	
	ctTraversed := int32(0)

	for tmpPtr!=nil {
		if posn == ctTraversed+1 {
			newNode.next = tmpPtr.next
			tmpPtr.next = &newNode
		}

		tmpPtr = tmpPtr.next
		ctTraversed++
	}

	lst.length++ // increment the length after insertion

	return 0

}

