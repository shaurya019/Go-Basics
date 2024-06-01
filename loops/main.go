package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	arr := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"};

	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}


	for i:= range arr{
		fmt.Println(arr[i])
	}

	for i,val := range arr {
		fmt.Printf("index is  and value is %v\n", i,val);
	}

}