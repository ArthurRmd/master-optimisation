package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Knapsack struct {
	size      int
	maxWeight int
	profit    []int
	weight    []int
}

type KnapsackResponse struct {
	profit int
	weight int
	items []bool
}

func getEval(knapsack Knapsack, binaries []string) KnapsackResponse{

	var weight int = 0
	var profit int = 0

	items := generateArray(knapsack.size, binaries)

	for i := 0; i < knapsack.size; i++ {
		if items[i] {
			weight += knapsack.weight[i]
			profit += knapsack.profit[i]
		}

	}

	if weight > knapsack.maxWeight {
		var beta int = 0
		var betaTemp int = 0

		for i := 0; i < knapsack.size ; i++  {
			betaTemp = knapsack.profit[i] / knapsack.weight[i]

			if betaTemp > beta {
				beta = betaTemp
			}

		}

		profit = profit - ( beta * (weight - knapsack.maxWeight))
	}

	return KnapsackResponse{profit: profit, weight: weight, items: items}
}

func getRandom( numberOfRandom int, knapsack Knapsack, binaries *[]string ) KnapsackResponse {

	bestKnapsackResponse := getEval(knapsack, *binaries)

	for i := 0; i < (numberOfRandom - 1); i++ {
		temp := getEval(knapsack, *binaries)

		if temp.profit > bestKnapsackResponse.profit {
			bestKnapsackResponse = temp
		}

	}

	return bestKnapsackResponse

}

func getKnapsackByFile(filename string) Knapsack {

	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("error opening file ", err)
		os.Exit(1)
	}

	defer f.Close()
	r := bufio.NewReader(f)

	var i uint8 = 0

	var size int
	var maxWeight int
	var profit []int
	var weight []int
	var temp int

	for {

		path, err := r.ReadString(10) // 0x0A separator = newline

		s := strings.Split(path, "\n")
		line := strings.Split(s[0], " ")

		if i == 0 {
			size, _ = strconv.Atoi(s[0])
		}
		if i == 1 {
			for j := 0; j < size; j++ {
				temp, _ = strconv.Atoi(line[j])
				profit = append(profit, temp)
			}
		}
		if i == 2 {
			for j := 0; j < size; j++ {
				temp, _ = strconv.Atoi(line[j])
				weight = append(profit, temp)
			}
		}
		if i == 3 {
			maxWeight, _ = strconv.Atoi(s[0])
		}

		if err == io.EOF {
			break
		}

		i++
	}

	knapsack := Knapsack{size: size, maxWeight: maxWeight, profit: profit, weight: weight}

	return knapsack

}