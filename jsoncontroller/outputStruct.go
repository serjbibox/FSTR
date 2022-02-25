package jsoncontroller

type response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewResponse() response {
	return response{}
}

const (
	new      = "new"
	pending  = "pending"
	resolved = "resolved"
	accepted = "accepted"
	rejected = "rejected"
)

var Status string = new
