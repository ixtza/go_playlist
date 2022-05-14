package response

type Response struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type ResponseSuccess struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
