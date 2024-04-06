package dto

type Response struct {
	Message   string      `json:"message"`
	Body      interface{} `json:"body,omitempty"`
	Error     string      `json:"error,omitempty"`
}
