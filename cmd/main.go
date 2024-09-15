package main

import (
	"fmt"

	stl "github.com/sangharshseth/Go-STL"
)

func main() {
	stack := stl.Stack[int]{}
	stack.Push(1)
	stack.Push(2)
	fmt.Print(stack.Top())
	stack.Pop()
	fmt.Print(stack.Top())
	fmt.Print(stack.Size())
}
