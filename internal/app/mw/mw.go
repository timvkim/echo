package mw

import (
	"log"
	"strings"

	"github.com/labstack/echo/v4"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		// get the value of the header
		val := ctx.Request().Header.Get("User-Role")

		// check if received value is equal to "admin"
		if strings.EqualFold(val, roleAdmin) {
			log.Println("red button user detected")
		}

		// call the next handler
		err := next(ctx)
		if err != nil {
			return err
		}

		return nil
	}
}
