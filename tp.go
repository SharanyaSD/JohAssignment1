package main

import (
"fmt"
"math"
)

func main() {
var r, area float64
const PI float64 = 3.14

fmt.Println("Enter radius: ")
fmt.Scanln(&r)

area = math.Pi * math.Pow(r, 2)
fmt.Println("Area of circle: ", area)
}


