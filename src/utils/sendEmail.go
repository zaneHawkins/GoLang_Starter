package utils

import (
	"fmt"
	"github.com/friendsofgo/errors"
	"net/smtp"
	"os"
	"strings"
)

type SMTPServer struct {
	Host string
	Port string
}

type Mail struct {
	To      []string
	Subject string
	Body    string
}

type Sender struct {
	Email    string
	Password string
}

func (s *SMTPServer) Address() string {
	return s.Host + ":" + s.Port
}

func SendEmail(email *Mail) error {

	smtpServer, sender, err := initializeMailSender()
	if err != nil {
		return err
	}

	message := buildMessage(sender, email)

	fmt.Println("Send email", string(message))

	auth := smtp.PlainAuth("", sender.Email, sender.Password, smtpServer.Host)

	err = smtp.SendMail(smtpServer.Address(), auth, sender.Email, email.To, message)
	if err != nil {
		return err
	}

	return nil
}

func initializeMailSender() (*SMTPServer, *Sender, error) {
	sender := Sender{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("EMAIL_PASS"),
	}

	if sender.Email == "" || sender.Password == "" {
		return nil, nil, errors.New("Sender Email or Password cannot be found")
	}

	smtpServer := SMTPServer{
		Host: "smtp.gmail.com",
		Port: "587",
	}

	return &smtpServer, &sender, nil
}

func buildMessage(sender *Sender, mail *Mail) []byte {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", sender.Email)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return []byte(msg)
}
