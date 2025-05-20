
/*

Map in Go lang

Map is a collection of key-value pairs. It is a built-in data structure in Go that is used to store key-value pairs. The keys are unique within a map while the values may not be. The key-value pairs are unordered.

syntax is

var map_name map[key_data_type]value_data_type
var map_name = map[key_data_type]value_data_type{key1: value1, key2: value2, key3: value3, ...}
var map_name = map[key_data_type]value_data_type{key1: value1, key2: value2, key3: value3, ...}



mapName := make(map[KeyType]ValueType)
mapname[KeyType1] = ValueType1
mapname[KeyType2] = ValueType2


*/

package main

import "fmt"

/*
func main() {
	person := make(map[string]string)
	person["name"] = "megamind"
	person["location"] = "mumbai"
	fmt.Println(person)

	fmt.Println(person["name"])
	delete(person,"location")
	fmt.Println(person["location"])

}

*/

/*
func main() {
	salaries := map[string]float64 {
		"manoj": 12.231,
		"sam": -10.000,
		"shubham": 100.00,
		"keshav": 50.00,
	}
	fmt.Println(salaries)

	fmt.Println("sanjana salaray", salaries["sanjana"])
} 


*/



/*
func main() {
	scores := map[string]int{"shubham": 80,"Raj":70}

	// value, exists := scores["Raj"]

	// delete(scores, "shubham") -----> delete index and value form map

	value, exists := scores["shubham"]

	// exists ----- > it check if key exist in map or not 

	if exists{
		fmt.Println("the value is", value)
	} else {
		fmt.Println("the value doesnt exist")
	}


}

*/




func main() {
	fruits := map[string]string{
		"mango":"yellow",
		"apple": "red",
	}

	for key, value := range fruits{
		fmt.Println("fruits", key, "color", value)
	}
}