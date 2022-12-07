package dataStructs

var VoiceOperators = []string{"TransparentCalls", "E-Voice", "JustPhone"}

type VoiceData struct {
	Country             string  `json:"country"`
	Load                int     `json:"load"`
	AvgAnswerTime       int     `json:"avgAnswerTime"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connectionStability"`
	TtfbClean           int     `json:"ttfbClean"`
	CallTime            int     `json:"callTime"`
	UnknowField         int     `json:"unknowField"`
}
