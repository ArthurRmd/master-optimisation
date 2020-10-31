package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {

	start := time.Now()
	//rand.Seed(time.Now().UnixNano())
	rand.Seed(2)

	binaries := generateBinaries()

	var knapsack Knapsack = getKnapsackByFile("data.txt")
	var response KnapsackResponse
	var response2 KnapsackResponse

	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		response = getRandom(50000, knapsack, binaries)
		wg.Done()
	}()


	go func() {
		response2 = getRandom(50000, knapsack, binaries)
		wg.Done()
	}()




	wg.Wait()

	fmt.Println(response.profit)
	fmt.Println(response2.profit)

	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
