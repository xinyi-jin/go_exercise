package mail

import (
	"github.com/jordan-wright/email"
)

var EmailSingleton = &email.Email{}

func init() {
	EmailSingleton = email.NewEmail()
	EmailSingleton.From = "Jordan Wright <iezhuhe@163.com>"
	EmailSingleton.To = []string{"1747402195@qq.com"}
	// EmailSingleton.Bcc = []string{"test_bcc@example.com"}
	// EmailSingleton.Cc = []string{"test_cc@example.com"}
	EmailSingleton.Subject = "Awesome Subject"
	EmailSingleton.Text = []byte("Text Body is, of course, supported!")
	EmailSingleton.HTML = []byte("<h1>Fancy HTML is supported, too!</h1>")
}
