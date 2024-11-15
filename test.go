package main // this means go makes executable from this file

import "fmt" // package from go

func main(){ //just one main function in app

	var name string = "Radovan" // explicitly giving type
	var username = "radovanrasha" // implicitly giving type
	var age int

	age = 25

	fmt.Println("Hello world\nMy name is", name, username, age)

	nametwo := "john" // column instead of var, shorhand can't be used outside of function

	fmt.Println(nametwo)

	ageTwo := 30

	fmt.Println(ageTwo)

	// bits & memory
	var numOne int8 = 25
	var numTwo uint8 = 25 //we are not including negative values just positive values
	var scoreOne float64 = 25.5
	scoreTwo := 25.5 //float64 because float64 is default

	fmt.Println(numOne, numTwo, scoreOne, scoreTwo)
}