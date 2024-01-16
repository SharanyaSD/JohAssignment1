package main
import "fmt"
//import "os"
//import "fmt";import "os"
/*
import (
	"fmt"
	"os"
)
*/

func main() {
	fmt.Println("Hello")
	fmt.Println("SWITCH statement")
	//break keyword is implicit. 
	fmt.Println("Choose a choice, enter a number")
	var choice int
	fmt.Scanln(&choice)
	switch(choice) {
		case 1:
			fmt.Print("Basic variables and display")
			var a int 
			var b bool
			var c float64
			var str string
			//fmt.Println(a,b,c,str)
			fmt.Printf("%T %T %T %T \n",a,b,c,str) 		//type of var
			fmt.Printf("%v %v %v %q \n",a,b,c,str)
		case 2:
			fmt.Print("If else statements")
			var input int
			fmt.Println("Enter a number: ")
			fmt.Scanln(&input)
			if(input%2==0) {
				fmt.Println(input ," is even no. ")
			} else if(input%2!=0) {
				fmt.Println(input ," is a odd no. ")
			} else {
				fmt.Println(input ,"is not number or 0")
			}
		case 3:
			fmt.Print("for loop")
			for as:=0;as<5;as++ {
				for bs:=5;bs>0;bs-- {
				fmt.Print(as,bs,"\n")
			}
			}
			
			/*
			//for loop infinitive
			for true  {
				fmt.Print("loop forever")
			}
			*/
			
		default:
			fmt.Print("None")
					
	
		
	
	}
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	}
