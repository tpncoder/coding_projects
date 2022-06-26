#include "modules.cpp"

class calculator
{
	public:
		float num1;
		float num2;
		int operation;
		int result;
		int extra_operation;

		void input(){
			cout << "Enter first no.: ";
			cin >> num1;
			
			cout << "Not required for Root and Rounding\n"
			     << "Enter second no.: ";
			cin >> num2;

			cout << "Enter operation: \n";
			cout 
				<< "1. Addition\n"
				<< "2. Subtraction\n"
				<< "3. Multiplication\n"
				<< "4. Division\n"
				<< "5. Other\n";
			cout << ">>> ";
			cin >> operation;
		};

		void calculation()
		{
			switch(operation){
				case 1:
					cout << "The answer is: " << num1 + num2;
					break;
				case 2:
					cout << "The answer is: " << num1 - num2;
					break;
				case 3:
					cout << "The answer is: " << num1 * num2;
					break;
				case 4:
					cout << "The answer is: " << num1 / num2;
					break;
				case 5:
					cout << "Enter the operation: \n";
					cout << " 1. Exponent\n 2. Root \n 3. Average \n";
					cout << ">>> ";
					cin >> extra_operation;

					switch(extra_operation){
						case 1:
							cout << "The answer is: " << pow(num1,num2);
							break;
						case 2:
							cout << "The answer is: " << sqrt(num1);
							break;
						case 3:
							cout << "The answer is: " << ceil(num1);
							break;
						default:
							cout << "\nOperation Not Found \n Exiting...";
							break;
					}
				default:
					cout << "\nOperation Not Found \n Exiting...";
					break;
		};						
	};
};
int main()
{
	calculator calc;
	calc.input();
	calc.calculation();
	return 0;
};
