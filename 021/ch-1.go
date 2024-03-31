// Solution of Task 1 of The Weekly Challenge 21
// https://theweeklychallenge.org/blog/perl-weekly-challenge-021/

/*
Test run:

$ go run ch-1.go 20
2.7182818284590455
*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	max := 100

	if len(os.Args) == 2 {
		max, _ = strconv.Atoi(os.Args[1])
	}

	e := 1.0
	f := 1.0

	for n := 1; n != max; n++ {
		f *= float64(n)
		e += 1.0 / f
	}

	fmt.Println(e)
}
