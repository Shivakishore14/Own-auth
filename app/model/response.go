package model

//WebResponse is the structure used for web api response
type WebResponse struct {
	Message string      `json:"message"`
	Error   error       `json:"error"`
	Data    interface{} `json:"data"`
}
