package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	strconv "strconv"
	"strings"
)

type Knapsack struct {
	size    int
	fitness int
	beta    float32
	profit  []int
	weight  []int
}

type KnapsackResponse struct {
	profit int
	weight int
	items  []bool
}

func getBestSolution(response *KnapsackResponse, personalResponse *KnapsackResponse) *KnapsackResponse {
	if response.profit > personalResponse.profit {
		return response
	}
	return personalResponse
}

func getKnapsackByFile(filename string) Knapsack {

	var f, err = os.Open(filename)

	if err != nil {
		fmt.Println("error opening file ", err)
		os.Exit(1)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	var size int
	var fitness int
	var profit []int
	var weight []int
	var temp int

	var i uint16 = 0

	for {

		path, err := r.ReadString(10) // 0x0A separator = newline

		s := strings.Split(path, "\n")
		line := strings.Split(s[0], " ")

		switch i {

		case 0:
			size, _ = strconv.Atoi(s[0])
			break

		case 1:
			for j := 0; j < size; j++ {
				temp, _ = strconv.Atoi(line[j])
				profit = append(profit, temp)
			}
			break

		case 2:
			for j := 0; j < size; j++ {
				temp, _ = strconv.Atoi(line[j])
				weight = append(weight, temp)
			}
			break

		case 3:
			fitness, _ = strconv.Atoi(s[0])
			break

		}

		if err == io.EOF {
			break
		}

		i++

	}

	var beta float32 = 0
	var betaTemp float32 = 0

	for i := 0; i < size; i++ {
		betaTemp = float32(profit[i]) / float32(weight[i])

		if betaTemp > beta {
			beta = betaTemp
		}

	}

	knapsack := Knapsack{size: size, fitness: fitness, profit: profit, weight: weight, beta: beta}

	return knapsack

}
