/*
🔹 Structs in Go (Structure)
A struct is a custom data type that groups related data together. It's like a blueprint for objects (just like classes in OOP, but without methods).


type StructName struct {
    field1 dataType1
    field2 dataType2
    // ...
}

func main(){
	struct1 := StructName{field1:"value_filed1",field2:"value_filed2",}
}

*/





/*

🔹 2. Pointer Receiver
The actual object is used. Changes affect the original struct.


func (p *Person) updateName(newName string) {
    p.name = newName // this updates the actual object
}

func main() {
    p := Person{"John", 30}
    p.updateName("Mike")
    fmt.Println(p.name) // Now prints "Mike"
}


✅ When to Use Pointer Receivers
When the method modifies the data

When you want to avoid copying large structs

To be consistent — many Go developers use pointer receivers by default

*/



package main

import "fmt"


type CollegeData struct{
	name string
	roll_no int
}

func main()  {

	sakec := CollegeData{name:"shubham", roll_no: 37}
	iitb := CollegeData{"Rishabh", 100}

	fmt.Println("name", sakec.name)
	fmt.Println("VS")
	fmt.Println("name", iitb.name)
	
	fmt.Println("roll_no", sakec.roll_no)
	fmt.Println("VS")
	fmt.Println("roll_no", iitb.roll_no)

}