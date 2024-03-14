package utils

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"stori-transactions/model"
	"strconv"
	"strings"
	"time"
)

func ReadTransactions(fileName string) []model.Transaction {
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("error getting current working directory:", err)
		return []model.Transaction{}
	}

	absolutePath := filepath.Join(currentDir, fileName)

	file, err := os.Open(absolutePath)

	if err != nil {
		log.Fatal("error reading the transactions csv file: ", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	recordsWithHeader, err := reader.ReadAll()

	if err != nil {
		fmt.Println("error reading transaction records from csv file")
	}

	records := recordsWithHeader[1:]
	var transactions []model.Transaction

	for i, record := range records {
		if len(record) == 3 {
			transactionIdStr := record[0]
			transactionId, err := strconv.ParseUint(transactionIdStr, 10, 64)

			if err != nil {
				fmt.Printf("error parsing transaction's 'id' for record #%d in csv file:%s\n", i, err.Error())
				continue
			}

			dateStr := record[1]
			dateParts := strings.Split(dateStr, "/")

			var date time.Time
			if len(dateParts) == 2 {
				month, err := strconv.Atoi(dateParts[0])

				if err != nil {
					fmt.Printf("error parsing transaction's month for record #%d in csv file\n", i)
					continue
				}

				day, err := strconv.Atoi(dateParts[1])

				if err != nil {
					fmt.Printf("error parsing transaction's day of month for record #%d in csv file\n", i)
					continue
				}

				date = time.Date(time.Now().Year(), time.Month(month), day, 0, 0, 0, 0, time.UTC)
			}

			transactionValueStr := record[2]
			transSign := string(transactionValueStr[0])

			var transactionType model.TransactionType
			if transSign == "+" {
				transactionType = model.Credit
			} else if transSign == "-" {
				transactionType = model.Debit
			} else {
				continue //FIXME
			}

			transactionValue, err := strconv.ParseFloat(transactionValueStr, 64)

			if err != nil {
				fmt.Printf("error parsing transaction's value for record #%d in csv file\n", i)
				continue
			}

			transaction := model.Transaction{
				ID:    transactionId,
				Date:  date,
				Value: transactionValue,
				Type:  transactionType,
			}

			transactions = append(transactions, transaction)
		}

	}

	return transactions
}
