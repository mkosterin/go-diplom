package main

import (
	"diplom/internal/dataStructs"
	"diplom/internal/repository"
	"diplom/internal/web"
	"fmt"
	"log"
	"net/http"
)

func main() {
	//Init config
	config := repository.ConfigReader()
	countries := repository.GetCountries()
	//*******************************************
	//stage 2
	//read, parse source CSV, save to slice "sms"
	smsDataPath := config.SmsSource
	sms := repository.ReadCsvFile(smsDataPath, countries)
	//write slice "sms" to CSV file
	smsDataSave := config.SmsTarget
	if err := repository.WriteCsvFile(&sms, smsDataSave); err != nil {
		log.Fatalf("error wriring SMS data to CSV: %s", err.Error())
	}
	//end of stage 2
	//*******************************************

	//*******************************************
	//stage 3
	mmsDataUrl := config.MmsSource
	response, err := http.Get(mmsDataUrl)
	if err != nil {
		log.Fatalf("MMS error making web request: %s", err.Error())
	}
	if response.StatusCode != 200 {
		log.Fatalf("MMS error making web request: %s", response.Status)
	}
	mms := repository.ReadMms(response, countries)
	fmt.Println("Debug output for MMS:")
	fmt.Println(mms)
	fmt.Println("---------------------")
	//end of stage 3
	//*******************************************

	//*******************************************
	//stage 4
	voiceDataPath := config.VoiceSource
	voice := repository.VoiceReadCsvFile(voiceDataPath, countries)
	//write slice "voice" to CSV file
	voiceDataSave := config.VoiceTarget
	if err := repository.VoiceWriteCsvFile(&voice, voiceDataSave); err != nil {
		log.Fatalf("error wriring SMS data to CSV: %s", err.Error())
	}
	//end of stage 4
	//*******************************************

	//*******************************************
	//stage 5
	mailDataPath := config.MailSource
	mail := repository.MailReadCsvFile(mailDataPath, countries)
	//write slice "mail" to CSV file
	mailDataSave := config.MailTarget
	if err := repository.MailWriteCsvFile(&mail, mailDataSave); err != nil {
		log.Fatalf("error wriring EMail data to CSV: %s", err.Error())
	}
	//end of stage 5
	//*******************************************

	//*******************************************
	//stage 6
	billingDataPath := config.BillingSource
	billing := repository.BillingReadFile(billingDataPath)
	//write slice "mail" to CSV file
	fmt.Println("Debug output for billing:")
	fmt.Println(billing)
	fmt.Println("-------------------------")
	//end of stage 6
	//*******************************************

	//*******************************************
	//stage 7
	supportDataUrl := config.SupportSource
	response, err = http.Get(supportDataUrl)
	if err != nil {
		log.Fatalf("Support error making web request: %s", err.Error())
	}
	if response.StatusCode != 200 {
		log.Fatalf("Support error making web request: %s", response.Status)
	}
	support := repository.ReadSupport(response)
	fmt.Println("Debug output for support:")
	fmt.Println(support)
	fmt.Println("-------------------------")
	//end of stage 7
	//*******************************************

	//*******************************************
	//stage 8
	accendentDataUrl := config.AccendendSource
	response, err = http.Get(accendentDataUrl)
	if err != nil {
		log.Fatalf("Accendent error making web request: %s", err.Error())
	}
	if response.StatusCode != 200 {
		log.Fatalf("Accendent error making web request: %s", response.Status)
	}
	accendent := repository.ReadAccendent(response)
	fmt.Println("Debug output for accendent:")
	fmt.Println(accendent)
	fmt.Println("---------------------------")
	//end of stage 8
	//*******************************************

	//*******************************************
	//stage 10
	status := dataStructs.NewResultT(sms, mms, voice, mail, billing, support, accendent, countries)
	//end of stage 10
	//*******************************************

	//*******************************************
	//stage 9
	serverUrl := config.WebListernerAddress
	web.Router(serverUrl, status)
	//end of stage 9
	//*******************************************
}
