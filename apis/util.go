package apis

import (
	"fmt"
	"net/http"
)

const (
	new      = "new"
	pending  = "pending"
	resolved = "resolved"
	accepted = "accepted"
	rejected = "rejected"
)

func SendHttp(w http.ResponseWriter, v ResponseInterface) {
	v.Send(w)
}

func SendErr(w http.ResponseWriter, statusCode int, err error) {
	er := ErrResponse{
		HTTPStatusCode: statusCode,
		ErrorText:      fmt.Sprint(err),
	}
	er.Send(w)
}
