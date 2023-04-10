package goroutines

import (
	"fmt"
	"time"
	"os"
)

/*
	Go routines share the same memory space, so different threads accessing the same resource should be synced
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


// func sendTobuf(chn <-chan int) {
// }

func sendToChannel(val int, chanl chan int) {
	fmt.Printf("\nInserting value: %v\n", val)
	// time.Sleep(time.Second)
	chanl <- val
	if(val >= 60) {
		// Sender can close a channel
		close(chanl)
	}
	fmt.Printf("\nInserting value success: %v\n", val)
}


func DemoBufChannel() {
	
	// Values are inserted and removed in FIFO order
	chnlbuf := make(chan int, 2) // can store upto 2 values in buffer, any addition value sent required receiver to be active for receiving

	// to get the length of buffered channel, len(channel), bo of elements stored in channel

	// inline function without explicit function declaration
	go func() {
		i := 0
		for {
			// time.Sleep(time.Second)
			
			fmt.Printf("\nreceiving index = %v, chnl = %#v\n", i, len(chnlbuf))
			
			// so as to not read from empty channel
			if len(chnlbuf) > 0 {
				// isOpen is true until channel is Open
				val, isOpen :=<- chnlbuf
				fmt.Printf("\nreceiving index=%v, value=%#v, open? = %v\n", i, val, isOpen)
			i++ } else {
				i++
				fmt.Println("Channel empty")
				continue
			}

		}
	}()

	sendToChannel(10, chnlbuf) // success, buffer: 1/2, non blobking
	sendToChannel(20, chnlbuf)  // success, buffer: 2/2, blocking
	sendToChannel(30, chnlbuf)  // overflow, requires a receiver running in a parallel goroutine
	sendToChannel(40, chnlbuf) 
	sendToChannel(50, chnlbuf) 
	sendToChannel(60, chnlbuf) 	


	//send:  
	// 10,20, 30
	// chanl: 
}


func sqrtChnl(chnl chan int) {
	val := 0
	for {
		val += 1
		chnl <- val * val
		
		if(val>100) {
			close(chnl) // its not necessary to close the chnl until we use something like range
			fmt.Println("Closing Channel ... ")
			break
		}
	}
}

func ChnlRange() {
	chanl := make(chan int)
	
	go sqrtChnl(chanl)
	// channel range, will keep receiving until channel is closed
	for val := range chanl {
		fmt.Println(val)
	}

}

/*
	Select blocks until one of the cases can run, then it executes the case.
	In case multiple cases can run, it chooses them at random.

	This is pretty useful while receiving from mutiple channels,
		In an iterative implementation: 
			val1 <- chan1; val2 <- chan2
		statement 2 waits for the first statement to execute, whats the point of concurrency then

		instead if we do this:
		select {
			case x := <-chan1: statement
			case y := <-chan2: statement
		}

*/

func ChnlSelect() {
	chanel := make(chan int)
	quit := make(chan int)

	go func() {
		x, y := 0, 1
		for {
			select {
				case chanel <- x:
					x, y = y, x+y
				case t := <-chanel:
					fmt.Println(t)
				case <-quit:
					fmt.Println("quit...")
				default: // works when no other case runs
					fmt.Println("Nothing to receive from chanels...")
					continue
			}
		}
	}()
	
	
		// run till i=9, then quit, channel select will select the first case of chanel<-x until we can consume, 
		// then close. This way we can close the connection from receiver
		for i:=0; i<10; i++ {
			fmt.Println(<-chanel)
		}
		quit<-0 

}



func Demo() {
	
	go runSimpleFunc("World") // spawns a new thread
	go runSimpleFunc("New")

	runSimpleFunc("hello")

}