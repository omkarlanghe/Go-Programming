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

	fmt.Println("salaries of employes are : ", salary)

	// find if key is present in map or not...
	newEmp := "omkar"
	value, ok := salary[newEmp]

	if ok == true {
		fmt.Println("Salary of", newEmp, "is", value)
	} else {
		fmt.Println("No", newEmp, "employee present")
	}

	birthYear := map[string]int{
		"omkar":    1996,
		"prafulla": 1994,
		"amit":     1995,
	}
	birthYear["kunal"] = 1996
	name := "kunal"
	fmt.Println("birth year of", name, "is", birthYear[name])

	//how to iterate over map..? we can use "range" form of for loop something like this..
	for key, value := range salary {
		fmt.Printf("salary of [%s] = %d\n", key, value)
	}
}

// In the above code, salary is nil hence it will be initialised using make function
