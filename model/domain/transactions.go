package domain

import (
	"time"
)

type Transaction struct {
	Id              uint8     `json:"id"`
	UserId          uint8     `json:"user_id"`
	NameTransaction string    `json:"name_transaction"`
	TypeTransaction string    `json:"type_transaction"`
	CategoryId      uint8     `json:"category_id"`
	Amount          int64     `json:"nominal"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DeletedAt       time.Time `json:"deleted_at"`
}

func localTime() time.Time {
	return time.Now().Local()
}

func NewTransaction(userId uint8, nameTransaction string, typeTransaction string, categoryId uint8, amount int64, description string) (*Transaction, error) {
	t := &Transaction{
		Id:              123,
		UserId:          userId,
		NameTransaction: nameTransaction,
		TypeTransaction: typeTransaction,
		CategoryId:      categoryId,
		Amount:          amount,
		Description:     description,
		CreatedAt:       localTime(),
		UpdatedAt:       localTime(),
	}

	return t, nil
}
