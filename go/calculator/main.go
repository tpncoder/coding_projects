package main

import (
	"fmt"
	"math"
)

var(
	firstNUM float64
	secondNUM float64
	operation string
	result float64
)

func userInput() {
	fmt.Print("Enter the first no. : ")
	fmt.Scanln(&firstNUM)
	fmt.Print("Enter the second no. : ")
	fmt.Scanln(&secondNUM)

	fmt.Print("Enter the Operation: ")
	fmt.Print(" \n 1. Addition\n 2. Subtraction\n 3. Multiplication\n 4. Division\n 5. Other\n")
	fmt.Print(">>> ")
	fmt.Scanln(&operation)
}

func calcution() {
	if operation == "1" || operation == "Addition" {
		result =firstNUM + secondNUM
		SResult := fmt.Sprintf("%v", result)
		fmt.Printf("The answer is " + SResult)

	} else if operation == "2" || operation == "Subtraction" {
		result =firstNUM - secondNUM
		SResult := fmt.Sprintf("%v", result)
		fmt.Printf("The answer is " + SResult)

	} else if operation == "3" || operation == "Multiplication" {	
		result =firstNUM * secondNUM
		SResult := fmt.Sprintf("%v", result)
		fmt.Printf("The answer is " + SResult)

	} else if operation == "4" || operation == "Division"{
		result =firstNUM / secondNUM
		SResult := fmt.Sprintf("%v", result)
		fmt.Printf("The answer is " + SResult)

	} else if operation == "5" || operation == "Other"{
		fmt.Print(" 1. Exponent\n 2. Root\n 3. Average\n" )
		fmt.Print(">>> ")
		fmt.Scanln(&operation)

		if operation == "1" || operation == "Exponent"{
			result = math.Pow(firstNUM, secondNUM)
			SResult := fmt.Sprintf("%v", result)
			fmt.Printf("The answer is " + SResult)

		} else if operation == "2" || operation == "Root"{
			result = math.Sqrt(firstNUM)
			SResult := fmt.Sprintf("%v", result)
			fmt.Println("The root of the first no. is " + SResult)
			
			result = math.Sqrt(secondNUM)
			SResult = fmt.Sprintf("%v", result)
			fmt.Println("The root of the second no. is " + SResult)

		} else if operation == "3" || operation == "Average"{
			result = firstNUM + secondNUM
			average := result / 2   
			SResult := fmt.Sprintf("%v", average)
			fmt.Printf("The answer is " + SResult) 
		}

	} else {
		fmt.Print("Operation not Found \n Exiting....")
	}
}

func main() {
	userInput()
	calcution()
}
