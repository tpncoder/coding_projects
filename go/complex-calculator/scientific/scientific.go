package scientific

import(
		"fmt"
		"math"
)

var(
		fnum float64;
		snum float64;
		result float64;
		//str_result string;
)

func Square(){
		fmt.Print("Enter the no. : ");
		fmt.Scanln(&fnum);

		result = math.Pow(fnum, 2);
		fmt.Println("The answer is ", result);
}

func Cube(){
		fmt.Print("Enter the no. : ");
		fmt.Scanln(&fnum);

		result = math.Pow(fnum, 3);
		fmt.Println("The answer is ", result);
}

func Exponent(){
		fmt.Print("Enter the first no. : ");
		fmt.Scanln(&fnum);
		fmt.Print("Enter the second no. : ");
		fmt.Scanln(&snum);

		result = math.Pow(fnum, snum);
		fmt.Println("The answer is ", result);
}

func Root(){
		fmt.Print("Enter the no. : ");
		fmt.Scanln(&fnum);

		result = math.Sqrt(fnum);
		fmt.Println("The answer is ", result);
}

func Cos(){
		fmt.Print("Enter the no. : ");
		fmt.Scanln(&fnum);

		result = math.Cos(fnum);
		fmt.Println("The answer is ", result);
}

func Sin(){
		fmt.Print("Enter the no. : ");
		fmt.Scanln(&fnum);

		result = math.Sin(fnum);
		fmt.Println("The answer is ", result);
}

func Tan(){
		fmt.Print("Enter the no. : ");
		fmt.Scanln(&fnum);

		result = math.Tan(fnum);
		fmt.Println("The answer is ", result);
}
