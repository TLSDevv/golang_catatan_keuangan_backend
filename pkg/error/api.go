package error

type APIError struct {
	StatusCode       int
	Code             string
	Message          string
	LocalizedMessage map[string]string
}

func (api APIError) Error() string {
	return api.Message
}
