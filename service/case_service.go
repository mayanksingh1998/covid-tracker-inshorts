package service

import (
	"covid-tracker/contract"
	"covid-tracker/manager"
	"fmt"
)

func GetUserLocationCases(location *contract.UserLocation) [] contract.Cases{
	state := GetLocationDetails(location)
	fmt.Println(state)
	cases := manager.GetCaseForCity(state)
	fmt.Println(cases)
	return cases
}