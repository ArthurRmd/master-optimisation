package main

import (
	"fmt"
	"runtime"
	"sync"
)

func tp1(){

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

}

func tp2(size int){


	var knapsack = getKnapsackByFile("data.txt")
	var response KnapsackResponse
	var response2 KnapsackResponse

	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		response = getRandomWithWalk(size / 2, knapsack)
		wg.Done()
	}()


	go func() {
		response2 = getRandomWithWalk(size / 2, knapsack)
		wg.Done()
	}()




	wg.Wait()

	fmt.Println(response.profit)
	fmt.Println(response2.profit)

}