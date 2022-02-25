package jsoncontroller

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewResponse() response {
	return response{}
}
