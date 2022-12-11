package repository

import (
	"diplom/internal/dataStructs"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

func BillingReadFile(filePath string) (response dataStructs.BillingData) {
	body, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	decNumber := number(string(body))
	binNumber := bitMask(decNumber)
	response = generateStruct(binNumber)
	return
}

func number(line string) (result uint8) {
	var buffer float64
	for i := 0; i < len(line); i++ {
		index := len(line) - 1 - i
		bit := string(line[index])
		if bit == "1" {
			buffer = buffer + math.Pow(2, float64(i))
		}
	}
	return uint8(buffer)
}

func bitMask(number uint8) (mask string) {
	return strconv.FormatInt(int64(number), 2)
}

func generateStruct(numberBin string) (response dataStructs.BillingData) {
	length := len(numberBin)
	numberBin = strings.Repeat("0", 6-length) + numberBin
	boolMask := make([]bool, 6)
	for i := 0; i < 6; i++ {
		if numberBin[i:i+1] == "1" {
			boolMask[i] = true
		}
	}
	response.CreateCustomer = boolMask[5]
	response.Purchase = boolMask[4]
	response.Payout = boolMask[3]
	response.Recurring = boolMask[2]
	response.FraudControl = boolMask[1]
	response.CheckoutPage = boolMask[0]
	return
}
