package web

import (
	"github.com/gin-gonic/gin"
	"github.com/quinntas/go-gin-template/src/internal/utils"
)

type HttpError struct {
	status  int
	message string
	body    utils.H
}

func NewHttpError(status int, message string, body utils.H) *HttpError {
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
	return NewHttpError(404, "Not Found", utils.H{})
}

func InternalServerError() *HttpError {
	return NewHttpError(500, "Internal Server Error", utils.H{})
}

func InvalidRequest() *HttpError {
	return NewHttpError(400, "Invalid Request", utils.H{})
}
