package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	start := time.Now()
	rand.Seed(484558488)

	//tp2(10000000)


	//bestImprovement(5000)

	var knapsack = getKnapsackByFile("data.txt")
	var items = generateRandomArrayOfBool(knapsack.size)

	response := bestImprovement(&knapsack,items,200)
	fmt.Println()

	}





	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
