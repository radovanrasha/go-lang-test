package main

import (
	"fmt"
	"sort"
	"strings"
)

func main(){ 
	greeting := "Hello world!"

	// string methods
	fmt.Println(strings.Contains(greeting, "Hello")) //its case sensitive
	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi")) //it does not change the original string
	fmt.Println(strings.ToUpper(greeting)) //it does not change the original string
	fmt.Println(strings.Index(greeting, "wor"))
	fmt.Println(strings.Split(greeting, " "))

	//
	ages := []int{45,50,20,53,12,64,73}

	sort.Ints(ages) // sort change original datapackage main
	fmt.Println(ages)

	index := sort.SearchInts(ages, 20) // it searches sorted array because data is altered in line 22
	// if there is not that number in array it returns the index where it should be
	fmt.Println(index)

	names := []string{"john", "mike","randall"}

	sort.Strings(names)

	fmt.Println(sort.SearchStrings(names, "rasha"))

}