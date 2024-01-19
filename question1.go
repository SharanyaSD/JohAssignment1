//Sharanya Datrange
/*
Write a program to store the day(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday) against the respective index of the day(1, 2, 3, 4, 5, 6, 7) in a map.
You can hard-code the map in your program.
Take an integer input from the user and print the day stored against that index and if nothing is to be found for that index,
Print "Not a day"
*/
package main
import (
	"fmt"
)

func main() {
	map1:=map[int]string {
		1:"Monday",
		2:"Tuesday",
		3:"Wednesday",
		4:"Thursday",
		5:"Friday",
		6:"Saturday",
		7:"Sunday",
	}

	fmt.Println("Enter a week day(number): ")
	var index int
	fmt.Scanln(&index)

	value,isPresent := map1[index]
	if isPresent {
		fmt.Println(value)
	} else {
		fmt.Println("Not a day")
	}


}

/*
Output -
sharanya@sharanya-ASUS-TUF-Gaming-A15-FA506IHRB-FA506IHRZ:~/Desktop/MyGoPracs$ go run question1.go
Enter a week day(number): 
6
Saturday
*/