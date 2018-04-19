package main

import (
	"fmt"

	gcfg "gopkg.in/gcfg.v1"
)

func main() {
	config := struct {
		Section struct {
			Enabled bool
			Path    string
		}
	}{}

	err := gcfg.ReadFileInto(&config, "config.ini")
	if err != nil {
		fmt.Println("Failed to read ini file: ", err)
	}

	fmt.Println(config.Section.Enabled)
	fmt.Println(config.Section.Path)

}
