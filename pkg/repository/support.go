package repository

import (
	"diplom/pkg/dataStructs"
	"encoding/json"
	"log"
	"net/http"
)

func ReadSupport(httpResponse *http.Response) []dataStructs.SupportData {
	var rowData, parsedData []dataStructs.SupportData
	parsedData = make([]dataStructs.SupportData, 0)
	err := json.NewDecoder(httpResponse.Body).Decode(&rowData)
	if err != nil {
		log.Fatal("Unable to parse JSON with Support ", err)
	}
	for i := 0; i < len(rowData); i++ {
		parsedData = append(parsedData, rowData[i])
	}
	return parsedData
}
