package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"meander"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	APIKey := &APIKeys{}
	APIKey = APIKey.readKeys()
	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})
	log.Println("Starting web server on 8080")
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}

type APIKeys struct {
	ConsumerKey string `yaml:"GOOGLE_PLACES_SECRET"`
}

func (k *APIKeys) readKeys() *APIKeys {
	yamlFile, err := ioutil.ReadFile("/Users/Leland/go/go_code/blueprints/meander/api_keys/secrets.yml")

	if err != nil {
		log.Printf("Error reading YAML: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, k)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return k
}
