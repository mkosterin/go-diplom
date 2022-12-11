package repository

import (
	"diplom/internal/dataStructs"
	"encoding/json"
	"log"
	"net/http"
)

func ReadSupport(httpResponse *http.Response) ([]dataStructs.SupportData, error) {
	var rowData, parsedData []dataStructs.SupportData
	parsedData = make([]dataStructs.SupportData, 0)
	err := json.NewDecoder(httpResponse.Body).Decode(&rowData)
	if err != nil {
		log.Printf("Unable to parse JSON with Support ", err)
		return parsedData, err
	}
	for i := 0; i < len(rowData); i++ {
		parsedData = append(parsedData, rowData[i])
	}
	return parsedData, nil
}
