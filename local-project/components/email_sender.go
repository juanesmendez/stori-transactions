package components

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/smtp"
	"os"
	"stori-transactions/environment"
	"stori-transactions/model"
)

type EmailSender interface {
	StartEmail(transactions []model.Transaction, calculator TransactionSummaryCalculator)
	SendEmail(body string) error
}

type EmailSenderImpl struct {
}

func NewEmailSenderImpl() EmailSender {

	return &EmailSenderImpl{}
}

func (emailSender *EmailSenderImpl) StartEmail(transactions []model.Transaction, calculator TransactionSummaryCalculator) {
	data := model.EmailData{
		ImageURL:       "https://www.storicard.com/static/images/thumbnail_storicard_small.png",
		Mail:           environment.ToEmail,
		Balance:        calculator.GetBalance(transactions),
		SummaryByMonth: calculator.GroupTransactionsByMonth(transactions),
		AverageByType:  calculator.AverageByType(transactions),
	}

	htmlBytes, err := os.ReadFile("email_template.html")
	if err != nil {
		log.Printf("error reading HTML file: %s", err.Error())
		return
	}

	htmlTemplate := string(htmlBytes)

	tmpl, err := template.New("emailTemplate").Parse(htmlTemplate)
	if err != nil {
		log.Printf("error parsing html template: %s", err.Error())
		return
	}

	var tplBuffer bytes.Buffer
	err = tmpl.Execute(&tplBuffer, data)
	if err != nil {
		log.Printf("error executing template: %s", err.Error())
		return
	}

	emailBody := tplBuffer.String()

	err = emailSender.SendEmail(emailBody)
	if err != nil {
		log.Printf("error sending email: %s", err.Error())
		return
	}

	log.Println("email sent successfully.")
}

func (emailSender *EmailSenderImpl) SendEmail(body string) error {
	from := environment.Email
	password := environment.EmailPassword

	to := []string{
		environment.ToEmail,
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Your transactions summary"
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html\r\n\r\n%s", to, subject, body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		log.Printf("error sending email: %s", err.Error())
		return err
	}
	log.Println("Email Sent Successfully!")

	return nil
}
