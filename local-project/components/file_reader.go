package components

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

type FileReader interface {
	ReadTransactionsFromCsv(fileName string) []model.Transaction
}

type FileReaderImpl struct{}

func NewFileReaderImpl() FileReader {

	return &FileReaderImpl{}
}

func (fileReader *FileReaderImpl) ReadTransactionsFromCsv(fileName string) []model.Transaction {
	records, err := fileReader.getRecords(fileName)

	if err != nil {
		return []model.Transaction{}
	}

	var transactions []model.Transaction

	for i, record := range records {
		if len(record) != 3 {
			continue
		}
		transactionId, err := fileReader.getTransactionId(record, i)

		if err != nil {
			continue
		}

		date, err := fileReader.getDate(record, i)

		if err != nil {
			continue
		}

		transactionType, err := fileReader.getTransactionType(record, i)

		if err != nil {
			continue
		}

		transactionValue, err := fileReader.getTransactionValue(record, i)

		if err != nil {
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

	return transactions
}

func (fileReader *FileReaderImpl) getRecords(fileName string) ([][]string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Printf("error getting current working directory: %s", err.Error())
		return [][]string{}, err
	}

	absolutePath := filepath.Join(currentDir, fileName)
	file, err := os.Open(absolutePath)

	if err != nil {
		log.Printf("error reading the transactions csv file: %s", err.Error())
		return [][]string{}, err
	}

	defer file.Close()

	reader := csv.NewReader(file)

	recordsWithHeader, err := reader.ReadAll()

	if err != nil {
		log.Printf("error reading transaction records from csv file: %s", err.Error())
		return [][]string{}, err
	}

	return recordsWithHeader[1:], nil
}

func (fileReader *FileReaderImpl) getTransactionId(record []string, rowNum int) (uint64, error) {
	transactionIdStr := record[0]
	transactionId, err := strconv.ParseUint(transactionIdStr, 10, 64)

	if err != nil {
		log.Printf("error parsing transaction's 'id' for record #%d in csv file:%s\n", rowNum, err.Error())
		return transactionId, err
	}

	return transactionId, nil
}

func (fileReader *FileReaderImpl) getDate(record []string, rowNum int) (time.Time, error) {
	dateStr := record[1]
	dateParts := strings.Split(dateStr, "/")

	if len(dateParts) != 2 {
		return time.Time{}, fmt.Errorf("'date' field of record #%d is  badly formatted: %s", rowNum, dateParts)
	}

	month, err := strconv.Atoi(dateParts[0])

	if err != nil {
		log.Printf("error parsing transaction's month for record #%d in csv file\n", rowNum)
		return time.Time{}, err
	}

	day, err := strconv.Atoi(dateParts[1])

	if err != nil {
		log.Printf("error parsing transaction's day of month for record #%d in csv file\n", rowNum)
		return time.Time{}, err
	}

	date := time.Date(time.Now().Year(), time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return date, nil
}

func (fileReader *FileReaderImpl) getTransactionType(record []string, rowNum int) (model.TransactionType, error) {
	transactionValueStr := record[2]
	transactionSign := string(transactionValueStr[0])

	if transactionSign == "+" {
		return model.Credit, nil
	}

	if transactionSign == "-" {
		return model.Debit, nil
	}

	return "", fmt.Errorf("'transaction' field of record #%d is  badly formatted: %s", rowNum, transactionValueStr)
}

func (fileReader *FileReaderImpl) getTransactionValue(record []string, rowNum int) (float64, error) {
	transactionValueStr := record[2]
	transactionValue, err := strconv.ParseFloat(transactionValueStr, 64)

	if err != nil {
		log.Printf("error parsing transaction's value for record #%d in csv file\n", rowNum)
		return transactionValue, err
	}

	return transactionValue, nil
}
