package goroutines

/*
	Mutex locks are used to lock resources(here values) used by goroutines
	which make sure that only one goroutine can access.
	To ensure consistency, reads and write bot should be locked
*/
import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex // Add mutex to the structure
	val map[string]int
}

func (ctr *SafeCounter) IncrementCounter(key string) {
	ctr.mu.Lock() // lock for write
	ctr.val[key]++ // increment value
	ctr.mu.Unlock() // unlock before exit
}


func (ctr *SafeCounter) Value(key string) int{
	ctr.mu.Lock() // lock read
	defer ctr.mu.Unlock() // defer until last statement i.e. return is executed
	return ctr.val[key]
}



func MutexExample() {
	ctr := SafeCounter{val: make(map[string]int)}	
	for i := 0; i <= 1000; i++ {
		go ctr.IncrementCounter("testKey")
		
		go func() {
			fmt.Println(time.Now(), ctr.Value("testKey") )
		}()
	} 

	time.Sleep(time.Second * 1)

	fmt.Println("Counter Value", ctr.Value("testKey"))


}