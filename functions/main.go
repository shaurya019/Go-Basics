package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter();
	res := help(3,5);
	fmt.Println("result",res);
	proRes := proAdder(2, 5, 8, 7, 3)
	fmt.Println("Pro result is: ", proRes)

}

func proAdder(values ...int) (int) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total

}


func help(x int, y int) int {
	return x+y;
}

func greeter () {
	fmt.Println("Namstey from golang")

}