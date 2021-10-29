package Max

import "reflect"

//Write a function that finds the maximum value in an int slice ([]int).
//Make it work for floats as well

func FindMaxInt(inputArr []int) int {
	max := -9999999999
	for i := 0; i < len(inputArr); i++ {
		if inputArr[i] > max {
			max = inputArr[i]
		}
	}
	return max
}

func FindMaxFloat32(inputArr []float32) float32 {
	var max float32 = -99999999999.0
	for i := 0; i < len(inputArr); i++ {
		if inputArr[i] > max {
			max = inputArr[i]
		}
	}
	return max
}

func Max(inputArr interface{}) interface{} {
	if reflect.TypeOf(inputArr).Kind() != reflect.Slice && reflect.TypeOf(inputArr).Kind() != reflect.Array {
		panic("Sort only works on arrays!")
	}

	switch reflect.TypeOf(inputArr).Elem().Kind() {
	case reflect.Int:
		return FindMaxInt(inputArr.([]int)) // something.([]int) is called a type assertion
	case reflect.String:
		return FindMaxFloat32(inputArr.([]float32))
	default:
		return nil
	}
}
