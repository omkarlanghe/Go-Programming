package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var name = "omkar langhe";
	fmt.Println("type %T value %d", name, unsafe.Sizeof(name));

	// Is their any way to create typed const? The answer is YES..
	const typedName string = "I love Go programming language";
	fmt.Println("Type is %T and value is %v", typedName, typedName);

	//typedName = "Can I change this"; => not allowed as typeName is of type const string

	var defaultName = "Omkar";
	type aliasOfStringType string; //alias for string type
	var customerName aliasOfStringType = "Omkar";

	fmt.Println(defaultName);
	fmt.Println(customerName);

	// defaultName = customerName; 
	// The above line of code is not allowed due to go's strongly typed policy check at compile time
}
