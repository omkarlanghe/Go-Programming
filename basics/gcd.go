package main
import "fmt"

func main() {
	var num1, num2 int = 0, 0;
	fmt.Println("\nEnter num1");
	fmt.Scan(&num1);
	fmt.Println("\nEnter num2");
	fmt.Scan(&num2);

	gcd := GCD(num1,num2);
	fmt.Println("GCD is: ", gcd);
}

func GCD(num1, num2 int) int {
	for i := 0 ; num1 != num2 ; i++ {
		if(num1 > num2) {
			num1 = num1 - num2;
		} else {
			num2 = num2 - num1;
		}
	}
	return num1;
}
