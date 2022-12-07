package repository

import (
	"diplom/pkg/dataStructs"
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func MailReadCsvFile(filePath string) (response []dataStructs.EmailData) {
	//Read source file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	var buffer dataStructs.EmailData
	for {
		line, _ := csvReader.Read()
		if line != nil {
			if mailChecker(line) {
				buffer.Country = line[0]
				buffer.Provider = line[1]
				buffer.DeliveryTime, _ = strconv.Atoi(line[2])
				response = append(response, buffer)
			}
		} else {
			break
		}
	}
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	return response
}

func MailWriteCsvFile(mailStore *[]dataStructs.EmailData, filePath string) error {
	recordsToWrite := make([][]string, 0)
	for i := 0; i < len(*mailStore); i++ {
		f0 := (*mailStore)[i].Country
		f1 := (*mailStore)[i].Provider
		f2 := strconv.Itoa((*mailStore)[i].DeliveryTime)
		f := []string{f0, f1, f2}
		recordsToWrite = append(recordsToWrite, f)
	}
	f, err := os.Create(filePath)
	if err != nil {
		log.Fatal("Unable to write output file "+filePath, err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	w.WriteAll(recordsToWrite)
	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
		return err
	}
	return nil
}

func mailChecker(line []string) bool {
	//Syntax check, according the rules
	if len(line) != 3 {
		return false
	}
	if countries[line[0]] == "" {
		return false
	}
	for i := 0; i < len(dataStructs.MailOperators); i++ {
		if line[1] == dataStructs.MailOperators[i] {
			return true
		}
	}
	return false
}
