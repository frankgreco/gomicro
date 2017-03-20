package utils

type Error interface {
	error
	Status() int
}

type StatusError struct {
	Code int
	Err  error
}

type JsonErr struct {
	Code int `json:"code"`
	Msg  string `json:"msg"`
}

func (se StatusError) Error() string {
	return se.Err.Error()
}

func (se StatusError) Status() int {
	return se.Code
}
