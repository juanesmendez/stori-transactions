package main

import (
	"fmt"
	"stori-transactions/utils"
)

func main() {

	transactions := utils.ReadTransactionsFromCsv("transactions.csv")

	for i, transaction := range transactions {
		fmt.Printf("Transaction #%d: %v\n", i, transaction)
	}

	fmt.Println("--------- SUMMARY ---------")
	fmt.Printf("Balance: %f\n", utils.GetBalance(transactions))
	fmt.Printf("Transactions and average value per month: %v\n", utils.GroupTransactionsByMonth(transactions))
	fmt.Printf("Average by type: %v\n", utils.AverageByType(transactions))

	utils.StartEmail(transactions)
}
