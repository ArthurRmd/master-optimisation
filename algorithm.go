package main


func random(numberIteration int) KnapsackResponse {

	var knapsack Knapsack = getKnapsackByFile("files/data.txt")

	var response = evaluate(&knapsack, generateRandomArray(knapsack.size))
	var responseTemp = KnapsackResponse{}

	for i := 0; i < numberIteration; i++ {
		responseTemp = evaluate(&knapsack, generateRandomArray(knapsack.size))
		response = *getBestSolution(&responseTemp, &response)
	}

	return response

}

func walk(numberIteration int) KnapsackResponse {
	var knapsack Knapsack = getKnapsackByFile("files/data.txt")
	var randomArray = generateRandomArray(knapsack.size)
	var response = evaluate(&knapsack, randomArray)
	var responseTemp = KnapsackResponse{}

	for i := 0; i < numberIteration; i++ {
		randomArray = changeBit(randomArray)
		responseTemp = evaluate(&knapsack, randomArray )
		response = *getBestSolution(&responseTemp, &response)
	}

	return response

}