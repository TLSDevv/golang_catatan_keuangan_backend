package web

type TransactionUpdateRequest struct {
	UserId          uint8  `json:"user_id"`
	NameTransaction string `json:"name_transaction"`
	TypeTransaction string `json:"type_transaction"`
	CategoryId      uint8  `json:"category_id"`
	Amount          int64  `json:"nominal"`
	Description     string `json:"description"`
}
