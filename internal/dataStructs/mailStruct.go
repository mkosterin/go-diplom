package dataStructs

var MailOperators = []string{"Gmail", "Yahoo", "Hotmail", "MSN", "Orange", "Comcast",
	"AOL", "Live", "RediffMail", "GMX", "Proton Mail", "Yandex", "Mail.ru"}

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}
