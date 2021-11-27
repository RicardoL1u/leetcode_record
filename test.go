package main

import "fmt"

func testPtr(b *int) {
	fmt.Print(b)
}

func main() {
	a := 10
	b := &a
	testPtr(b)
	fmt.Println(b)
}
