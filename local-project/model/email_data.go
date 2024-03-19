package model

import (
	"time"
)

type EmailData struct {
	ImageURL       string
	Mail           string
	Balance        float64
	SummaryByMonth map[time.Month]MonthlyTransactionSummary
	AverageByType  map[TransactionType]float64
}
