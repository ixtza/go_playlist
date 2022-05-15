package response

type CollaborationResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type CollaborationResponseSuccess struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
