package responses

type Result[T any] struct {
	IsSuccessful bool   `json:"is_successful"`
	Data         *T     `json:"data"`
	ErrorMessage string `json:"error_message"`
}
