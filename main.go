package main

import (
	"fmt"
	"github.com/buptmiao/parallel"
	"math/rand"
	"time"
)

func main() {

	start := time.Now()
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(2)

	binaries := generateBinaries()
	binaries2 := generateBinaries()
	//binaries3 := generateBinaries()

	var knapsack Knapsack = getKnapsackByFile("data.txt")
	var knapsack2 Knapsack = getKnapsackByFile("data.txt")
	//var knapsack3 Knapsack = getKnapsackByFile("data.txt")

	var response1 KnapsackResponse
	var response2 KnapsackResponse
	var response3 KnapsackResponse

	//	response := getRandom(50000, knapsack, binaries)
	//	response2 := getRandom(50000, knapsack2, binaries2)

	numberPerThread := 350000
	p := parallel.NewParallel()

	p.Register(getRandom, numberPerThread, knapsack, binaries).SetReceivers(&response1)
	p.Register(getRandom, numberPerThread, knapsack2, binaries2).SetReceivers(&response2)
	//p.Register(getRandom, numberPerThread, knapsack3, binaries3).SetReceivers(&response3)

	p.Run()

	bestProfit := 0

	if response1.profit> response2.profit {
		bestProfit = response1.profit
	} else {
		bestProfit = response2.profit
	}

	if bestProfit < response3.profit {
		bestProfit = response3.profit
	}

	fmt.Println(bestProfit)



	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
