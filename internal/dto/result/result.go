package dto

type SuccessResult struct {
	// ResponseCode string      `json:"code"`
	Status int         `json:"code"`
	Data   interface{} `json:"data"`
}
type SuccessReauth struct {
	// ResponseCode string      `json:"code"`
	Status int         `json:"code"`
	Data   interface{} `json:"message"`
}

type ErrorResult struct {
	// ResponseCode string `json:"code"`
	Status  int    `json:"code"`
	Message string `json:"message"`
}
type ErrorResultDto struct {
	// ResponseCode string `json:"code"`
	Status  int    `json:"code"`
	Message string `json:"message"`
	Error   error  `json:"error"`
}
