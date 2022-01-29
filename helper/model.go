package helper

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/dto"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/entity"
)

func UserToDTO(u entity.User) dto.User {
	return dto.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Fullname: u.Fullname,
	}
}

func UsersToDTO(u []entity.User) []dto.User {
	users := []dto.User{}

	for _, value := range u {
		user := UserToDTO(value)
		users = append(users, user)
	}

	return users
}

func UserDTOToUser(u dto.User) entity.User {
	return entity.User{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Fullname: u.Fullname,
	}
}

func TransactionToDTO(t entity.Transaction) dto.Transaction {
	return dto.Transaction{
		ID:            t.ID,
		TrcName:       t.TrcName,
		Amount:        t.Amount,
		CreatedAt:     t.CreatedAt,
		TransactionAt: t.TransactionAt,
		TrcType:       t.TrcType,
	}
}

func TransactionsToDTO(t []entity.Transaction) []dto.Transaction {
	transactions := []dto.Transaction{}

	for _, value := range t {
		transaction := TransactionToDTO(value)
		transactions = append(transactions, transaction)
	}

	return transactions
}
