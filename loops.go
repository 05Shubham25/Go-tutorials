// Go only support FOR LOOP
// There is no WHILE or DO WHILE loop in Go
// implementation of WHILE or DO WHILE loop can be done using FOR loop
// strusture of FOR loop in Go
// for initialization; condition; post {
// 	// code
// }
// initialization: executed before the first iteration 


package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

