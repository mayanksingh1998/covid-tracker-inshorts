package contract

type LocationDetails struct {
	Items []struct {
		Address map[string]string `json:"address"`
	} `json:"items"`
}

type UserLocation struct {
	Latitude string `json:"latitude" form:"latitude" query:"latitude"`
	Longitude string `json:"longitude" form:"longitude" query:"longitude"`
}