
package main
import "fmt"

// Slice in Go lang
/*
Slice is a variable-length sequence which stores elements of a similar type, you are not allowed to store different type of elements in the same slice.

syntax is

var slice_name []variable_type

var slice_name = []variable_type{value1, value2, value3, ...}
var slice_name = []variable_type{value1, value2, value3, ...}


arr := [5]int{10, 20, 30, 40, 50}
slice := arr[1:4]  // Slice contains {20, 30, 40}
slice := arr[:4]   // Slice contains {10, 20, 30, 40}
slice := arr[1:]   // Slice contains {20, 30, 40, 50}
slice := arr[:]    // Slice contains {10, 20, 30, 40, 50}

// skip next number of elements
slice := arr[:len(arr)-1]      // remove last element
slice := arr[1:]               // remove first element
slice := arr[2:]               // remove first two elements
slice := arr[:len(arr)-2]      // remove last two elements

*/


// using make function




func main() {
	slice := make([]int,5) // create a slice of len of 5 with default value 0
	new_slice := append(slice,4,5)
	fmt.Println(new_slice) 



	// append element in middle
	new_slice = append(new_slice[:2], 10, 20, 30, 40, 50)
	fmt.Println(new_slice)

	// remove element from slice
	new_slice = append(new_slice[:2], new_slice[3:]...)
	fmt.Println(new_slice)

	// append elemt in middle step by step and easy way 
	new_slice = append(new_slice[:2], 30, 40, 50) // append 30, 40, 50 at index 2
	fmt.Println(new_slice) 

	// append slice to another slice
	slice1 := []int{1,2,3,4,5}
	slice2 := []int{6,7,8,9,10}
	slice1 = append(slice1, slice2...) // ... is used to append all elements of slice2
	fmt.Println(slice1)
	
}