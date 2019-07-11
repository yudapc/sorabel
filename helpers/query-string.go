package helpers

import (
	"strconv"

	"github.com/labstack/echo"
)

func QueryString(context echo.Context) (int, int, int, string) {
	page, _ := strconv.Atoi(context.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit, _ := strconv.Atoi(context.QueryParam("limit"))
	if limit < 1 {
		limit = 5
	}
	orderBy := context.QueryParam("orderBy")
	if orderBy == "" {
		orderBy = "id"
	}
	sortBy := context.QueryParam("sortBy")
	if sortBy == "" {
		sortBy = "asc"
	}
	offset := (page - 1) * limit
	order := orderBy + " " + sortBy
	return page, limit, offset, order
}
