package meander

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	yaml "gopkg.in/yaml.v2"
)

type Place struct {
	*googleGeometry `json:"geometry"`
	Name            string         `json:"name"`
	Icon            string         `json:"icon"`
	Photos          []*googlePhoto `json:"photos"`
	Vicinity        string         `json:"vicinity"`
}

type googleResponse struct {
	Results []*Place `json:"results"`
}

type googleGeometry struct {
	*googleLocation `json:"location"`
}

type googleLocation struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type googlePhoto struct {
	PhotoRef string `json:"photo_reference"`
	URL      string `json:"url"`
}

type Query struct {
	Lat          float64
	Lng          float32
	Journey      []string
	Radius       int
	CostRangeStr string
}

func (p *Place) Public() interface{} {
	return map[string]interface{}{
		"name":     p.Name,
		"icon":     p.Icon,
		"photos":   p.Photos,
		"vicinity": p.Vicinity,
		"lat":      p.Lat,
		"lng":      p.Lng,
	}
}

func (q *Query) find(types string) (*googleResponse, error) {
	APIKey := &APIKeys{}
	APIKey = APIKey.readKeys()
	u := "https://maps.googleapis.com/maps/api/place/nearbysearch/json"
	vals := make(url.Values)
	vals.Set("location", fmt.Sprintf("%g,%g", q.Lat, q.Lng))
	vals.Set("radius", fmt.Sprintf("%d", q.Radius))
	vals.Set("types", types)
	vals.Set("key", APIKey.ConsumerKey)
	if len(q.CostRangeStr) > 0 {
		r, err := ParseCostRange(q.CostRangeStr)
		if err != nil {
			return nil, err
		}
		vals.Set("minprice", fmt.Sprintf("%d", int(r.From)-1))
		vals.Set("maxprice", fmt.Sprintf("%d", int(r.To)-1))
	}
	res, err := http.Get(u + "?" + vals.Encode())
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var response googleResponse
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, nil
}

//Run runs the query concurrently, and returns the results.
func (q *Query)

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
