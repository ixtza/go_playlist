package response

type PlaylistResponse struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type PlaylistResponseSuccess struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
