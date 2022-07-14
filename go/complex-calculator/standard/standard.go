package standard

import(
		"fmt"
)

var(
		fnum float64;
		snum float64;
		result float64;
		str_result string;
)

func Add(){
		fmt.Print("Enter first no. : ");
		fmt.Scanln(&fnum);
		fmt.Print("Enter second no. : ");
		fmt.Scanln(&snum);

		result = fnum + snum;
		str_result = fmt.Sprintf("%v", result);
		fmt.Println("The answer is", str_result);
}

func Subtract(){
		fmt.Print("Enter first no. : ");
		fmt.Scanln(&fnum);
		fmt.Print("Enter second no. : ");
		fmt.Scanln(&snum);	

		result = fnum - snum;
		str_result = fmt.Sprintf("%v", result);
		fmt.Println("The answer is", str_result);
}

func Multiply(){
		fmt.Print("Enter first no. : ");
		fmt.Scanln(&fnum);
		fmt.Print("Enter second no. : ");
		fmt.Scanln(&snum);

		result = fnum * snum;
		str_result = fmt.Sprintf("%v", result);
		fmt.Println("The answer is", str_result);
}

func Divide(){
		fmt.Print("Enter first no. : ");
		fmt.Scanln(&fnum);
		fmt.Print("Enter second no. : ");
		fmt.Scanln(&snum);

		result = fnum / snum;
		str_result = fmt.Sprintf("%v", result);
		fmt.Println("The answer is", str_result);
}
