package repository

import (
	"diplom/internal/dataStructs"
	"encoding/json"
	"log"
	"net/http"
)

func ReadAccendent(httpResponse *http.Response) ([]dataStructs.IncidentData, error) {
	var rowData, parsedData []dataStructs.IncidentData
	parsedData = make([]dataStructs.IncidentData, 0)
	err := json.NewDecoder(httpResponse.Body).Decode(&rowData)
	if err != nil {
		log.Printf("Unable to parse JSON with Accendents ", err)
		return parsedData, err
	}
	for i := 0; i < len(rowData); i++ {
		parsedData = append(parsedData, rowData[i])
	}
	return parsedData, nil
}
