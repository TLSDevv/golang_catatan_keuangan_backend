package entities

import (
	"errors"
	"time"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/pkg"
)

var (
	ErrUserIDRequired           = errors.New("user_id is required")
	ErrTransactionNameRequired  = errors.New("transaction_name is required")
	ErrCategoryRequired         = errors.New("category is required")
	ErrTransactionTypeIsInvalid = errors.New("transaction_type is invalid")
	ErrAmountRequired           = errors.New("amount is required")
	ErrNoDataFound              = errors.New("No data found")
	ErrTransactionAtRequired    = errors.New("transaction_at is invalid")

	TransactionDomainErrors = []error{
		ErrUserIDRequired,
		ErrTransactionNameRequired,
		ErrCategoryRequired,
		ErrTransactionTypeIsInvalid,
		ErrAmountRequired,
		ErrNoDataFound,
		ErrTransactionAtRequired,
	}
)

type (
	TransactionInput struct {
		TransactionName string `json:"transaction_name"`
		Category        string `json:"category"`
		TransactionType int    `json:"transaction_type"`
		Amount          int    `json:"amount"`
		TransactionAt   string `json:"transaction_at"`
	}

	CountResult struct {
		Total int `json:"total"`
	}

	Transaction struct {
		ID              int       `json:"id"`
		UserID          int       `json:"user_id"`
		TransactionName string    `json:"transaction_name"`
		Category        string    `json:"category"`
		TransactionType int       `json:"transaction_type"`
		Amount          int       `json:"amount"`
		TransactionAt   time.Time `json:"transaction_at"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		DeletedAt       time.Time `json:"deleted_at"`
	}
)

func NewTransaction(userID int, transactionName string, category string, transactionType int, amount int, transactionAt time.Time) (Transaction, error) {
	transaction := Transaction{
		UserID:          userID,
		TransactionName: transactionName,
		Category:        category,
		TransactionType: transactionType,
		Amount:          amount,
		TransactionAt:   transactionAt,
	}
	err := transaction.Validate()
	if err != nil {
		return Transaction{}, err
	}

	return transaction, nil
}

func (t *Transaction) Update(ti TransactionInput) error {
	if len(ti.TransactionName) != 0 {
		t.TransactionName = ti.TransactionName
	}
	if len(ti.Category) != 0 {
		t.Category = ti.Category
	}
	if ti.TransactionType == 0 || ti.TransactionType == 1 {
		t.TransactionType = ti.TransactionType
	}
	if ti.Amount != 0 {
		t.Amount = ti.Amount
	}
	if len(ti.TransactionAt) != 0 {
		tAt, err := pkg.StringDateToDateTime(ti.TransactionAt)
		if err != nil {
			return err
		}
		t.TransactionAt = *tAt
	}

	return nil
}

func (ti TransactionInput) Validate() []string {
	var errors []string

	if len(ti.TransactionName) == 0 {
		errors = append(errors, ErrTransactionNameRequired.Error())
	}
	if len(ti.Category) == 0 {
		errors = append(errors, ErrCategoryRequired.Error())
	}
	if ti.TransactionType < 0 || ti.TransactionType > 1 {
		errors = append(errors, ErrTransactionTypeIsInvalid.Error())
	}
	if ti.Amount == 0 {
		errors = append(errors, ErrAmountRequired.Error())
	}
	if len(ti.TransactionAt) == 0 {
		errors = append(errors, ErrTransactionAtRequired.Error())
	}

	if errors != nil {
		return errors
	}

	return nil
}

func (t Transaction) Validate() error {
	if len(t.TransactionName) == 0 {
		return ErrTransactionNameRequired
	}
	if len(t.Category) == 0 {
		return ErrCategoryRequired
	}
	if t.TransactionType < 0 || t.TransactionType > 1 {
		return ErrTransactionTypeIsInvalid
	}
	if t.Amount == 0 {
		return ErrAmountRequired
	}
	return nil
}

func (t Transaction) Delete() {
	t.DeletedAt = time.Now()
}
