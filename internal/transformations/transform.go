package transformations

import "github.com/labstack/echo"

// Transformation is a shared interface that all transformations must implement to be valid
// where eval() returns a map of header injections to be added to the response
type Transformation interface {
	Transform(c echo.Context) map[string]string
}
