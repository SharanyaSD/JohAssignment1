package main

import (
	"fmt"
)

type Option struct{}

func main() {
	var input int
	fmt.Println("Enter option value 1 to 4: ")
	fmt.Scanln(&input)

	opt := Option{}

	switch input {
	case 1:
		opt.AcceptAnything(1)
	case 2:
		opt.AcceptAnything("Hello")
	case 3:
		opt.AcceptAnything(true)
	case 4:
		opt.AcceptAnything("Custom")
	default:
		fmt.Println("Invalid option")
	}
}

func (o *Option) AcceptAnything(i interface{}) {
	switch value := i.(type) {
	case int:
		fmt.Println("This is a value of type Integer:", value)
	case string:
		fmt.Println("This is a value of type String:", value)
	case bool:
		fmt.Println("This is a value of type Boolean:", value)
	default:
		fmt.Println("Unknown type")
	}
}

/*
Output
go run A4question1.go
Enter option value 1 to 4: 
2
This is a value of type String: Hello
*/
