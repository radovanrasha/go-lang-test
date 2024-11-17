package main

import "fmt"

func loop(){ 
	x := 0

	// like while loop
	for x < 5{
		fmt.Println("value of x is", x)
		x++
	}

	// standard for loop
	for i := 0; i < 10; i++ {
		fmt.Println("value of i is:", i)
	}

	names := []string{"john", "victor", "michael"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Printf("value at index %v is %v\n", index, value)
	}

	// we can use just what we want, underscore changes index
	for _, value := range names {
		fmt.Printf("value is %v\n", value)
		value = "new name" // this will not change original array, cuz value is like local copy of it
	}

	fmt.Println(names)

}