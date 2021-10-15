package email

import (
	"fmt"
	"net/smtp"
)

const (
	mailNoReply = "noreply@icosmonaut.xyz"
	smtpServer  = "localhost:25"
)

func SendMail(to, subject, msg string) error {
	body := fmt.Sprintf("To: %s\r\n"+
		"Subject: %s\r\n"+
		"\r\n"+
		"%s\r\n", to, subject, msg)
	return smtp.SendMail(smtpServer, nil, mailNoReply, []string{to}, []byte(body))
}
