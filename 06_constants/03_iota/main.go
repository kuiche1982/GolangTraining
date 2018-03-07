package main

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const (
	d = iota // 0
	q = "others"
	e = iota // 2
	f = iota // 3
)

func main() {
	fmt.Println(a, b, c)
	fmt.Println(d, e, f)
}
