package contract


type State struct{
	Name string  `json:"name"`
	StateCode string `json:"state_code"`
	AbbreviationCode string `json:"abbreviation_code"`
}

type Test struct {
	string `json:"state"`
}