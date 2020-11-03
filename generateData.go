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


func generateRandomArrayOfBool(size int) []bool {

	var arrayOfBool = make([]bool, size)

	for i := 0; i < size; i++ {
		temp := rand.Intn(2)

		if temp == 1 {
			arrayOfBool[i] = true
		}

	}

	return arrayOfBool

}

func changeBitArrayOfBool( array []bool) []bool {

	temp := rand.Intn(len(array))
	changeBit(array, temp)
	return array

}

func changeBit(array []bool,index int) []bool  {
	var newArray = make([]bool, len(array))
	copy(newArray, array)
	newArray[index] = !newArray[index]
	return newArray
}


func generateArrayWalk( size int ) []bool {

	if len(firstArray) == 0 {

		firstArray = generateRandomArrayOfBool(size)

	} else {
		changeBitArrayOfBool(firstArray)
	}

	return firstArray

}
