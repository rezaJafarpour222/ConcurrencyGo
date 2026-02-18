package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}
func main() {
	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsilon",
	}
	wg.Add(9)
	for i, x := range words {
		printSomething(fmt.Sprintf("%d: %s", i, x), &wg)
	}
	wg.Wait()
}
