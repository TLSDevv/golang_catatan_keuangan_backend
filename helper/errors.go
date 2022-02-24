package helper

type Errors struct {
	Err  string
	Data interface{}
}

func (e *Errors) Error() string {
	return e.Err
}

func NewErrors(message string, data interface{}) error {
	return &Errors{
		Err:  message,
		Data: data,
	}
}
