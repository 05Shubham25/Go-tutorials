// functions in GO lang

// unction syntax
/*



func function_name( [parameter list] ) [return_types] {
   body of the function
   return [expression]
}




func: Keyword used to declare a function.
function_Name: Name of the function (should be descriptive, e.g., addNumbers, getUser).
parameters: Input values for the function. You can pass 0 or more parameters.
returnType: Specifies the type of the returned value (int, string, float64, etc.).
return: Returns the result of the function.
*/

package main

import "fmt"

/*
Exampple function without parameters and return type
func Greet() {
	fmt.Println("Hello, World!")
}

func main() {
	Greet()
}
*/

/*
function with return value

func add(a int , b int ) int{
	return a + b
}

func main(){
	result := add(10, 20)
	fmt.Println("the return value is: ", result)
}
*/

/*
Function with parameter

func Greet(name string) {
	fmt.Println("Hello, ", name)
	fmt.Println("How are you?", name)
}


func main(){
	Greet("shubham")
	Greet("San")
}
*/

/*
Multiple return values
func divide(a,b int) (int ,int) {
	Quotient := a / b
	Reminder := a % b
	return Quotient,Reminder
}

func main(){
	Quo, Rem :=  divide(40,30)
	fmt.Println("Quotient is ", Quo)
	fmt.Println("Reminder is ", Rem)
	fmt.Println(Quo+ Rem)
}

*/

/*
Function with naed return value
func calulate( a, b int)(sum int , product int){
	sum = a + b
	product =  a*b
	return
}

func main(){
	s , p := calulate(5 ,5)
	fmt.Println("Sum",s)
	fmt.Println("product",p)
}

*/

/*
Variadit Functions ( Multiple arguments)
*/







/*

func main() {
	array := [5]int{1, 2, 3, 4, 5}

	for index, value := range array {
		fmt.Println("index", index, "value", value)

	}
}

*/






/*
REVERSE ARRAY


func sum(numbers ...int) int {
	total := 0

	for _, num :=  range numbers {
		total += num
	}
	return total
}


func main() {
	Sum := sum(1,2,3,4,5)
	fmt.Println("Sum", Sum)
}


*/

/*
func main() {
    day := 3

    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    default:
        fmt.Println("Invalid day")
    }
}


OUTPUT :- Wwdnesday
*/
