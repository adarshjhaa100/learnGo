package goroutines

import (
	"fmt"
	"time"
	"os"
)

/*
	Go routines share the asme memory space, so different threads accessing the same resource should be synced
*/

func runSimpleFunc(s string) {
	for {
		// wait 1s and print s
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("\nstring = %v, time = %v, pid = %v, %v\n", 
		s, time.Now(), os.Getpid(), os.Getppid())

	}

}




type ChanTp struct {
	val int
	threadName string
}

func channelThreadCall(arr []int, thname string, chanl chan ChanTp) {
	
	sum := 0 
	// fmt.Println(arr)

	for _, val := range arr {
		sum += val 
	}

	fmt.Printf("\nFor thread = %v, sum = %v\n", thname, sum)

	chanl <- ChanTp {sum, thname}
}


// channels: a pipe/coduit(through which something flows) where we can send/receive values
/*
	IMPORTANT:
	For unbuffered channel

	S -> C -> R

	S, R should be running at the same time, as there's no buffer in a channel to hold value, it just acts as a pipe
*/

func DemoUnbufChannel() {

	arr := []int {10,20,-3, 4,45,-20}
	
	chanl := make(chan ChanTp) // unbuffered, so will block until there's a receiver

	fmt.Printf("\nChannel desc = %#v\n", chanl)

	go channelThreadCall(arr[:len(arr)/2], "thread1", chanl)
	go channelThreadCall(arr[len(arr)/2:], "thread2", chanl)

	
	// this line will block the channel, and since there's no receiver waiting to receive data(the receive code should actually be running) 
	// in parallel will result in deadlock

	// chanl <- ChanTp{10, "main"} 
	

	x := <-chanl// receive from channel (here, channel is empty)
	y := <-chanl

	// go channelThreadCall(arr[len(arr)/2:], "thread3", chanl)
	
	fmt.Println(x, y, x.val + y.val )
}

/*
	For buffered channel. 
	sends are blocked when buffer is full and receive is blocked when buffer empty
	e.g. C[1] with 1 buffer size, S(v) sending value v

	S(v) -> C // channel will hold 1 value from S, will not block, C = [v](val sent by c)
	S(v1) -> C(v) -> R(v1) // Now channel will block as we dont have empty space,   
	C -> R(v) // Channel will allow receive without sender, C = [], empty
	C -> R // Receive blocked, will require sender
*/
func DemoBufChannel() {}




func Demo() {
	
	go runSimpleFunc("World") // spawns a new thread
	go runSimpleFunc("New")

	runSimpleFunc("hello")

}