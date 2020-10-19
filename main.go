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

	var knapsack Knapsack = getKnapsackByFile("data.txt")
	response := getRandom(100000, knapsack, binaries)
	fmt.Println(response.profit)

	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
