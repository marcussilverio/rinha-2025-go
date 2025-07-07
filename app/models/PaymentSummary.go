package models

type PaymentSummary struct {
	UUID      string  `json:"UUID"`
	IsDefault bool    `json:"isDefault"`
	Amount    float64 `json:"amount"`
}

const (
	TableName        = "payment_summary"
	CreateTableQuery = `CREATE TABLE IF NOT EXISTS payment_summary (
    		UUID VARCHAR(36) PRIMARY KEY,
    		isDefault BOOLEAN NOT NULL,
    		amount FLOAT NOT NULL
    		);`
)
