package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	start := time.Now()
	rand.Seed(2)

	//tp2(1000000)

	for i := 0; i <  1000000; i++ {
		generateArrayWalk(1000)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
