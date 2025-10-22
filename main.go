package main

import (
	"fmt"
	"github.com/atiixx/deepdive/core/list"
)

func main() {
	l := list.List[int]{}
	fmt.Println("List created!")
	l.Append(2)
	l.Append(5)
	l.Print()
	l.Prepend(1)
	l.Print()
	l2 := list.List[int]{}
	l2.Prepend(3)
	l2.Print()
	l2.Prepend(16)
	l2.Print()
	fmt.Println(l2.Find(3))
	fmt.Println(l2.Find(2))
}
