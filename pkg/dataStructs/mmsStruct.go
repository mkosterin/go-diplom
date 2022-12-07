package dataStructs

var MmsOperators = []string{"Topolo", "Kildy", "Rond"}

type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}
