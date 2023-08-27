// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql"
)

type Account struct {
	ID        int64        `json:"id"`
	Username  string       `json:"username"`
	Type      string       `json:"type"`
	Document  string       `json:"document"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	Balance   int64        `json:"balance"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Entry struct {
	ID        int64        `json:"id"`
	Amount    int64        `json:"amount"`
	AccountID int64        `json:"account_id"`
	CreatedAt sql.NullTime `json:"created_at"`
}

type Transaction struct {
	ID         int64        `json:"id"`
	Amount     int64        `json:"amount"`
	SenderID   int64        `json:"sender_id"`
	ReceiverID int64        `json:"receiver_id"`
	CreatedAt  sql.NullTime `json:"created_at"`
}
