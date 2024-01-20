package main

import "fmt"

func main() {
 var l,b int
 fmt.Println("Enter length of rectangle: ")
 fmt.Scanln(&l)
 fmt.Println("Enter breadth of rectangle: ")
 fmt.Scanln(&b)
	
 rect:=Rectangle {
  length: l,
  breadth:b,
 }

 fmt.Println("Area of rectangle: ",rect.Area())
 fmt.Println("Perimeter of rectangle: ",rect.Perimeter())
	
}

type Rectangle struct {
 length, breadth int
}

func (r *Rectangle) Area() int {
 area:=r.length*r.breadth
 return area
}

func (r *Rectangle) Perimeter() int {
 perimeter:=2*(r.length+r.breadth)
 return perimeter
}


/*
sharanya@sharanya-ASUS-TUF-Gaming-A15-FA506IHRB-FA506IHRZ:~/Desktop/MyGoPracs$ go run A4question2.go
Enter length of rectangle: 
8
Enter breadth of rectangle: 
33
Area of rectangle: 264
Perimeter of rectangle: 82

*/
