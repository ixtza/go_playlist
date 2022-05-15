package response

type AuthResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type AuthResponseSuccess struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
