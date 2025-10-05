package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("helo", "world")
	fmt.Println(a, b)
}
