package notification

import (
	"github.com/go-resty/resty/v2"
)

type Message struct {
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Color string `json:"color"`
	Text  string `json:"text"`
}

func Error(err error, mattermostWebhook string) {

	var m Message
	var data Attachment

	if err == nil {
		data.Color = "#B0fc38"
		data.Text = "Success"
	} else {
		data.Color = "#FF0000"
		data.Text = err.Error()
	}
	m.Attachments = append(m.Attachments, data)

	resty.New().R().SetBody(m).Post(mattermostWebhook)

}
