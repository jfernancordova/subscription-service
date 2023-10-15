package main

import (
	"sync"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

// Message represents an email message
type Mail struct {
	Domain      string
	Host        string
	Port        int
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
	// to be used for performing the actual sending of the mail
	Wait      *sync.WaitGroup
	MailChan  chan Message
	ErrorChan chan error
	DoneChan  chan bool
}

// Message represents an email message
type Message struct {
	From        string
	FromName    string
	To          string
	Subject     string
	Attachments []string
	Data        any
	DataMap     map[string]any
	Template    string
}

func (m *Mail) sendMail(msg Message, errorChan chan error) error {
	if msg.Template != "" {
		msg.Template = "mail"
	}

	if msg.From == "" {
		msg.From = m.FromAddress
	}

	if msg.FromName == "" {
		msg.FromName = m.FromName
	}

	data := map[string]any{
		"message": msg.Data,
	}
	msg.DataMap = data

	formattedMessage, err := m.buildHTMLMessage(msg)
	if err != nil {
		errorChan <- err
	}

	plainMessage, err := m.buildPlainMessage(msg)
	if err != nil {
		errorChan <- err
	}

	server := mail.NewSMTPClient()
	server.Host = m.Host
	server.Port = m.Port
	server.Username = m.Username
	server.Password = m.Password
	server.Encryption = m.getEncryption(m.Encryption)
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		errorChan <- err
	}

	email := mail.NewMSG()
	email.SetFrom(msg.From).AddTo(msg.To).SetSubject(msg.Subject)
	email.SetBody(mail.TextPlain, plainMessage)
	email.AddAlternative(mail.TextHTML, formattedMessage)

	if len(msg.Attachments) > 0 {
		for _, attachment := range msg.Attachments {
			email.AddAttachment(attachment)
		}
	}

	err = email.Send(smtpClient)
	if err != nil {
		errorChan <- err
	}

	return nil
}

func (m *Mail) buildHTMLMessage(msg Message) (string, error) {
	return "", nil
}

func (m *Mail) buildPlainMessage(msg Message) (string, error) {
	return "", nil
}

func (m *Mail) getEncryption(e string) mail.Encryption {
	switch e {
	case "ssl":
		return mail.EncryptionSSLTLS
	case "tls":
		return mail.EncryptionSTARTTLS
	case "none":
		return mail.EncryptionNone
	default:
		return mail.EncryptionSTARTTLS
	}
}
