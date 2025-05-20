//Array in Go lang 

/* 
syntax is 

1) var variable_name [size] variable_type
2) var variable_name [size] variable_type = [size]variable_type{value1, value2, value3, ...}
3) var variable_name = [size]variable_type{value1, value2, value3, ...}
*/


package main
import "fmt"

/*
this is how to assign array 3 ways 
func main() {
	var arr [5]int  //Array of size 4
	arr[0] = 10
	arr[1]= 20
	arr[2] = 50
	arr[3] = 30
	arr[4]= 40
	short_arr := [5]int{1,2,3,4,5}  //  Shorthand declaration

	inferred_arr := [...]int{1,2,3,4,5,6,7,8} //✅ Array with inferred size

	fmt.Println(arr, short_arr, inferred_arr)
	fmt.Println(arr[2], short_arr[1], inferred_arr[4]) // accessing Array elements

}

*/

/*
Iterating Over Arrays

func main(){

	// Using for loop
	numbers := [...] int{1,2,3,4,5,6}
	for i := 0; i < len(numbers); i++{
		fmt.Println("Index:", i , " Value:", numbers[i],)
	}



	// using range
	for index, value :li= range(numbers){       
		fmt.Println("Index" ,index, "value",value)
	}

}

*/

// reversearray



func reverseArray(arr [5]int ) {

	// Print the original array
	fmt.Println("original Array:", arr)

	// print reverse array 
	fmt.Print("Reversed Array: [")

	// loop through the array in reverse order
	for i := len(arr)-1 ; i>=0; i-- {
		fmt.Print(arr[i])

		// print comma 
		if i != 0 {
			fmt.Print(",")
		}
	}

	fmt.Print("]")
	
}


func main()  {
	numbers :=[5]int{2,3,4,5,6}
	reverseArray(numbers)
}