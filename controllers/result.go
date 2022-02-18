package controllers

import (
	"github.com/gin-gonic/gin"
)

func resultData(returnData interface{}) gin.H {
	return gin.H{
		"code": 0,
		"data": returnData,
	}
}

func errorText(code int, text string) gin.H {
	return gin.H{
		"code": code,
		"text": text,
	}
}
