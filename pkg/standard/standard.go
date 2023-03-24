package standard

import "time"

const (
	ResponseSuccessCode = 1000
	ResponseErrorCode   = 5000
)

const (
	ResponseSuccessMessage = "success"
	ResponseErrorMessage   = "error"
)

const (
	AuthKey = "AUTH|%s"
)
var Now = time.Now