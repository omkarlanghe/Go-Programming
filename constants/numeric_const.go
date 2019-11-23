package main

import (
	"fmt"
)

func main() {
	const a = 5;
	var intVar int = a;
	var int32Var int32 = a;
	var float64Var float64 = a;
	var complex64Var complex64 = a;
	fmt.Println("intVar : ", intVar, "\nint32Var : ", int32Var, "\nfloat64Var : ", float64Var, "\ncomplex64Var : ", complex64Var);
	fmt.Printf("intVar type %T : ", intVar, "\nint32Var type %T: ", int32Var, "\nfloat64Var type %T: ", float64Var, "\ncomplex64Var type %T: ", complex64Var);

	var i = 5
	var f = 5.6
	var c = 5 + 6i
	fmt.Printf("\n\ni's type %T, f's type %T, c's type %T", i, f, c);
}
