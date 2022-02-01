package handler

// type ITransactionHandler interface {
// 	FindAll(w http.ResponseWriter, r *http.Request)
// }

// type TransactionHandler struct {
// 	ts service.ITransactionService
// }

// func NewTransactionHandler(r *mux.Router, ts service.ITransactionService) ITransactionHandler {
// 	th := TransactionHandler{
// 		ts: ts,
// 	}

// 	r.HandleFunc("/transactions", th.FindAll).Methods("GET")

// 	return th
// }

// func (th TransactionHandler) FindAll(w http.ResponseWriter, r *http.Request) {
// 	result, err := th.ts.FindAll(r.Context())

// 	if err != nil {
// 		SendNoData(w, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	SendWithData(w, http.StatusOK, "", result)
// 	return
// }
