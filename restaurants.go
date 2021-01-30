package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"googlemaps.github.io/maps"
	"net/http"
)

type restaurant struct {
	Name string `json:"name"`
	Address string `json:"address"`
	Rating float32 `json:"rating"`
	Ratings int `json:"ratings"`
	Price int `json:"price"`
	Photo string `json:"photo"`
}

func Restaurants(w http.ResponseWriter, r *http.Request) {
	var loc struct {
		Lat  float64 `json:"lat"`
		Lng float64 `json:"long"`
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if err := json.NewDecoder(r.Body).Decode(&loc); err != nil {
		_, _ = fmt.Fprint(w, "Error")
		return
	}

	rest, err := recommendation(loc.Lat, loc.Lng)
	if err != nil {
		fmt.Println("Error getting recommendations")
	}
	js, err := json.Marshal(rest)
	if err != nil {
		fmt.Println("Error marshalling")
	}
	_, err = w.Write(js)
	if err != nil {
		fmt.Println("Error writing json")
	}
}

func recommendation(lat float64, long float64) ([]restaurant, error) {
	client, err := maps.NewClient(maps.WithAPIKey("AIzaSyCrV3oxI6JEXV797k61ujfif0_tjG9Xckc"))
	if err != nil {
		return nil, errors.New("invalid Google API key")
	}

	req := &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: lat,
			Lng: long,
		},
		Radius:   10000,
		Type:     "restaurant",
		OpenNow: true,
	}

	res, err := client.NearbySearch(context.Background(), req)
	if err != nil {
		return nil, errors.New("invalid places request")
	}

	ret := make([]restaurant, len(res.Results))
	for i := 0; i < len(res.Results); i++ {
		ret[i] = restaurant{
			res.Results[i].Name,
			res.Results[i].Vicinity,
			res.Results[i].Rating,
			res.Results[i].UserRatingsTotal,
			res.Results[i].PriceLevel,
			res.Results[i].Photos[0].HTMLAttributions[0],
		}
	}

	return ret, nil
}
