package main

import (
	"diplom/pkg/dataStructs"
	"diplom/pkg/repository"
	"diplom/pkg/web"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func main() {
	//read config file
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	//*******************************************
	//stage 2
	//read, parse source CSV, save to slice "sms"
	smsDataPath := viper.GetString("sms.sourceDirectory") + "/" + viper.GetString("sms.sourceFileName")
	sms := repository.ReadCsvFile(smsDataPath)
	//write slice "sms" to CSV file
	smsDataSave := viper.GetString("sms.targetDirectory") + "/" + viper.GetString("sms.targetFileName")
	if err := repository.WriteCsvFile(&sms, smsDataSave); err != nil {
		log.Fatalf("error wriring SMS data to CSV: %s", err.Error())
	}
	//end of stage 2
	//*******************************************

	//*******************************************
	//stage 3
	mmsDataUrl := viper.GetString("mms.scheme") + "://" +
		viper.GetString("mms.address") + ":" +
		viper.GetString("mms.port") + "/" +
		viper.GetString("mms.endpoint")
	response, err := http.Get(mmsDataUrl)
	if err != nil {
		log.Fatalf("MMS error making web request: %s", err.Error())
	}
	if response.StatusCode != 200 {
		log.Fatalf("MMS error making web request: %s", response.Status)
	}
	mms := repository.ReadMms(response)
	fmt.Println("Debug output for MMS:")
	fmt.Println(mms)
	fmt.Println("---------------------")
	//end of stage 3
	//*******************************************

	//*******************************************
	//stage 4
	voiceDataPath := viper.GetString("voice.sourceDirectory") + "/" + viper.GetString("voice.sourceFileName")
	voice := repository.VoiceReadCsvFile(voiceDataPath)
	//write slice "voice" to CSV file
	voiceDataSave := viper.GetString("voice.targetDirectory") + "/" + viper.GetString("voice.targetFileName")
	if err := repository.VoiceWriteCsvFile(&voice, voiceDataSave); err != nil {
		log.Fatalf("error wriring SMS data to CSV: %s", err.Error())
	}
	//end of stage 4
	//*******************************************

	//*******************************************
	//stage 5
	mailDataPath := viper.GetString("mail.sourceDirectory") + "/" + viper.GetString("mail.sourceFileName")
	mail := repository.MailReadCsvFile(mailDataPath)
	//write slice "mail" to CSV file
	mailDataSave := viper.GetString("mail.targetDirectory") + "/" + viper.GetString("mail.targetFileName")
	if err := repository.MailWriteCsvFile(&mail, mailDataSave); err != nil {
		log.Fatalf("error wriring EMail data to CSV: %s", err.Error())
	}
	//end of stage 5
	//*******************************************

	//*******************************************
	//stage 6
	billingDataPath := viper.GetString("billing.sourceDirectory") + "/" + viper.GetString("billing.sourceFileName")
	billing := repository.BillingReadFile(billingDataPath)
	//write slice "mail" to CSV file
	fmt.Println("Debug output for billing:")
	fmt.Println(billing)
	fmt.Println("-------------------------")
	//end of stage 6
	//*******************************************

	//*******************************************
	//stage 7
	supportDataUrl := viper.GetString("support.scheme") + "://" +
		viper.GetString("support.address") + ":" +
		viper.GetString("support.port") + "/" +
		viper.GetString("support.endpoint")
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
	accendentDataUrl := viper.GetString("accendent.scheme") + "://" +
		viper.GetString("accendent.address") + ":" +
		viper.GetString("accendent.port") + "/" +
		viper.GetString("accendent.endpoint")
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
	countries := repository.GetCountries()
	status := dataStructs.NewResultT(sms, mms, voice, mail, billing, support, accendent, countries)
	//end of stage 10
	//*******************************************

	//*******************************************
	//stage 9
	serverUrl := viper.GetString("web.address") + ":" + viper.GetString("web.port")
	web.Router(serverUrl, status)
	//end of stage 9
	//*******************************************
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
