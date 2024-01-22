//Sharanya Datrange
/*
Write a Program to fulfil the following conditions.
Input: A string containing words separated by space
Output: A slice containing words with the highest frequency in the input string.
Note: The order of the words in the resultant slice should be the same as the order they appear in the input string.
Condition: Words are case-sensitive. Joe is not the same as joe.
Given Code Template: A code template is provided that takes a string as an input and turns it into a slice of strings.
Example Test Case 1:
Input: My name is Joe and My Father's name is also Joe
Output: [My name is Joe]
*/
package main

import (
	"fmt"
	"strings"
	//"math"
	"bufio"
	"os"
)


func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string input: ")
	inputStr, _ := reader.ReadString('\n')
	//splitString := strings.Split(inputstr, " ")
	
	//wordFrequency:=map[string] int {}
	splitString := strings.Fields(inputStr)
	wordFrequency := make(map[string]int)

	for _, word := range splitString {
		wordFrequency[word]++
	}

	maxWords, maxFrequency := findMaxFrequency(wordFrequency)
	fmt.Println("Words with max frequency:", maxWords, "- Highest occurrence:", maxFrequency)
}




//math.Max?

func findMaxFrequency(wordFreq map[string]int) ([]string, int) {
	var maxFrequency int
	var multiple []string

	for word, frequency := range wordFreq {
		if frequency > maxFrequency {
			maxFrequency = frequency
			multiple = []string{word}
		} else if frequency == maxFrequency {
			multiple = append(multiple, word)
		}
	}
	return multiple, maxFrequency
}
/*
sharanya@sharanya-ASUS-TUF-Gaming-A15-FA506IHRB-FA506IHRZ:~/Desktop/MyGoPracs$ go run question3.go
Enter a string input: 
hi hi hello hello
Words with max frequency: [hello hi] - Highest occurrence: 2
*/
