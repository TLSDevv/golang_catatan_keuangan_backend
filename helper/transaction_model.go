package helper

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/domain"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

func ToTransaction(transaction web.TransactionCreateRequest) domain.Transaction {
	return domain.Transaction{
		UserId:          transaction.UserId,
		NameTransaction: transaction.NameTransaction,
		TypeTransaction: transaction.TypeTransaction,
		CategoryId:      transaction.CategoryId,
		Amount:          transaction.Amount,
		Description:     transaction.Description,
	}
}

func ToTransactionResponse(transaction domain.Transaction) web.TransactionResponse {
	return web.TransactionResponse{
		Id:              transaction.Id,
		UserId:          transaction.UserId,
		NameTransaction: transaction.NameTransaction,
		TypeTransaction: transaction.TypeTransaction,
		CategoryId:      transaction.CategoryId,
		Amount:          transaction.Amount,
		Description:     transaction.Description,
	}
}

func ToTransactionResponses(transactions []domain.Transaction) (transactionResponse []web.TransactionResponse) {
	var transactionResponses []web.TransactionResponse

	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, ToTransactionResponse(transaction))
	}
	return
}
