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

func Notification(err string) {

	var m Message
	var data Attachment
	// err := os.Args[1]

	// if len(os.Args) < 2 {
	// 	fmt.Println("Please provide two input arguments.")
	// 	return
	// }

	if err == "nil" {
		data.Color = "#B0fc38"
		data.Text = "Success"
	} else {
		data.Color = "#FF0000"
		data.Text = err
	}
	m.Attachments = append(m.Attachments, data)

	resty.New().R().SetBody(m).Post("http://mattermost.sytes.net/hooks/5bzes4jw5b8dbneepimad4hb4o")

}
