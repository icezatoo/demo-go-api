package dto

// https://blog.depa.do/post/gin-validation-errors-handling

type ResponseSucessful struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

type ResponseError struct {
	StatusCode int         `json:"statusCode"`
	Error      interface{} `json:"error"`
}
