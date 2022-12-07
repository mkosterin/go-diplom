package repository

import (
	"diplom/pkg/dataStructs"
	"encoding/json"
	"log"
	"net/http"
)

func ReadMms(httpResponse *http.Response) []dataStructs.MMSData {
	var rowData, parsedData []dataStructs.MMSData
	parsedData = make([]dataStructs.MMSData, 0)
	err := json.NewDecoder(httpResponse.Body).Decode(&rowData)
	if err != nil {
		log.Fatal("Unable to parse JSON with MMS ", err)
	}
	for i := 0; i < len(rowData); i++ {
		if mmsChecker(rowData[i]) == true {
			parsedData = append(parsedData, rowData[i])
		}
	}
	return parsedData
}

func mmsChecker(line dataStructs.MMSData) bool {
	//Syntax check, according the rules
	if countries[line.Country] == "" {
		return false
	}
	for i := 0; i < len(dataStructs.MmsOperators); i++ {
		if line.Provider == dataStructs.SmsOperators[i] {
			return true
		}
	}
	return false
}
