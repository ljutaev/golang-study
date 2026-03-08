package main

import "fmt"

func main() {

	// pointer
	first := 100
	second := &first

	fmt.Println(second)
	fmt.Println(*second)
}
