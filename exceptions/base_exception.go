package exceptions

import "net/http"

type BaseException struct {
	Id         string `json:"id"`
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode uint16 `json:"statusCode"`
}

func NewBaseException(status int, id string, msg string) *BaseException {
	b := new(BaseException)

	b.Id = id
	b.Message = msg
	b.StatusCode = uint16(status)
	b.Error = http.StatusText(status)

	return b
}
