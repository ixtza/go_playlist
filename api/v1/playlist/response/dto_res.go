package response

type ResponseError struct {
	StatusCode string      `json:"status"`
	Message    interface{} `json:"message"`
}

type ResponseSuccess struct {
	StatusCode string      `json:"status"`
	Message    interface{} `json:"data"`
}

type ResponseCreated struct {
	StatusCode string      `json:"status"`
	Data       interface{} `json:"data"`
}
