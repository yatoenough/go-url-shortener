package response

import (
	"net/http"
	"time"
)

type Response struct {
	StatusCode int       `json:"statusCode"`
	Error      string    `json:"error,omitempty"`
	Timestamp  time.Time `json:"timestamp"`
}

func OK() Response {
	return Response{
		StatusCode: http.StatusOK,
		Timestamp:  time.Now(),
	}
}

func Error(status int, msg string) Response {
	return Response{
		StatusCode: status,
		Error:      msg,
		Timestamp:  time.Now(),
	}
}
