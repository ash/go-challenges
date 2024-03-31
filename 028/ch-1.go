// Solution of Task 1 of The Weekly Challenge 28
// https://theweeklychallenge.org/blog/perl-weekly-challenge-028/

// $ go run ch-1.go
// 19:22:06

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now.Format("15:04:05"))
}
