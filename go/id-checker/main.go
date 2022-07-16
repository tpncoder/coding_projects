package main

import "fmt"

var choice string

func main() {
	checkID()
}

func checkID() {
	id := [10]string{"James", "Tom", "Tejas"}

	var userID string
	fmt.Print("Enter your ID: ")
	fmt.Scanln(&userID)

	for i := 0; i < 10; i++ {
		if id[i] == userID {
			fmt.Println("ID exists")
			break
		} else {
			fmt.Println("ID doesn't exist")
			break
		}
	}
}
