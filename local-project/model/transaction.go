package model

import "time"

type Transaction struct {
	ID    uint64          `json:"id"`
	Date  time.Time       `json:"date"`
	Value float64         `json:"transaction"`
	Type  TransactionType `json:"type"`
}

type TransactionType string

const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"
)
