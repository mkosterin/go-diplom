package repository

import (
	"diplom/internal/dataStructs"
	"errors"
	"log"
	"net/http"
)

func collectData(config dataStructs.Config, countries map[string]string) (rawStruct dataStructs.RawStruct) {
	rawStruct.Error = make([]error, 0)

	smsDataPath := config.SmsSource
	var smsErr error
	rawStruct.SMS, smsErr = SmsReadCsvFile(smsDataPath, countries)
	if smsErr != nil {
		log.Printf("error reading SMS data: %s", smsErr.Error())
		rawStruct.Error = append(rawStruct.Error, smsErr)
	}
	//write slice "sms" to CSV file
	smsDataSave := config.SmsTarget
	if err := SmsWriteCsvFile(&rawStruct.SMS, smsDataSave); err != nil {
		log.Printf("error writing SMS data to CSV: %s", err.Error())
		rawStruct.Error = append(rawStruct.Error, err)
	}

	mmsDataUrl := config.MmsSource
	response, mmsErr := http.Get(mmsDataUrl)
	if mmsErr != nil {
		log.Printf("MMS error making web request: %s", mmsErr.Error())
		rawStruct.Error = append(rawStruct.Error, mmsErr)
	} else if response.StatusCode != 200 {
		log.Printf("MMS error making web request: %s", response.Status)
		rawStruct.Error = append(rawStruct.Error, errors.New("MMS Response code not 200"))
	} else {
		rawStruct.MMS, mmsErr = ReadMms(response, countries)
	}
	if mmsErr != nil {
		log.Printf("MMS error making web request: %s", mmsErr.Error())
		rawStruct.Error = append(rawStruct.Error, mmsErr)
	}

	voiceDataPath := config.VoiceSource
	var voiceErr error
	rawStruct.VoiceData, voiceErr = VoiceReadCsvFile(voiceDataPath, countries)
	if voiceErr != nil {
		log.Printf("error reading Voice data: %s", voiceErr.Error())
		rawStruct.Error = append(rawStruct.Error, voiceErr)
	}
	//write slice "voice" to CSV file
	voiceDataSave := config.VoiceTarget
	if voiceErr := VoiceWriteCsvFile(&rawStruct.VoiceData, voiceDataSave); voiceErr != nil {
		log.Printf("error wriring Voice data to CSV: %s", voiceErr.Error())
		rawStruct.Error = append(rawStruct.Error, voiceErr)
	}

	mailDataPath := config.MailSource
	var mailErr error
	rawStruct.Email, mailErr = MailReadCsvFile(mailDataPath, countries)
	if mailErr != nil {
		log.Printf("error reading EMail data: %s", mailErr.Error())
		rawStruct.Error = append(rawStruct.Error, mailErr)
	}
	//write slice "mail" to CSV file
	mailDataSave := config.MailTarget
	if mailErr = MailWriteCsvFile(&rawStruct.Email, mailDataSave); mailErr != nil {
		log.Printf("error wriring EMail data to CSV: %s", mailErr.Error())
		rawStruct.Error = append(rawStruct.Error, mailErr)
	}

	var billingErr error
	billingDataPath := config.BillingSource
	rawStruct.Billing, billingErr = BillingReadFile(billingDataPath)
	if billingErr != nil {
		log.Printf("error reading Billing data: %s", billingErr.Error())
		rawStruct.Error = append(rawStruct.Error, billingErr)
	}

	supportDataUrl := config.SupportSource
	var supportErr error
	response, supportErr = http.Get(supportDataUrl)
	if supportErr != nil {
		log.Printf("Support error making web request: %s", supportErr.Error())
		rawStruct.Error = append(rawStruct.Error, supportErr)
	} else if response.StatusCode != 200 {
		log.Printf("Support error making web request: %s", response.Status)
		rawStruct.Error = append(rawStruct.Error, errors.New("SUPPORT Response code not 200"))
	} else {
		rawStruct.Support, supportErr = ReadSupport(response)
	}
	if supportErr != nil {
		log.Printf("Support error making web request: %s", supportErr.Error())
		rawStruct.Error = append(rawStruct.Error, supportErr)
	}

	accendentDataUrl := config.AccendendSource
	var accendentErr error
	response, accendentErr = http.Get(accendentDataUrl)
	if accendentErr != nil {
		log.Printf("Accendent error making web request: %s", accendentErr.Error())
		rawStruct.Error = append(rawStruct.Error, accendentErr)
	} else if response.StatusCode != 200 {
		log.Printf("Accendent error making web request: %s", response.Status)
		rawStruct.Error = append(rawStruct.Error, errors.New("ACCENDENT Response code not 200"))
	} else {
		rawStruct.Incidents, accendentErr = ReadAccendent(response)
	}
	if accendentErr != nil {
		log.Printf("Accendent error making web request: %s", accendentErr.Error())
		rawStruct.Error = append(rawStruct.Error, accendentErr)
	}

	return rawStruct
}
