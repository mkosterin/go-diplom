package dataStructs

var SmsOperators = []string{"Topolo", "Kildy", "Rond"}

type SmsData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}
