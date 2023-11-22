package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	// userNames := []string{}

	userNames[0] = "Julie"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Manuel")

	fmt.Println(userNames)

	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.7

	courseRatings.output()

	// fmt.Println(courseRatings)

	for index, value := range userNames {
		// ...
		fmt.Println("Index:", index)
		fmt.Println("Value:", value)
	}

	for key, value := range courseRatings {
		// ...
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
