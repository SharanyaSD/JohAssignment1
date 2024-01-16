//Sharanya Datrange

/*
2. Write a program to calculate the area of the circle, First line has a value of the radius of the circle
constraint
1. Use const PI from the package math package
2. Use the Pow function from the math package
3. Round area float value to 2 decimal places


*/

package main

import (
	"fmt"
	"math"
	)

func main() {
	var radius,area float64 
	fmt.Print("Enter radius: ")
	fmt.Scanln(&radius)
	area= math.Pi * math.Pow(radius,2)
	roundArea:=math.Round(area*100)/100
	
	fmt.Println("Area of circle is: ", roundArea)
	
	
}

/*
sharanya@sharanya-ASUS-TUF-Gaming-A15-FA506IHRB-FA506IHRZ:~/Desktop/MyGoPracs$ go run q2AreaCircle.go
Enter radius: 14
Area of circle is:  615.75


*/
O
