package main

import (
	"math/rand"
	"strconv"
)

var firstArray []bool

func generateArray(numberOfElement int, arrayBinary *[]string) []bool {

	var array = make([]bool, numberOfElement)

	numberOfElement = numberOfElement / 10

	for i := 0; i < numberOfElement; i++ {

		temp := rand.Intn(1024)

		var chaine = (*arrayBinary)[temp]
		for j := 0; j < 10; j++ {
			if chaine[j] == 48 {
				array[(10*i)+j] = true
			}
		}

	}
	return array
}

func generateBinaries() []string {

	var binaries []string
	var arrayZero = [11]string{
		"000000000",
		"000000000",
		"00000000",
		"0000000",
		"000000",
		"00000",
		"0000",
		"000",
		"00",
		"0",
		"",
	}

	for i := 0; i < 1024; i++ {
		temp := strconv.FormatInt(int64(i), 2)
		temp = arrayZero[len(temp)] + temp
		binaries = append(binaries, temp)
	}

	return binaries

}

func generateArrayWalk( size int ) []bool {

	if len(firstArray) == 0 {

		var array = make([]bool, size)

		for i := 0; i < size; i++ {
			temp := rand.Intn(2)

			if temp == 1 {
				array[i] = true
			}

		}

		firstArray = array
	} else {

		temp := rand.Intn(size)

		if firstArray[5] == false {
			firstArray[5] = true
		} else {
			firstArray[5] = false
		}

	}

	return firstArray


}
