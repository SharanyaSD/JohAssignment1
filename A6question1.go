
//Sharanya Datrange
package main


import (
	"fmt"
)

func main() {
	MySlice := SliceInput()
	accessSlice(MySlice)
	fmt.Println("Testing Panic and Recover")
}

func SliceInput() []int {
	var size int
	fmt.Println("Enter the size of the slice:")
	fmt.Scanln(&size)
	MySlice := make([]int, size)
	fmt.Printf("Enter %d elements for the slice:\n", size)
	for i := 0; i < size; i++ {
		fmt.Printf("Element %d: ", i)
		fmt.Scanln(&MySlice[i])
	}
	return MySlice
}

func accessSlice(MySlice []int) {
	var index int
	fmt.Println("Enter Input: ")
	fmt.Scanln(&index)
	defer func(){
		if r:=recover(); r!=nil {
			fmt.Printf("internal error: %v\n",r)
		}
	} ()
	fmt.Printf("item %v, value %v ",index,MySlice[index])
	
}

/*

sharanya@sharanya-ASUS-TUF-Gaming-A15-FA506IHRB-FA506IHRZ:~/Desktop/MyGoPracs$ go run A6question1.go
Enter the size of the slice:
3
Enter 3 elements for the slice:
Element 0: 1
Element 1: 2
Element 2: 3
Enter Input: 
9
internal error: runtime error: index out of range [9] with length 3
Testing Panic and Recover
sharanya@sharanya-ASUS-TUF-Gaming



*/



