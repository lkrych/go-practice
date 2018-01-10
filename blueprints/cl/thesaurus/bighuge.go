package thesaurus

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	ThesaurusKey string `yaml:"THESAURUS_API_KEY"`
}

type synonyms struct {
	Noun *words `json:"noun"`
	Verb *words `json:"verb"`
}

type words struct {
	Syn []string `json:"syn"`
}

func (c *Config) Synonyms(term string) ([]string, error) {
	var syns []string
	response, err := http.Get("http://words.bighugelabs.com/api/2/" + c.ThesaurusKey + "/" + term + "/json")
	if err != nil {
		return syns, errors.New("bighuge: Failed when looking for synonyms for " + term + " " + err.Error())
	}
	var data synonyms
	defer response.Body.Close()
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		return syns, err
	}
	if data.Noun != nil {
		syns = append(syns, data.Noun.Syn...)
	}
	if data.Verb != nil {
		syns = append(syns, data.Verb.Syn...)
	}
	return syns, nil
}

func (c *Config) GetConfig() *Config {
	yamlFile, err := ioutil.ReadFile("secrets.yml")

	if err != nil {
		log.Printf("Error reading YAML: %v", err)
	}
	err = yaml.Unmarshal(yamlFile, c)

	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}
