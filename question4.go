//Sharanya Datrange
/*
Complete the program to return 3 slices of a given array, based on the following conditions.
Given Array : ["qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"]
Hint: Hard-code the given array into your program.
Input: Two space-separated integers representing index1 and index2.
Output: Output will contain 3 slices
1. slice containing all the elements from the start of the array to Index1
2. slice containing all the elements from index1 to index2
3. slice containing all the elements from index2 to the end of the array
*/
package main
import "fmt"

func main() {
	array:=[...] string {"qwe", "wer", "ert", "rty", "tyu", "yui", "uio", "iop"}
	var index1,index2 int
	fmt.Println("Enter index 1: ")
	fmt.Scanln(&index1)
	fmt.Println("Enter index 2: ")
	fmt.Scanln(&index2)
	if index1 >= 0 && index1 < len(array) && index2 > index1 && index2 <= len(array) {
		slice1:=array[:index1]
		slice2:=array[index1:index2]
		slice3:=array[index2:]
		fmt.Println(slice1,slice2,slice3)
	} else {
		fmt.Println("Incorrect Indexes")
	}

}

/*
Output - 
sharanya@sharanya-ASUS-TUF-Gaming-A15-FA506IHRB-FA506IHRZ:~/Desktop/MyGoPracs$ go run question4.go
Enter index 1: 
2
Enter index 2: 
4
[qwe wer] [ert rty] [tyu yui uio iop]
*/