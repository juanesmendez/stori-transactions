package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"stori-transactions/components"
	"stori-transactions/model"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, s3Event events.S3Event) {
	log.Println("entered lambda function handler...")

	sess := session.Must(session.NewSession())
	emailSender := components.NewEmailSenderImpl()
	fileReader := components.NewFileReaderImpl()
	transSummaryCalculator := components.NewTransactionSummaryCalculatorImpl()

	if len(s3Event.Records) > 0 {
		log.Println("received s3 event record...")
		record := s3Event.Records[0]

		s3Object := record.S3
		bucketName := s3Object.Bucket.Name
		objectKey := s3Object.Object.Key

		content, err := fileReader.GetFileContentFromS3(sess, bucketName, objectKey)
		if err != nil {
			log.Printf("error downloading file from S3: %v", err)
			return
		}

		recordsWithHeader, err := fileReader.ParseCSVContent(content)
		if err != nil {
			log.Printf("error parsing CSV content: %v", err)
			return
		}
		log.Println("parsed csv content successfully")
		records := recordsWithHeader[1:]
		transactions := fileReader.ReadTransactions(records)
		printTransactionsData(transactions, transSummaryCalculator)

		log.Println("starting to send email...")
		emailSender.StartEmail(transactions, transSummaryCalculator)
	} else {
		log.Println("no S3 events found in the event payload")
	}
}

func printTransactionsData(transactions []model.Transaction, calculator components.TransactionSummaryCalculator) {
	for i, transaction := range transactions {
		log.Printf("Transaction #%d: %v\n", i, transaction)
	}

	log.Println("--------- SUMMARY ---------")
	log.Printf("Balance: %f\n", calculator.GetBalance(transactions))
	log.Printf("Transactions and average value per month: %v\n", calculator.GroupTransactionsByMonth(transactions))
	log.Printf("Average by type: %v\n", calculator.AverageByType(transactions))
}
