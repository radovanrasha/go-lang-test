package main

import "fmt"

func arraysmain(){ 
	//arrays
	var ages [3]int = [3]int{20, 25, 30} // 3 represents length, we cant change length
	var agesTwo = [3]int{20, 25, 30} // simplier syntax

	names := [4]string{"radovan", "john", "doe", "mario"}
	names[1] = "johnny"

	fmt.Println(ages, len(agesTwo), names)


	//slices-more like arrays in other languages, we can change length (use arrays under the hood)
	var scores = []int{100,150,50}
	scores[0] = 25
	scores = append(scores, 75) //append does not change the original array, it returns a new one

	fmt.Println(scores)

	//slice ranges
	rangeOne := names[2:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	rangeOne = append(rangeOne, "arthur", "michael")
	//if we change the slice, we change the original array but only if we does not exceed the original array length
	//if we exceed the original array length, we create a new array and original array wont be changed

	fmt.Println(rangeOne, rangeTwo, rangeThree, names)

}