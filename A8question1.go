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
		defer wg.Done()
		mutex.Lock()
		n++
		mutex.Unlock()
		
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		mutex.Lock()
		if isEven(n) {
			fmt.Println(n, " is even")
		} else {
			fmt.Println(n, " is odd")
		}
		mutex.Unlock()
		
	}()
	wg.Wait()

	fmt.Println("Main completed")
}


