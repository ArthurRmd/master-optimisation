package main

func evaluate(knapsack *Knapsack, items []bool) KnapsackResponse {

	var weight int = 0
	var profit int = 0

	for i := 0; i < knapsack.size; i++ {
		if items[i] {
			weight += knapsack.weight[i]
			profit += knapsack.profit[i]
		}
	}

	if weight > knapsack.fitness {
		var temp = float32(profit) - (knapsack.beta * float32(weight-knapsack.fitness))
		profit = int(temp)
	}

	return KnapsackResponse{profit: profit, weight: weight, items: items}
}
