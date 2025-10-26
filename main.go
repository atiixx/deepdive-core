package main

import (
	"fmt"
	"hash/fnv"
)

func main() {
	test := fnv.New64()
	test.Write([]byte("test"))
	test2 := fnv.New64()
	test2.Write([]byte("tesa"))

	for {
		fmt.Println(test2.Sum64())
		fmt.Println(test.Sum64())
	}

}
