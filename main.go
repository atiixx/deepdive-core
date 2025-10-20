package main

import (
	"fmt"
	"github.com/atiixx/deepdive/core/list"
)

func main() {
	l := list.List{4}
	fmt.Println("List created!")
	fmt.Println(l.GetValue())
	l.SetValue(2)
	fmt.Println(l.GetValue())
}
