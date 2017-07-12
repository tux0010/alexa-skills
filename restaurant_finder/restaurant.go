package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	log "github.com/sirupsen/logrus"
)

type Restaurant struct {
	Name    string
	Address string
	Phone   string
}

type YelpResp struct {
	Total      int `json:"total"`
	Businesses []struct {
		Phone    string `json:"phone"`
		Name     string `json:"name"`
		Location struct {
			City     string `json:"city"`
			Address1 string `json:"address1"`
			Zipcode  string `json:"zip_code"`
		} `json:"location"`
	} `json:"businesses"`
}

// TODO: Split this into a service (source from Yelp, Infatuation, etc)
func restaurantRecommendation(cuisine string) (*Restaurant, error) {
	apiToken := ""

	v := url.Values{}
	v.Set("term", fmt.Sprintf("%s restaurants", cuisine))
	v.Set("location", "77904")
	v.Set("open_now", "true")

	u := fmt.Sprintf("https://api.yelp.com/v3/businesses/search?%s", v.Encode())
	req, err := http.NewRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiToken))

	log.Printf("Searching Yelp for %s", u)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var yr YelpResp
	str, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(str, &yr)
	if err != nil {
		return nil, err
	}

	log.Info(yr)

	if yr.Total == 0 {
		return &Restaurant{}, nil
	}

	r := Restaurant{
		Name:    yr.Businesses[0].Name,
		Address: yr.Businesses[0].Location.Address1,
		Phone:   yr.Businesses[0].Phone,
	}

	return &r, nil
}
