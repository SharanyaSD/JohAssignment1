//Sharanya Datrange

package main

import (
	"fmt"
	"errors"
)

func main() {
	MySlice:=SliceInput()
	err := acceptSlice(MySlice)
	if err!=nil {
		fmt.Println(err)
	}

}

func SliceInput() []int {
	var size int
	fmt.Println("Enter size of slice: ")
	fmt.Scanln(&size)
	MySlice:=make([] int,size)
	fmt.Printf("Enter %d elements for slice ", size)
	for i:=0;i<size;i++ {
		fmt.Printf("element %d : ",i)
		fmt.Scanln(&MySlice[i])
	}
	return MySlice
}

func acceptSlice(MySlice []int) (error) {
	var index int
	fmt.Println("Enter index to search value: ")
	fmt.Scanln(&index)
	
	if index < 0 || index >= len(MySlice) {
		return errors.New("index out of range")
	}

	fmt.Printf("item %v, value %v\n", index, MySlice[index])
	return nil


}

/*
sharanya@sharanya-ASUS-TUF-Gaming-A15-FA506IHRB-FA506IHRZ:~/Desktop/MyGoPracs$ go run A6question2.go
Enter size of slice: 
5
Enter 5 elements for slice element 0 : 1
element 1 : 2
element 2 : 3
element 3 : 4
element 4 : 5
Enter index to search value: 
9
index out of range

*/
