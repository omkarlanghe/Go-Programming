package main

import "fmt"

func main() {
	var array [3]int
	fmt.Println(array) //all the elements in an array are automatically initialized to 0

	// assining elements to an array
	array[0] = 1
	array[1] = 2
	array[2] = 3
	fmt.Println(array)

	//declaring array using short hand declaration
	cars := [3]string{"aston martin", "bugatti", "mercedis"}
	// fmt.Println(cars)

	for i := 0; i < 3; i++ {
		fmt.Println("cars => ", cars[i])
	}

	//It is not necessary all the elements has to be assigned, those which are not assigned automatically assigned to 0
	var age [5]int
	age[0] = 23
	age[3] = 22

	for i := 0; i < len(age); i++ {
		fmt.Println("age => ", age[i])
	}

	//let compiler decides a length for you based on the elements you put in an array
	lazy := [...]int{1, 2, 33, 53, 531}

	fmt.Println(lazy)
}
