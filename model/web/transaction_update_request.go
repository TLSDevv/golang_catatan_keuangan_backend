package web

type TransactionUpdateRequest struct {
	Id              uint8   `validate:"required" json:"id"`
	UserId          uint8   `validate:"required" json:"user_id"`
	NameTransaction string `validate:"required,max=50,min=1" json:"name_transaction"`
	TypeTransaction string `validate:"required,max=10,min=1" json:"type_transaction"`
	CategoryId      uint8   `validate:"required" json:"category_id"`
	Amount          int64  `validate:"required" json:"nominal"`
	Description     string `validate:"required,max=250,min=1" json:"description"`
}
