package service

import (
	"github.com/TLSDevv/golang_catatan_keuangan_backend/api/repository"
	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

func NewService(t repository.TransactionRepository) TransactionServiceInterface {
	return &transactionService{
		transactionRepository: t,
	}
}

func (s *transactionService) ListTransaction(limit int, page int) ([]web.TransactionResponse, error) {
	return s.transactionRepo.List(limit, page)
}

func (s *transactionService) GetTransaction(id string) (web.TransactionResponse, error) {
	return s.transactionRepo.GetByID(id)
}

func (s *transactionService) CreateTransaction(t web.TransactionResponse) error {

}
