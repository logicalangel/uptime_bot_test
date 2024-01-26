package dto

type Response struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseError struct {
	Status  bool         `json:"status"`
	Message string       `json:"message"`
	Fields  []FieldError `json:"fields,omitempty"`
}

type FieldError struct {
	Field string `json:"field"`
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

func (e ResponseError) Error() string {
	return e.Message
}
