package types

//ResponseError default response structure
type ResponseError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
