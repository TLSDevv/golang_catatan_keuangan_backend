package service

import (
	"context"

	"github.com/TLSDevv/golang_catatan_keuangan_backend/model/web"
)

type TransactionServiceInterface interface {
	ListTransaction(ctx context.Context, int, page int) []web.TransactionResponse
	GetTransaction(ctx context.Context, idTransaction int) web.TransactionResponse
	CreateTransaction(ctx context.Context, t web.TransactionCreateRequest)
	UpdateTransaction(ctx context.Context, idTransaction int, t web.TransactionCreateRequest)
	DeleteTransaction(ctx context.Context, idTransaction int)
}
