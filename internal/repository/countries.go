package repository

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

func GetCountries() map[string]string {
	resp, err := http.Get("http://country.io/names.json")
	if err != nil {
		log.Printf("error read countries: %s", err.Error())
		os.Exit(1)
	}
	defer resp.Body.Close()
	jsonCountriesData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("error unmarshal countries: %s", err.Error())
		os.Exit(1)
	}
	countriesMap := make(map[string]string)
	err = json.Unmarshal([]byte(jsonCountriesData), &countriesMap)
	return countriesMap
}
