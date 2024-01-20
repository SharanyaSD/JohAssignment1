package main

import "fmt"

func main() {
 var choice int
 fmt.Print("Enter 1 for Rectangle or 2 for Square: ")
 fmt.Scan(&choice)

 switch choice {
 case 1:
  var length, breadth int
  fmt.Print("Enter length of the rectangle: ")
  fmt.Scan(&length)
  fmt.Print("Enter breadth of the rectangle: ")
  fmt.Scan(&breadth)

  rectangle := Rectangle{length,breadth}
  fmt.Println("Rectangle:")
  print(rectangle)

 case 2:
  var side int
  fmt.Print("Enter side of the square: ")
  fmt.Scan(&side)

  square := Square{side}
  fmt.Println("Square:")
  print(square)

 default:
  fmt.Println("Invalid choice. Enter 1 or 2.")
 }
}

type Quadrilateral interface {
 Area() int
 Perimeter() int
}

type Rectangle struct {
 length, breadth int
}

type Square struct {
 side int
}

func (r Rectangle) Area() int {
 area := r.length * r.breadth
 return area
}

func (r Rectangle) Perimeter() int {
 perimeter := 2 * (r.length + r.breadth)
 return perimeter
}

func (s Square) Area() int {
 area := s.side * s.side
 return area
}

func (s Square) Perimeter() int {
 perimeter := 4 * s.side
 return perimeter
}

func print(q Quadrilateral) {
 fmt.Println("Area: ", q.Area())
 fmt.Println("Perimeter is: ", q.Perimeter())
}

/*
Output 
sharanya@sharanya-ASUS-TUF-Gaming-A15-FA506IHRB-FA506IHRZ:~/Desktop/MyGoPracs$ go run A4question3.go
Enter 1 for Rectangle or 2 for Square: 1
Enter length of the rectangle: 15
Enter breadth of the rectangle: 1
Rectangle:
Area:  15
Perimeter is:  32
sharanya@sharanya-ASUS-TUF-Gaming-A15-FA506IHRB-FA506IHRZ:~/Desktop/MyGoPracs$ go run A4question3.go
Enter 1 for Rectangle or 2 for Square: 2
Enter side of the square: 15
Square:
Area:  225
Perimeter is:  60
*/
