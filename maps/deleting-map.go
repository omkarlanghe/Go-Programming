package main

import (
	"fmt"
)

func main() {
	var salary map[string]int
	if salary == nil {
		fmt.Println("map is nil. going to make one..")
		salary = make(map[string]int)
		salary["roshan"] = 240000
		salary["ajinkya"] = 400000
		salary["ram"] = 500000
	}

	// fmt.Println("salaries of employes are : ", salary)

	// how to iterate over map..? we can use "range" form of for loop something like this..
	for key, value := range salary {
		fmt.Printf("salary of [%s] = %d\n", key, value)
	}

	// Deleting element from map
	fmt.Println("After Deleting")
	delete(salary, "ram")
	for key, value := range salary {
		fmt.Printf("salary of [%s] = %d\n", key, value)
	}

	fmt.Println("lenght of map salary is", len(salary))
}

/*
	maps are reference type, i.e if we assign one map object to another newly created map object, changes made in new map object reflects on orignal map as,
	they both point to same memory location / data structure
*/
