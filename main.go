package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"log"
	"stori-transactions/model"
	"stori-transactions/utils"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, s3Event events.S3Event) {
	log.Println("entered lambda function handler...")

	sess := session.Must(session.NewSession())

	if len(s3Event.Records) > 0 {
		log.Println("received s3 event record...")
		record := s3Event.Records[0]

		s3Object := record.S3
		bucketName := s3Object.Bucket.Name
		objectKey := s3Object.Object.Key

		content, err := utils.GetFileContentFromS3(sess, bucketName, objectKey)
		if err != nil {
			log.Printf("error downloading file from S3: %v", err)
			return
		}

		recordsWithHeader, err := utils.ParseCSVContent(content)
		if err != nil {
			log.Printf("error parsing CSV content: %v", err)
			return
		}
		log.Println("parsed csv content successfully")
		records := recordsWithHeader[1:]
		transactions := utils.ReadTransactions(records)
		printTransactionsData(transactions)

		log.Println("starting to send email...")
		utils.StartEmail(transactions)
	} else {
		log.Println("no S3 events found in the event payload")
	}
}

func printTransactionsData(transactions []model.Transaction) {
	for i, transaction := range transactions {
		log.Printf("Transaction #%d: %v\n", i, transaction)
	}

	log.Println("--------- SUMMARY ---------")
	log.Printf("Balance: %f\n", utils.GetBalance(transactions))
	log.Printf("Transactions and average value per month: %v\n", utils.GroupTransactionsByMonth(transactions))
	log.Printf("Average by type: %v\n", utils.AverageByType(transactions))
}
