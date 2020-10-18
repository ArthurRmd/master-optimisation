package main

import (
	"math/rand"
	"strconv"
)

func generateArray(numberOfElement int, arrayBinary []string) []bool {

	var array = make([]bool, numberOfElement)

	numberOfElement = numberOfElement /10

	for i := 0; i < numberOfElement; i++ {

		temp := rand.Intn(1024)

		chaine := arrayBinary[temp]

		for j := 0; j < 10; j++ {
			if chaine[j] == 48 {
				array[(10*i) + j] = true
			}
		}


	}
	return array
}

func generateBinaries() []string  {

	var binaries []string

	for i := 0; i < 1024; i++ {
		temp := strconv.FormatInt(int64(i), 2)
		for len(temp) < 10 {
			temp = "0" + temp
		}
		binaries = append(binaries, temp)
	}

	return binaries

}
