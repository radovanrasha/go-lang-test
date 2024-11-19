package main

import "fmt"

func booleans()  {
	age := 45
	
	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 45)
	fmt.Println(age != 50)

	if age > 30 {
		fmt.Println("Age is more than 30")
	} else if age > 40{
		fmt.Println("age is more than 40")
	}else{
		fmt.Println("Age is less than 30")
	}

	names := []string{"john", "mike","randall"}

	for index, value := range names {
		if index == 1 {
			fmt.Println("countinuing", index)
			continue
		}

		fmt.Printf("The value at pos %v is %v", index, value)
	}
}