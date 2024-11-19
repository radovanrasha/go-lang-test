package main

import (
	"fmt"
	"math"
)

func sayHello(n string)  {
	fmt.Printf("good morning %v\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)	
	}
}

//this function have return so we need to declare return type
func circleArea(r float64) float64{
	return math.Pi * r * r
}

func main()  {
	sayHello("radovanrasha")
	names := []string{"radovan", "rasha", "john", "michael"}
	cycleNames(names, sayHello)

	a1 := circleArea(5)

	fmt.Println(a1)
}