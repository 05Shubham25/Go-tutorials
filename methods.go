/*
✅ Method Syntax in Go
go

func (receiverName ReceiverType) methodName(parameters) returnType {
    // method body
}

✨ Explanation:
func → keyword to define the method

(receiverName ReceiverType) → the struct (or type) the method belongs to

methodName() → the name of your method

parameters → optional inputs

returnType → optional return type
*/


package main

import "fmt"


/*
A Simple Method

type OfficeSup struct{
	name string
	// role string
	emp_id int
}

func (e OfficeSup) greet(){
	fmt.Println("Hello ",e.name , "Sir! Welcome")
	fmt.Println("Your role is ", e.role)
	fmt.Println("And your emp_id is ", e.emp_id)
}

func main(){
	emp1 := OfficeSup{"Shubham", "Manager", 1000}
	emp1.greet() // method call
}

*/



type OfficeSup struct {
	name   string
	emp_id int
}

// 🔄 Method with formatted string output
func (e OfficeSup) sayAge() string {
	return fmt.Sprintf("%s has emp_id of %d", e.name, e.emp_id)
}

func main() {
	Employee := OfficeSup{"Shubham", 100}
	fmt.Println(Employee.sayAge())
}




