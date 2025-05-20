package main

import (
	"fmt" 
	"reflect"
	
)

func main(){
	var name string = "Shubham"
	var age int = 21
	var isCool bool = true
	fmt.Println("Hello World")
	fmt.Println("My name is", name, "and I am", age, "years old.")
	fmt.Println("Is cool", isCool)

	location := "India"
	rollNo := 123
	fmt.Println("I am from", location, "and my roll number is", rollNo)
	fmt.Println(reflect.TypeOf(location))
	fmt.Println(reflect.TypeOf(rollNo))
}
