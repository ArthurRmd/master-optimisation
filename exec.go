package main

import (
	"fmt"
	"runtime"
	"sync"
)

func tp1(size int) KnapsackResponse {

	binaries := generateBinaries()

	var knapsack Knapsack = getKnapsackByFile("data.txt")
	var response KnapsackResponse
	var response2 KnapsackResponse

	runtime.GOMAXPROCS(4)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		response = getRandom(size/2, knapsack, binaries)
		wg.Done()
	}()

	go func() {
		response2 = getRandom(size/2, knapsack, binaries)
		wg.Done()
	}()

	wg.Wait()

	if  response.profit > response2.profit {
		fmt.Println(response.profit)
		return response
	} else {
		fmt.Println(response2.profit)
		return response2
	}

}

func tp2(size int) {

	var knapsack = getKnapsackByFile("data.txt")
	var response KnapsackResponse
	var response2 KnapsackResponse

	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		response = getRandomWithWalk(size/2, knapsack)
		wg.Done()
	}()

	go func() {
		response2 = getRandomWithWalk(size/2, knapsack)
		wg.Done()
	}()

	wg.Wait()

	fmt.Println(response.profit)
	fmt.Println(response2.profit)

}

func firstImprovement(numberOfEval int) {

		var knapsack = getKnapsackByFile("data.txt")

		var items = generateRandomArrayOfBool(knapsack.size)
		response := makeEval(&knapsack,items)

		fmt.Println(response.profit)

		for i := 0; i < numberOfEval; i++ {
			tempItems := generateRandomArrayOfBool(knapsack.size)
			response := makeEval(&knapsack,tempItems)
			response = response
			for j := 0; j < 100; j++ {

			}

		}




}

func bestImprovement(knapsack *Knapsack, items []bool, numberOfEval int) KnapsackResponse {


	var maxEval = makeEval(knapsack, items)

	for i := 0; i < numberOfEval; i++ {
		maxEval = getMaxChangeBit(knapsack, maxEval)
	}

	return maxEval


}


func getMaxChangeBit(knapsack *Knapsack, response KnapsackResponse) KnapsackResponse {

	basicValue := response

	for i := 0; i < 1000; i++ {

		tempItems := changeBit(response.items, i)
		response2 := makeEval(knapsack, tempItems)

		if response2.profit > basicValue.profit {
			basicValue = response2
		}

	}


	return basicValue

}
