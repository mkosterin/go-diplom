package dataStructs

type RawStruct struct {
	SMS       []SmsData
	MMS       []MMSData
	VoiceData []VoiceData
	Email     []EmailData
	Billing   BillingData
	Support   []SupportData
	Incidents []IncidentData
	Error     []error
}

type ResultT struct {
	Status bool       `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

type ResultSetT struct {
	SMS       [][]SmsData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceData              `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   BillingData              `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}

/*
func NewResultT(sms []SmsData, mms []MMSData,
	voice []VoiceData, email []EmailData, billing BillingData,
	support []SupportData, incident []IncidentData, countries map[string]string) *ResultT {

	var result ResultT
	result.Data = *newResultSetT(sms, mms, voice, email, billing, support, incident, countries)
	result.Error = ""
	result.Status = true
	return &result
}

func newResultSetT(sms []SmsData, mms []MMSData,
	voice []VoiceData, email []EmailData, billing BillingData,
	support []SupportData, incident []IncidentData, countries map[string]string) *ResultSetT {
	var result ResultSetT

	//SMS section
	for key := range sms {
		sms[key].Country = countries[sms[key].Country]
	}
	result.SMS = make([][]SmsData, 2)
	result.SMS[0] = make([]SmsData, len(sms))
	result.SMS[1] = make([]SmsData, len(sms))
	sort.Slice(sms, func(i, j int) (less bool) {
		return sms[i].Country < sms[j].Country
	})
	for i := range sms {
		result.SMS[0][i] = sms[i]
	}
	sort.Slice(sms, func(i, j int) (less bool) {
		return sms[i].Provider < sms[j].Provider
	})
	for i := range sms {
		result.SMS[1][i] = sms[i]
	}

	//MMS section
	for key := range mms {
		mms[key].Country = countries[mms[key].Country]
	}
	result.MMS = make([][]MMSData, 2)
	result.MMS[0] = make([]MMSData, len(mms))
	result.MMS[1] = make([]MMSData, len(mms))
	sort.Slice(mms, func(i, j int) (less bool) {
		return mms[i].Country < mms[j].Country
	})
	for i := range mms {
		result.MMS[0][i] = mms[i]
	}
	sort.Slice(mms, func(i, j int) (less bool) {
		return mms[i].Provider < mms[j].Provider
	})
	for i := range mms {
		result.MMS[1][i] = mms[i]
	}

	//Voice section
	result.VoiceCall = voice

	//Email section
	emailCountries := make(map[string][]EmailData)
	sortedEmailCountries := make(map[string][]EmailData)
	result.Email = make(map[string][][]EmailData)
	for i := range email {
		emailCountries[email[i].Country] = append(emailCountries[email[i].Country], email[i])
	}
	for key, value := range emailCountries {
		sort.Slice(value, func(i, j int) (less bool) {
			return value[i].DeliveryTime < value[j].DeliveryTime
		})
		sortedEmailCountries[key] = value
	}
	for key, value := range sortedEmailCountries {
		result.Email[key] = make([][]EmailData, 2)
		result.Email[key][0] = make([]EmailData, 3)
		result.Email[key][1] = make([]EmailData, 3)
		result.Email[key][0] = value[:3]
		result.Email[key][1] = value[len(value)-4 : len(value)-1]
	}

	//Billing section
	result.Billing = billing

	//Support section
	result.Support = make([]int, 2)
	ticketsCount := 0
	averageTime := 60 / 18
	for i := range support {
		ticketsCount = ticketsCount + support[i].ActiveTickets
	}
	if ticketsCount < 9 {
		result.Support[0] = 1
	} else if ticketsCount >= 8 && ticketsCount < 16 {
		result.Support[0] = 2
	} else {
		result.Support[0] = 3
	}
	result.Support[1] = averageTime * ticketsCount

	//Incidents section
	//Получите данные об истории инцидентов.
	//Отсортируйте полученные данные так, чтобы все инциденты со статусом active оказались наверху списка,
	//а остальные ниже. Порядок остальной сортировки не важен.
	result.Incidents = make([]IncidentData, 0)
	for i := range incident {
		if incident[i].Status == "active" {
			result.Incidents = append(result.Incidents, incident[i])
		}
		if incident[i].Status == "closed" {
			result.Incidents = append(result.Incidents, incident[i])
		}
	}

	return &result
}
*/
