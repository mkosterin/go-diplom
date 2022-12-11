package dataStructs

var VoiceOperators = []string{"TransparentCalls", "E-Voice", "JustPhone"}

type VoiceData struct {
	Country             string  `json:"country"`
	Load                int     `json:"bandwidth"`
	AvgAnswerTime       int     `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TtfbClean           int     `json:"ttfb"`
	CallTime            int     `json:"median_of_call_time"`
	UnknowField         int     `json:"voice_purity"`
}
