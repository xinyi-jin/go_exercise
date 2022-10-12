package mail

import (
	"fmt"
	"net/smtp"
)

// go1.5及之后版本使用
// 实测go1.15版本也可使用
func SendMail() {
	err := EmailSingleton.Send("smtp.163.com:25", smtp.PlainAuth("", "iezhuhe@163.com", "MMSZFTEYAFVBRKNK", "smtp.163.com"))
	if err != nil {
		fmt.Printf("EmailSingleton.Send failed %v", err)
	}
}
