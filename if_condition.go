package main

import "fmt"


func main() {

	// age input from user
	var age int
	fmt.Println("Enter your age: ")
	fmt.Scanln(&age)

	if age != int(age) {
		fmt.Println("Please enter a valid age.")
		return
	}

	// if condition
	if age != int(age) {
		fmt.Println("Please enter a valid age.")
	} else if age >= 18 {
		fmt.Println("You are eligible to vote.")
	} else if age == 17  {
		fmt.Println("You are 1 year to go to be  eligible to vote.")
	} else {
		fmt.Println("You are not eligible to vote.")
	}

}