//Sharanya Datrange

package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	n := 0
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		mutex.Lock()
		n++
		mutex.Unlock()
		defer wg.Done()
	}()

	wg.Add(1)
	go func() {
		mutex.Lock()
		if isEven(n) {
			fmt.Println(n, " is even")
		} else {
			fmt.Println(n, " is odd")
		}
		mutex.Unlock()
		defer wg.Done()
	}()
	wg.Wait()

	fmt.Println("Main completed")
}
