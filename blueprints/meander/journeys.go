package meander

import "strings"

type j struct {
	Name       string
	PlaceTypes []string
}

//Journeys describes the places you will go to in a day/night
//we use an interface slice because we will later implement a ay of exposing public data
//regardless of actual types
var Journeys = []interface{}{
	j{Name: "Romantic", PlaceTypes: []string{"park", "bar", "movie_theater", "restaurant", "florist"}},
	j{Name: "Shopping", PlaceTypes: []string{"department_store", "cafe", "clothing_store", "jewelry_store", "shoe_store"}},
	j{Name: "Night Out", PlaceTypes: []string{"bar", "casino", "food", "night_club", "bar", "bar", "hospital"}},
	j{Name: "Culture", PlaceTypes: []string{"museum", "cafe", "cemetary", "library", "art_gallery"}},
	j{Name: "Pamper", PlaceTypes: []string{"hair_care", "beauty_salon", "cafe", "spa"}},
}

func (j j) Public() interface{} {
	return map[string]interface{}{
		"name":    j.Name,
		"journey": strings.Join(j.PlaceTypes, "|"),
	}
}
