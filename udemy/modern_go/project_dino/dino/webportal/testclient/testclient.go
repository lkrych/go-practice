package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Animal struct {
	ID         int
	AnimalType string
	Nickname   string
	Zone       int
	Age        int
}

func main() {
	data := &Animal{
		AnimalType: "Gecko",
		Nickname:   "Veronica",
		Zone:       3,
		Age:        24,
	}

	var b bytes.Buffer
	json.NewEncoder(&b).Encode(data)

	// resp, err := http.Post("http://localhost:8080/api/animals/add", "application/json", &b)
	// if err != nil || resp.StatusCode != 200 {
	// 	log.Fatal(err)
	// }

	resp, err := http.Post("http://localhost:8080/api/animals/edit/Veronica", "application/json", &b)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal(err)
	}

}
