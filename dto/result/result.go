package dto

type SuccessResult struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type ErrorResult struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
