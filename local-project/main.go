package main

import (
	"log"
	"stori-transactions/utils"
)

func main() {

	transactions := utils.ReadTransactionsFromCsv("transactions.csv")

	for i, transaction := range transactions {
		log.Printf("Transaction #%d: %v\n", i, transaction)
	}

	log.Println("--------- SUMMARY ---------")
	log.Printf("Balance: %f\n", utils.GetBalance(transactions))
	log.Printf("Transactions and average value per month: %v\n", utils.GroupTransactionsByMonth(transactions))
	log.Printf("Average by type: %v\n", utils.AverageByType(transactions))

	utils.StartEmail(transactions)
}