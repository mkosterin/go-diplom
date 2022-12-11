package repository

import (
	"diplom/internal/dataStructs"
	"encoding/csv"
	"log"
	"os"
)

func SmsReadCsvFile(filePath string, countries map[string]string) (response []dataStructs.SmsData, err error) {
	//Read source file
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("Unable to read input file "+filePath, err)
		return
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	var buffer dataStructs.SmsData
	for {
		line, _ := csvReader.Read()
		if line != nil {
			if smsChecker(line, countries) {
				buffer.Country = line[0]
				buffer.Bandwidth = line[1]
				buffer.ResponseTime = line[2]
				buffer.Provider = line[3]
				response = append(response, buffer)
			}
		} else {
			break
		}
	}
	if err != nil {
		log.Printf("Unable to parse file as CSV for "+filePath, err)
		return
	}
	return response, nil
}

func SmsWriteCsvFile(smsStore *[]dataStructs.SmsData, filePath string) error {
	recordsToWrite := make([][]string, 0)
	for i := 0; i < len(*smsStore); i++ {
		f0 := (*smsStore)[i].Country
		f1 := (*smsStore)[i].Bandwidth
		f2 := (*smsStore)[i].ResponseTime
		f3 := (*smsStore)[i].Provider
		f := []string{f0, f1, f2, f3}
		recordsToWrite = append(recordsToWrite, f)
	}
	f, err := os.Create(filePath)
	if err != nil {
		log.Printf("Unable to write output file "+filePath, err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.WriteAll(recordsToWrite)
	if err := w.Error(); err != nil {
		log.Printf("error writing csv:", err)
		return err
	}
	return nil
}

func smsChecker(line []string, countries map[string]string) bool {
	//Syntax check, according the rules
	if len(line) != 4 {
		return false
	}
	if countries[line[0]] == "" {
		return false
	}
	for i := 0; i < len(dataStructs.SmsOperators); i++ {
		if line[3] == dataStructs.SmsOperators[i] {
			return true
		}
	}
	return false
}
