package service_error

type ServiError struct {
	Code int
	Msg  string
}

func NewServiceError() *ServiError {
	return &ServiError{Code: 500, Msg: "Internal Server Error"}
}
