//Sharanya Datrange


/*
Write a program to calculate the simple interest 
First-line has the comma-separated values of the Principal, rate and time (in years) respective
*constraints: Round simple interest float value to 2 decimal places
*/

package main

import (
	"fmt"
	"math"
	)

func main() {
	var principal, rate float64
	var time int
	fmt.Print("Enter principal amount: ")
	fmt.Scanln(&principal)
	fmt.Print("Enter rate: ")
	fmt.Scanln(&rate)
	fmt.Print("Enter time in years: ")
	fmt.Scanln(&time)
	var simpleInterest float64
	simpleInterest= (principal*rate*float64(time))/100
	roundedSI:=math.Round(simpleInterest*100)/100
	fmt.Println("Simple Interest is: ",roundedSI)
	
}

/*
Output - 
sharanya@sharanya-ASUS-TUF-Gaming-A15-FA506IHRB-FA506IHRZ:~/Desktop/MyGoPracs$ go run q1Simple_Interest.go
Enter principal amount: 456
Enter rate: 52
Enter time in years: 7
Simple Interest is:  1659.84


*/
