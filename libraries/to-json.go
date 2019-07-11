package libraries

import "github.com/labstack/echo"

func ToJson(context echo.Context, code int, message string, data interface{}) error {
	return context.JSON(code, H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
