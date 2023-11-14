// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

func foo(wg *sync.WaitGroup) {
	fmt.Printf("new routine")
	wg.Done()
}

func mainWait() {
	var wg sync.WaitGroup
	wg.Add(1)
	go foo(&wg)
	wg.Wait()
	fmt.Println("Hello, 世界")
}
