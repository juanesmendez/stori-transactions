package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"log"
	"stori-transactions/model"
	"stori-transactions/utils"
	"strings"
)

func handler(ctx context.Context, s3Event events.S3Event) {
	log.Println("entered lambda function handler...")

	sess := session.Must(session.NewSession())

	if len(s3Event.Records) > 0 {
		log.Println("received s3 event record...")
		record := s3Event.Records[0]

		s3Object := record.S3
		bucketName := s3Object.Bucket.Name
		objectKey := s3Object.Object.Key

		content, err := getFileContentFromS3(sess, bucketName, objectKey)
		if err != nil {
			log.Printf("error downloading file from S3: %v", err)
			return
		}

		recordsWithHeader, err := parseCSVContent(content)
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

func getFileContentFromS3(sess *session.Session, bucketName, objectKey string) (string, error) {
	svc := s3.New(sess)

	input := &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	}

	result, err := svc.GetObject(input)
	if err != nil {
		return "", err
	}

	contentBytes, err := io.ReadAll(result.Body)
	if err != nil {
		return "", err
	}

	return string(contentBytes), nil
}

func parseCSVContent(content string) ([][]string, error) {
	reader := csv.NewReader(strings.NewReader(content))

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func printTransactionsData(transactions []model.Transaction) {
	for i, transaction := range transactions {
		fmt.Printf("Transaction #%d: %v\n", i, transaction)
	}

	fmt.Println("--------- SUMMARY ---------")
	fmt.Printf("Balance: %f\n", utils.GetBalance(transactions))
	fmt.Printf("Transactions and average value per month: %v\n", utils.GroupTransactionsByMonth(transactions))
	fmt.Printf("Average by type: %v\n", utils.AverageByType(transactions))
}

func main() {
	lambda.Start(handler)
}
