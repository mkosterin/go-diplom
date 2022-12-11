package main

import (
	"diplom/internal/repository"
	"diplom/internal/web"
)

func main() {
	//Init config
	config := repository.ConfigReader()
	countries := repository.GetCountries()
	serverUrl := config.WebListernerAddress
	web.Router(serverUrl, config, countries)
}
