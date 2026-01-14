package main

import "fmt"

type Vertes1 struct {
	X int
	Y int
}

func main() {
	v := Vertes1{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}