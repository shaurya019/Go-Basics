package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")
	var arr = []string{"1","2","3"};
	fmt.Println(arr);

	arr = append(arr, "4","5");

	arr = append(arr[2:])
	arr = append(arr[:3])


	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777

	highScores = append(highScores, 555, 666, 321)

	sort.Ints(highScores);
	sorted := sort.IntsAreSorted(highScores)
    fmt.Println("Are the high scores sorted?", sorted)


	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
	
}