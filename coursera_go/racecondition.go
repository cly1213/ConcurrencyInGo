package main

import "fmt"

func main() {
	a := 0
	time := 100

	for i := 0; i < time; i++ {
		go func() {
			a++
		}()
	}

	fmt.Printf("a: %d\n", a)
}

/*
Race condition ?
1. It occurs when two or more process can access and change the shared data at the same time.
It occurred because there were conflicting accesses to a resource .
Critical section problem may cause race condition.

2. The result could be random number but less than or equal to 100.
*/
