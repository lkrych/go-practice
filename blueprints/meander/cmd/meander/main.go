package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"meander"
	"net/http"
	"strconv"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	var APIKeys = &APIKeys{}
	APIKeys = APIKeys.readKeys()

	http.HandleFunc("/journeys", cors(func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	}))
	http.HandleFunc("/recommendations", cors(func(w http.ResponseWriter, r *http.Request) {
		urlQuery := r.URL.Query()
		q := &meander.Query{
			Journey: strings.Split(urlQuery.Get("journey"), "|")}
		var err error
		q.Lat, err = strconv.ParseFloat(urlQuery.Get("lat"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q.Lng, err = strconv.ParseFloat(urlQuery.Get("lng"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q.Radius, err = strconv.Atoi(urlQuery.Get("radius"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		q.CostRangeStr = urlQuery.Get("cost")
		places := q.Run(APIKeys.ConsumerKey)
		respond(w, r, places)
	}))
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

func cors(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w, r)
	}
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
