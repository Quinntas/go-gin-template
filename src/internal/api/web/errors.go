package web

import "github.com/gin-gonic/gin"

type HttpError struct {
	status  int
	message string
	body    gin.H
}

func NewHttpError(status int, message string, body gin.H) *HttpError {
	return &HttpError{status, message, body}
}

func (he *HttpError) Error() string {
	return he.message
}

func (he *HttpError) serialize(c *gin.Context) {
	he.body["message"] = he.message
	JsonResponse(c, he.status, he.body)
}

func NotFound() *HttpError {
	return NewHttpError(404, "Not Found", gin.H{})
}

func InternalServerError() *HttpError {
	return NewHttpError(500, "Internal Server Error", gin.H{})
}
