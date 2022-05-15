package response

type UserResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type UserResponseSuccess struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
