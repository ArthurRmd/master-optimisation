package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main() {

	start := time.Now()
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(2)

	binaries := generateBinaries()
	//fmt.Println(len(binaries[len(binaries) - 1]))


	if true {
		var knapsack Knapsack = getKnapsackByFile("data.txt")
		response := getRandom(100000, knapsack, &binaries)
		fmt.Println(response.profit)

	} else {

		for i := 0; i < 100000 ; i++ {
			generateArray(1000, binaries)
		}

	}

	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
