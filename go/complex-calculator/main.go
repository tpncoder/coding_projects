package main

import(
		"fmt"
		"complex-calculator/standard"
		"complex-calculator/scientific"
)

var(
		calculator_choice int8;
		operation int8;
)

func main(){
	fmt.Println("Calculator\nChoose calculator\n 1. Standard\n 2. Scientific");
    fmt.Print(">>> ");
	fmt.Scanln(&calculator_choice);
		

	if calculator_choice == 1{
		fmt.Println(" 1. Addition\n 2. Subtraction\n 3. Multiplication\n 4. Division");
		fmt.Print(">>> ");
		fmt.Scanln(&operation);

		if operation == 1{
				standard.Add();
		} else if operation == 2{
				standard.Subtract();
		} else if operation == 3{
				standard.Multiply();
		} else if operation == 4{
				standard.Divide();
		} else{
				fmt.Print("Error");
		}
	} else if calculator_choice == 2{
		fmt.Println(" 1. Square\n 2. Cube\n 3. Exponent\n 4. Root\n 5. Sin\n 6. Cos\n 7. Tan");
		fmt.Print(">>> ");
		fmt.Scanln(&operation);

		if operation == 1{
				scientific.Square();
		} else if operation == 2{
				scientific.Cube();
		} else if operation == 3{
				scientific.Exponent();
		} else if operation == 4{
				scientific.Root();
		} else if operation == 5{
				scientific.Sin();
		} else if operation == 6{
				scientific.Cos();
		} else if operation == 7{
				scientific.Tan();
		} else{
			fmt.Print("Error");
		
		};
	};
};
