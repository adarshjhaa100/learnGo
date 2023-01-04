package goroutines

/*
	Binary Tree in Golang
	Exercise: Given two trees, find if they're equivalent using concurrency mechanism
*/
import (
	"golang.org/x/tour/tree"
	"fmt"
	"time"
)


// type Tree struct {
// 	left, right *Tree
// 	val int
// }


// Two approaches to call walk(according to me)
/*
	1. For recursive approach, 
		we dont know when the gorouting calling walk will stop sending values. 
		Note: walk wont stop on nil!!! 
		So,  we call WalK() using anonymous goroutine after which the channel will be closed 
		use range to read values
	
		2. For iterative approach, 
		when using a stack based approach, we can send values until none remain in stack and then close channel 


*/

func Walk(node *tree.Tree, ch chan int) {
	// fmt.Println("walk: ", node)
	
	if(node == nil) {
		return
	}

	Walk(node.Left, ch)
	ch <- node.Value
	Walk(node.Right, ch)
	
	
}


func Same(t1, t2 *tree.Tree) bool{
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()
	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()
	
	Loop: // adding label in for loop (we can break out of the loop using an inner statement like break Labelname)
	for {
		select {
			case x, ok1 := <-ch1:
				y, ok2 := <-ch2

				if(ok1 && ok2) {
					if(x != y) {
						return false
					}
					fmt.Println("x, y =", x, y, time.Now())	
				} else {
					break Loop
				}

			case y, ok2 := <-ch2:
				x, ok1 := <-ch1
				if(ok1 && ok2) {
					if(x!=y) {
						return false
					}

				} else {
					break Loop
				}
		}
	}

	return true

}


func Soln1() {
	// tree.New(k) returns a randomly built binary tree holding k, 2k,...10k
	ch := make(chan int)
	tree1 := tree.New(3)
	fmt.Println(tree1)

	// anomymous function to close the channel, if we try to close in nil
	go func() {
		Walk(tree1, ch)
		fmt.Println("Channel Closed")
		close(ch)
	}()

	// channel need to be closed
	for x := range ch {
		fmt.Println(x)
	}

}

func Soln2() {
	isSame := Same(tree.New(1), tree.New(1))
	// Same(tree.New(1), tree.New(1))
	fmt.Println("Same trees?: ", isSame)

	isSame = Same(tree.New(1), tree.New(2))
	// Same(tree.New(1), tree.New(1))
	
	fmt.Println("Same trees?: ", isSame)


}