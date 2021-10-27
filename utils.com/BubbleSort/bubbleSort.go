package Bubblesort

// https://stackoverflow.com/questions/18041334/convert-interface-to-int

import (
	"reflect"
)

func intSort(unsortedArr []int) []int {
	for i := 0; i < len(unsortedArr); i++ {
		for j := i + 1; j < len(unsortedArr); j++ {
			if unsortedArr[i] > unsortedArr[j] {
				temp := unsortedArr[j]
				unsortedArr[j] = unsortedArr[i]
				unsortedArr[i] = temp
			}
		}
	}

	return unsortedArr
}

func stringSort(unsortedArr []string) []string {
	for i := 0; i < len(unsortedArr); i++ {
		for j := i + 1; j < len(unsortedArr); j++ {
			if unsortedArr[i] > unsortedArr[j] {
				temp := unsortedArr[j]
				unsortedArr[j] = unsortedArr[i]
				unsortedArr[i] = temp
			}
		}
	}

	return unsortedArr
}

func Sort(unsortedArr interface{}) interface{} {

	if reflect.TypeOf(unsortedArr).Kind() != reflect.Slice && reflect.TypeOf(unsortedArr).Kind() != reflect.Array {
		panic("Sort only works on arrays!")
	}

	switch reflect.TypeOf(unsortedArr).Elem().Kind() {
	case reflect.Int:
		intSort(unsortedArr.([]int)) // something.([]int) is called a type assertion
	case reflect.String:
		stringSort(unsortedArr.([]string))
	}

	return unsortedArr
}
