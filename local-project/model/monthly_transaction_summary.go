package model

type MonthlyTransactionSummary struct {
	TransactionsCount int
	Total             float64
	Average           float64
	DataByType        map[TransactionType]SummaryTransactionType
}
