package main

import "fmt"

func isEven(a int) bool {

	if a%2 == 0 {
		return true
	} else {
		return false
	}
}

func isPrime(a int) bool {

	if isEven(a) == true && a != 2 {
		fmt.Println("ce n'est pas un nombre premier")
		return false
	} else {
		fmt.Println(" c'est un nombre premier")
		return true
	}
}

func main() {
	fmt.Println(isPrime(8))
}
