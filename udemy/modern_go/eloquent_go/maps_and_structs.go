package main

import "fmt"

func main() {
	x := make(map[string]int)
	x["first"] = 1
	x["second"] = 2
	fmt.Println(x)

	// check if a key exists in map using an
	// initialization expression, if ok, print the value!
	if v, ok := x["first"]; ok {
		fmt.Println(v)
	}

	//Another way to initialize a map is using literal init
	y := map[string]int{
		"third": 3,
		"four":  4,
	}
	fmt.Println("Before deleting the key third from map y")
	fmt.Println(y)

	// to delete a key
	delete(y, "third")
	fmt.Println("After deleting the key third from map y")
	fmt.Println(y)

	//let's talk about structs

	type person struct {
		Name    string
		Age     int
		Address string
	}

	jason := person{
		Name:    "Jason",
		Age:     5,
		Address: "Germany",
	}

	fmt.Println(jason.Name)
}
