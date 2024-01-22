//Sharanya Datrange
/*
I can be placed before V (5) and X (10) to make 4 and 9. 
X can be placed before L (50) and C (100) to make 40 and 90. 
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer.
*/
package main
import "fmt"

func main() {
	var input string
	fmt.Printf("Enter Roman number: ")
	fmt.Scanln(&input)
	
	var storeval int = 0
	for i:=0;i<len(input);i++ {
		switch input[i] {
		case 'I':
			
			if input[i+1] == 'V' || input[i+1] == 'X' {
				storeval -= 1
			} else {
				storeval += 1
			}
			
		case 'V':
			storeval += 5
		case 'X':
			if input[i+1] == 'L' || input[i+1] == 'C' {
				storeval -= 10
			} else {
				storeval += 10
			}
		case 'L':
			storeval += 50
		case 'C':
			if input[i+1] == 'D' || input[i+1] == 'M' {
				storeval -= 100
			} else {
				storeval += 100
			}
		case 'D':
			storeval += 500
		case 'M':
			storeval += 1000
		default:
			fmt.Println("Wrong input. Enter valid Roman number")
		}
	}
		
		
	fmt.Println("Number corresponding is: ",storeval)
}
	