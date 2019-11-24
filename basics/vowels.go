package main

import "fmt"

func main() {
	var ch string;
	fmt.Println("Enter Vowel");
	fmt.Scan(&ch);

	switch ch {
	case "a", "e", "i", "o", "u", "A", "E", "I", "O", "U":
		fmt.Println("Its a Vowel");
	default:
		fmt.Println("Its not a Vowel");
	}
}
