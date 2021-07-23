package service

import (
	"covid-tracker/constants"
	"covid-tracker/contract"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func GetLocationDetails(location *contract.UserLocation) string{
	var latitude float64
	var longitude float64

	latitude, _ = strconv.ParseFloat(location.Latitude, 8)
	longitude, _ = strconv.ParseFloat(location.Longitude, 8)

	url := "https://revgeocode.search.hereapi.com/v1/revgeocode?apiKey=" + constants.LOCATION_API_KEY + "&at=" + fmt.Sprint(latitude) + "," + fmt.Sprint(longitude)
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	var locationDetails contract.LocationDetails
	json.Unmarshal(body, &locationDetails)
	var state = ""
	for _, add := range locationDetails.Items {
		state = add.Address["state"]
	}
	return state
}
