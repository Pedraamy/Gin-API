package dto

type Response struct {
	Response interface{} `json:"message"`
}

type Success struct {
	ID interface{} `json:"ID"`
}
