package handler

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errors.New(message))
}