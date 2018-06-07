package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}
	data := struct {
		Name string
	}{"Luiga Sefter"}

	//when you execute a template you need to pass in two arguments
	// 1) where you want to write the template output
	// 2) the data to be used when executing the template

	//this will output to the command line
	err = t.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}
}
