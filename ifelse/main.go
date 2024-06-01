package main

import "fmt"

func main() {
	fmt.Println("If else in golang")
	loginCount := 10
	var res  string

	if loginCount < 10{
		res = "1234";
	} else if loginCount > 10{
		res = "5678";
	} else {
		res = "0";
	}
	fmt.Println(res);
}