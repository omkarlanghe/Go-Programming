package main

import "fmt"

func main() {
	var age int = 23; //variable declaration with initialization
	fmt.Println("Age is => ", age);
	fmt.Println("\n");

	age = 22;
	fmt.Println("Age updated => ", age);

	var topic = "Go for fun"; //Type Inference: based on R-Val, Go automatically detects the type of R-Val and make "topic" variable of that type.
	fmt.Println(topic);

	// multiple variable declaration in go
	var x, y, name = 10, 20, "omkar langhe"; //method 1;
	fmt.Println("x = ",x, " y = ", y, " name = ", name);

	// method 2
	var(
	x_1 = 10
	y_1 int
	name_1 = "jai shree ram"
	);

	fmt.Println("x1 = ", x_1, " y = ", y_1, " name = ", name_1);
}
