package exception

type NotFoundError struct {
	Error string
}

func NewNotFoundError(err string) NotFoundError {
	return NotFoundError{Error: err}
}

// type BadRequestError struct {
// 	Error string
// }

// func NewBadRequestError(err string) BadRequestError {
// 	return BadRequestError{Error: err}
// }
