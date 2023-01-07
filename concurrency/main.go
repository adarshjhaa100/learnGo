package main

import (
	"fmt"
	goroutines "learn/concurrency/goroutines"

)

func main() {
	fmt.Printf("\nGoroutines\n")
	// goroutines.Demo() // goroutine demo
	// goroutines.DemoUnbufChannel()
	// goroutines.DemoBufChannel()
	// goroutines.ChnlRange()
	// goroutines.ChnlSelect()
	// goroutines.Soln1()
	// goroutines.Soln2()
	// goroutines.MutexExample()
	goroutines.RunSolnWebCrawler()
}