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
	//MySlice := [] int {2,4,6,8,10,12,14,16,18,20}
	var index int
	fmt.Println("Enter Input: ")
	fmt.Scanln(&index)
	found:=false
	for i:=0;i<len(MySlice);i++ {
		if index == i {
			fmt.Printf("item %v, value %v ",i,MySlice[i])
			found =true
			break
		}
	}
	if !found {
		defer func(){
			if r:=recover(); r!=nil {
				fmt.Printf("internal error: %v\n",r)
			}
		} ()
		panic("Index out of bound")
		
	}
	
}

/*

sharanya@sharanya-ASUS-TUF-Gaming-A15-FA506IHRB-FA506IHRZ:~/Desktop/MyGoPracs$ go run A6question1.go
Enter the size of the slice:
5
Enter 5 elements for the slice:
Element 0: 10
Element 1: 20
Element 2: 30
Element 3: 40
Element 4: 50
Enter Input: 
9
internal error: Index out of bound
Testing Panic and Recover



*/


