package main

import "fmt"
import "math"

func main() {
	name, age := "omkar langhe", 23; //short hand declarations
	fmt.Println(name, " : ", age);

	a, b := 10.00, 20.00;
	fmt.Println("Minimum is : ", math.Min(a,b));
}
