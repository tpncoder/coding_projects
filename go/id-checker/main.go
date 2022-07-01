package main

import(
	"fmt"
)

var choice string

func checkID(){
	var userID string
	fmt.Print("Enter your ID: ")
	fmt.Scanln(&userID)
	id := []string{"James", "Tom", "Tejas"}

	switch(userID){
		case "James":
			fmt.Print("Your ID exists!")
			break
		case "Tom":
			fmt.Print("Your ID exists!")
			break
		case "Tejas":
			fmt.Print("Your ID exists!")
			break
		default:
			fmt.Print("Your ID doesn't exist!\n")
			fmt.Println("Would you like to create one?\n (y or n)")
			fmt.Print(">>>")
			fmt.Scanln(&choice)
			switch(choice){
				case "y":
					id = append(id, userID)
					fmt.Println("Your ID has been added")
					fmt.Printf("The current ID list is " + "%v",id)
					break
				case "n":
					break;
			}
		}
	}

func main(){
	checkID()
}
