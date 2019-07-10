package libraries

import "github.com/labstack/echo"

func ToJson(c echo.Context, code int, message string, data interface{}) error {
	return c.JSON(code, H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
