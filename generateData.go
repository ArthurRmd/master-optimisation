
package main

import (
	"math/rand"
)

func generateRandomArray(size int) []bool {

	var arrayOfBool = make([]bool, size)

	for i := 0; i < size; i++ {
		if rand.Intn(2) == 1{
			arrayOfBool[i] = true
		}
	}

	return arrayOfBool
	



}

func changeBit( array []bool) []bool {
	temp := rand.Intn(len(array))
	return changeBitIndex(array, temp)


}

func changeBitIndex(array []bool,index int) []bool  {
	var newArray = make([]bool, len(array))
	copy(newArray, array)
	newArray[index] = !newArray[index]
	return newArray
}