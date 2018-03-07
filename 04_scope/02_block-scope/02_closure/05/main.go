package main

import "fmt"

func main() {

	counter1 := getCounter()
	counter2 := getCounter()
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter1())
	fmt.Println(counter2())
	fmt.Println(counter2())
	fmt.Println(counter2())
}

func getCounter() func() int {
	x := 0
	fun := func() int {
		x++
		return x
	}
	return fun
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
