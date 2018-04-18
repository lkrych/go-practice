package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//type capable of holding JSON
type configuration struct {
	Enabled bool
	Path    string
}

func main() {
	file, err := os.Open("config.json") //open file
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	conf := configuration{}
	err = decoder.Decode(&conf) //parse file into struct
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(conf.Path)
}
