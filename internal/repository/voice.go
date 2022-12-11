package repository

import (
	"diplom/internal/dataStructs"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func VoiceReadCsvFile(filePath string, countries map[string]string) ([]dataStructs.VoiceData, error) {
	//Read source file
	var response []dataStructs.VoiceData
	f, err := os.Open(filePath)
	if err != nil {
		log.Printf("Unable to read input file "+filePath, err)
		return response, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	var buffer dataStructs.VoiceData
	for {
		line, _ := csvReader.Read()
		if line != nil {
			if voiceChecker(line, countries) {
				buffer.Country = line[0]
				buffer.Load, _ = strconv.Atoi(line[1])
				buffer.AvgAnswerTime, _ = strconv.Atoi(line[2])
				buffer.Provider = line[3]
				ConnectionStability, _ := strconv.ParseFloat(line[4], 32)
				buffer.ConnectionStability = float32(ConnectionStability)
				buffer.TtfbClean, _ = strconv.Atoi(line[5])
				buffer.CallTime, _ = strconv.Atoi(line[6])
				buffer.UnknowField, _ = strconv.Atoi(line[7])
				response = append(response, buffer)
			}
		} else {
			break
		}
	}
	if err != nil {
		log.Printf("Unable to parse file as CSV for "+filePath, err)
	}
	return response, err
}

func VoiceWriteCsvFile(voiceStore *[]dataStructs.VoiceData, filePath string) error {
	recordsToWrite := make([][]string, 0)
	for i := 0; i < len(*voiceStore); i++ {
		f0 := (*voiceStore)[i].Country
		f1 := strconv.Itoa((*voiceStore)[i].Load)
		f2 := strconv.Itoa((*voiceStore)[i].AvgAnswerTime)
		f3 := (*voiceStore)[i].Provider
		f4 := fmt.Sprintf("%f", (*voiceStore)[i].ConnectionStability)
		f5 := strconv.Itoa((*voiceStore)[i].TtfbClean)
		f6 := strconv.Itoa((*voiceStore)[i].CallTime)
		f7 := strconv.Itoa((*voiceStore)[i].CallTime)
		f := []string{f0, f1, f2, f3, f4, f5, f6, f7}
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

func voiceChecker(line []string, countries map[string]string) bool {
	//Syntax check, according the rules
	if len(line) != 8 {
		return false
	}
	if countries[line[0]] == "" {
		return false
	}
	for i := 0; i < len(dataStructs.VoiceOperators); i++ {
		if line[3] == dataStructs.VoiceOperators[i] {
			return true
		}
	}
	return false
}
