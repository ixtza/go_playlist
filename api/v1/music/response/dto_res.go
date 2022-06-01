package response

type MusicResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type MusicResponseSuccess struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
