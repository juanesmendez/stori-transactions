package utils

import (
	"stori-transactions/model"
	"time"
)

func GetBalance(transactions []model.Transaction) float64 {
	var balance float64

	for _, transaction := range transactions {
		balance += transaction.Value
	}

	return balance
}

func GroupTransactionsByMonth(transactions []model.Transaction) map[time.Month]MonthlyTransactionSummary {
	transSummaryByMonth := make(map[time.Month]MonthlyTransactionSummary)

	for _, transaction := range transactions {
		if _, ok := transSummaryByMonth[transaction.Date.Month()]; !ok {
			transSummaryByMonth[transaction.Date.Month()] = MonthlyTransactionSummary{
				DataByType: make(map[model.TransactionType]SummaryTransactionType),
			}
		}

		monthSummary := transSummaryByMonth[transaction.Date.Month()]
		monthSummary.TransactionsCount++
		monthSummary.Total += transaction.Value

		dataByType := monthSummary.DataByType[transaction.Type]
		dataByType.TransactionsCount++
		dataByType.Total += transaction.Value

		monthSummary.DataByType[transaction.Type] = dataByType

		transSummaryByMonth[transaction.Date.Month()] = monthSummary
	}

	for month, summary := range transSummaryByMonth {
		average := 0.0
		if summary.TransactionsCount > 0 {
			average = summary.Total / float64(summary.TransactionsCount)
		}
		summary.Average = average

		monthlyAverage := 0.0
		for transType, transTypeData := range summary.DataByType {
			if transTypeData.TransactionsCount > 0 {
				monthlyAverage = transTypeData.Total / float64(transTypeData.TransactionsCount)
			}
			transTypeData.Average = monthlyAverage
			summary.DataByType[transType] = transTypeData
		}

		transSummaryByMonth[month] = summary
	}

	return transSummaryByMonth
}

func AverageByType(transactions []model.Transaction) map[model.TransactionType]float64 {
	averageByTransactionType := make(map[model.TransactionType]float64)
	transactionsSummaryByType := make(map[model.TransactionType]SummaryTransactionType)

	for _, transaction := range transactions {
		transactionSummary := transactionsSummaryByType[transaction.Type]
		transactionSummary.TransactionsCount++
		transactionSummary.Total += transaction.Value
		transactionsSummaryByType[transaction.Type] = transactionSummary
	}

	for transactionType, summaryTransactionType := range transactionsSummaryByType {
		average := 0.0
		if summaryTransactionType.TransactionsCount > 0 {
			average = summaryTransactionType.Total / float64(summaryTransactionType.TransactionsCount)
		}
		averageByTransactionType[transactionType] = average
	}

	return averageByTransactionType
}

type MonthlyTransactionSummary struct {
	TransactionsCount int
	Total             float64
	Average           float64
	DataByType        map[model.TransactionType]SummaryTransactionType
}

// FIXME Mover esta struct de aquí
type SummaryTransactionType struct {
	TransactionsCount int
	Total             float64
	Average           float64
}
