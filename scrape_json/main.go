package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type ReferenceJSON struct {
	MiscInformation struct {
		PhoneMake  string `json:"phoneMake"`
		PhoneModel string `json:"phoneModel"`
		UserData   struct {
			Name string `json:"name"`
		} `json:"userData"`
		Geolocation struct {
			Latitude   float64 `json:"latitude"`
			Longitude  float64 `json:"longitude"`
			LastUpdate string  `json:"lastUpdate"`
		} `json:"GeoLocation"`
	} `json:"misc_information"`
}

type UsableJSON struct {
	Position []float64
}

func main() {
	fmt.Println("Entering main")
	readJSONFiles("/Users/Leland/Desktop/smileID/tanzjson/json")
}

//must take in absolute path
func readJSONFiles(directory string) {
	fmt.Println("Entering readJSONFiles with", directory)
	//read directory of json
	files, err := ioutil.ReadDir(fmt.Sprintf("%v", directory))
	if err != nil {
		fmt.Println("Cannot read directory")
		log.Fatal(err)
	}

	fmt.Println("Found", directory, "files")

	objs := []*UsableJSON{}
	//iterate through files
	for _, f := range files {
		//save scraped json to object and add it to collection
		file, err := ioutil.ReadFile(fmt.Sprintf("%v/%v", directory, f.Name()))
		if err != nil {
			fmt.Printf("There was an error reading %v \n", f.Name())
			log.Fatal(err)
		}
		fmt.Println("Read file", f.Name())
		_, uJSON := jsonToStruct(file)
		objs = append(objs, uJSON)
	}

	//write scraped json to file
	marshalled, err := json.Marshal(objs)
	if err != nil {
		fmt.Println("There was an error marshalling the JSON")
		log.Fatal(err)
	}
	err = ioutil.WriteFile("output.json", marshalled, 0644)
	if err != nil {
		fmt.Println("There was an error writing the marshalled JSON")
		log.Fatal(err)
	}
}

func jsonToStruct(file []byte) (*ReferenceJSON, *UsableJSON) {
	r := &ReferenceJSON{}
	u := &UsableJSON{}
	if err := json.Unmarshal(file, &r); err != nil {
		fmt.Println("Error unmarshalling json")
		log.Fatal(err)
		return r, u
	}
	latLng := []float64{}
	latLng = append(latLng, r.MiscInformation.Geolocation.Latitude)
	latLng = append(latLng, r.MiscInformation.Geolocation.Longitude)
	u.Position = latLng
	// fmt.Printf("%v \n", r)
	return r, u
}
