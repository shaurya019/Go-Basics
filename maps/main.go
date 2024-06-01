package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	arr := make(map[string]string);

	arr["JS"] = "Javascript"
	arr["RB"] = "Ruby"
	arr["PY"] = "Python"

	fmt.Println("List of all arr: ", arr)
	fmt.Println("JS shorts for: ", arr["JS"])

	for id,val := range arr {
		fmt.Printf("For key v, value is %v\n", id,val);
	}
}