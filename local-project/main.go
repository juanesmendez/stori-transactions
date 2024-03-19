package main

import (
	"log"
	"stori-transactions/components"
	"stori-transactions/model"
)

func main() {
	fileReader := components.NewFileReaderImpl()
	emailSender := components.NewEmailSenderImpl()
	transSummaryCalc := components.NewTransactionSummaryCalculatorImpl()

	transactions := fileReader.ReadTransactionsFromCsv("transactions.csv")
	printTransactionsData(transactions, transSummaryCalc)
	emailSender.StartEmail(transactions, transSummaryCalc)
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
