package dto

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

type Params struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func OK() Params {
	return Params{
		Status: StatusOK,
	}
}

func Error(msg string) Params {
	return Params{
		Status: StatusError,
		Error:  msg,
	}
}
