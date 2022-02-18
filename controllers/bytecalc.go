package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

func CalcByte(c *gin.Context) {
	unit := c.Query("unit")
	valueStr := c.Query("value")

	value, err := cast.ToUint64E(valueStr)
	if err != nil {
		c.JSON(http.StatusOK, errorText(1, "value must be a number"))
		return
	}

	switch unit {
	case "bit":
	case "byte":
		value *= 8
	case "kilobyte":
		value *= 8192
	case "megabyte":
		value *= 8388608
	case "gigabyte":
		value *= 8589934592
	case "terabyte":
		value *= 8796093022208
	default:
		c.JSON(http.StatusOK, errorText(1, "invalid unit"))
		return
	}

	c.JSON(http.StatusOK, resultData(map[string]float64{
		"bit":      float64(value),
		"byte":     float64(value) / 8,
		"kilobyte": float64(value) / 8192,
		"megabyte": float64(value) / 8388608,
		"gigabyte": float64(value) / 8589934592,
		"terabyte": float64(value) / 8796093022208,
	}))
}
