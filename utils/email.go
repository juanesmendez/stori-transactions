package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
	"os"
	"stori-transactions/environment"
	"stori-transactions/model"
	"time"
)

type EmailData struct {
	RecipientName string

	ImageURL string

	Mail           string
	Balance        float64
	SummaryByMonth map[time.Month]MonthlyTransactionSummary
	AverageByType  map[model.TransactionType]float64
}

func SendEmail(body string) error {
	from := environment.Email
	password := environment.EmailPassword

	to := []string{
		"", //FIXME Sacar de header de un request
	}

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	subject := "Your transactions summary"
	message := fmt.Sprintf("To: %s\r\nSubject: %s\r\nContent-Type: text/html\r\n\r\n%s", to, subject, body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, []byte(message))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Email Sent Successfully!")

	return nil
}

func StartEmail(transactions []model.Transaction) {
	data := EmailData{
		RecipientName: "John Doe", //FIXME De donde sacar el nombre?

		//FIXME
		ImageURL: "https://www.storicard.com/static/images/thumbnail_storicard_small.png",

		Mail:           "some@email.com", //FIXME: Traerlo del header
		Balance:        GetBalance(transactions),
		SummaryByMonth: GroupTransactionsByMonth(transactions),
		AverageByType:  AverageByType(transactions),
	}

	htmlBytes, err := os.ReadFile("email_template.html")
	if err != nil {
		fmt.Println("error reading HTML file:", err)
		return
	}

	htmlTemplate := string(htmlBytes)

	tmpl, err := template.New("emailTemplate").Parse(htmlTemplate)
	if err != nil {
		fmt.Println("error parsing html template:", err)
		return
	}

	var tplBuffer bytes.Buffer
	err = tmpl.Execute(&tplBuffer, data)
	if err != nil {
		fmt.Println("error executing template:", err)
		return
	}

	emailBody := tplBuffer.String()

	err = SendEmail(emailBody)
	if err != nil {
		fmt.Println("error sending email:", err)
		return
	}

	fmt.Println("email sent successfully.")
}
