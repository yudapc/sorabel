package helpers

import "github.com/labstack/echo"

func ToJson(context echo.Context, code int, message string, data interface{}) error {
	return context.JSON(code, H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}

func ToJsonBadRequest(context echo.Context, data interface{}) error {
	return context.JSON(400, H{
		"code":    400,
		"message": "failed",
		"data":    data,
	})
}
