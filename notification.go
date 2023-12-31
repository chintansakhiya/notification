package notification

import (
	"github.com/go-resty/resty/v2"
)

type Data struct {
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Color string `json:"color"`
	Text  string `json:"text"`
}

func Error(err error, mattermostWebhook string) error {

	var m Data
	var data Attachment

	if err == nil {
		data.Color = "#B0fc38"
		data.Text = "Success"
	} else {
		data.Color = "#FF0000"
		data.Text = err.Error()
		err = nil
	}
	m.Attachments = append(m.Attachments, data)

	_,err=resty.New().R().SetBody(m).Post(mattermostWebhook)
	return err
}

func Message(text string, mattermostWebhook string) error {

	var m Data
	var data Attachment

	data.Color = "#B0fc38"
	data.Text = "Message :" + text

	m.Attachments = append(m.Attachments, data)

	_, err := resty.New().R().SetBody(m).Post(mattermostWebhook)
	return err
}
