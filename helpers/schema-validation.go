package helpers

import (
	"github.com/labstack/echo"
	"github.com/xeipuuv/gojsonschema"
)

func SchemaValidation(c echo.Context, filePath string) error {
	schemaLoader := gojsonschema.NewReferenceLoader(GetFile(filePath))
	document := gojsonschema.NewGoLoader(c.Request().Body)
	_, err := gojsonschema.Validate(schemaLoader, document)
	if err != nil {
		panic(err.Error())
		return err
	}
	return nil
}
