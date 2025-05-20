// stack

/*
Initialize an empty list as stack

Push(x):
    Append x to the stack

Pop():
    If stack is not empty:
        Remove and return the last element
    Else:
        Print "Stack Underflow"

Peek():
    If stack is not empty:
        Return last element without removing
    Else:
        Print "Stack is empty"

IsEmpty():
    Return True if stack length is 0, else False

Size():
    Return length of stack

*/

package main

import (
	"fmt"
)


type Stack struct{
    item []int
}

// push item into stack
func (s *Stack) push(x int){
    s.item = append(s.item,x)
}

// pop item of the stack

func (s *Stack) pop() int{

    if len(s.item) == 0 {
        fmt.Println("stack is underflow!!!")
        return -1
    }

    getlastvalue := s.item[len(s.item)-1]
    s.item = s.item[:len(s.item)-1]
    return getlastvalue
}

// peek the last element 

func (s *Stack) peek() int{
    if len(s.item) == 0{
        fmt.Println("stack is empty")
        return -1
    }

    return s.item[len(s.item)-1] 

}

// IsEmpty is stack is empty

func (s *Stack) IsEmpty() bool {
    return len(s.item) == 0
}

func (s *Stack) Size() int {
    return len(s.item)
}


func main() {
    Stack :=  Stack{}

    // push 

    Stack.push(20)
    Stack.push(10)
    Stack.push(30)
    fmt.Println("Push to stake", Stack.item)
    // pop

    Stack.pop()

    fmt.Println("removed element", Stack.pop())    
    fmt.Println("removed element", Stack.item)    

    //
    Stack.peek()
    fmt.Println("removed element", Stack.item)    


}


