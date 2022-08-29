package responses

type SerendipiaResponse struct {
	Status  int                    `json:"satatus"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
