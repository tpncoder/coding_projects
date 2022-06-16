#include "modules.cpp"

class Program{
	public:
		int users[30] = {1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20};
		int target;
		char choice;
		int new_id_unc;
		int new_id;

		void input(){
			cout << "Enter your ID no. : ";
			cin >> target;
		};
		void check_user(){	
			bool exists = std::find(std::begin(users), std::end(users), target) != std::end(users);
			if(exists){
				cout << "Your ID Exists";
			}
			else {
				cout << "Your ID doesn't exist\n";
				cout << "Would you like to create one? (y or n)\n";
				cout << ">>>";
				cin >> choice;

				if(choice = 'y')
				{
					cout << "Enter your new ID no. : ";
					cin >> new_id_unc;
					string new_id_un = to_string(new_id_unc);
					cout << "Are you sure you want your ID no. to be " + new_id_un + ": ";
					cin >> choice;
					if(choice = 'y')
					{
						int new_id = new_id_unc;
						users[3] = new_id;
					}
					else{
						int new_id_un = 0;
					}
				}
				else{
					cout << "OK";
					cout << "Program Exited ...";
					return;
				}
			};
		};
};

int main()
{
	Program p;
	p.input();
	p.check_user();
	return 0;
};
