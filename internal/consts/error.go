package consts

import "errors"

var (
	ErrInternalError  = errors.New("internal_error")
	ErrApiNotfound    = errors.New("api_notfound")
	ErrBadBodyRequest = errors.New("request_body_malformed")
)
