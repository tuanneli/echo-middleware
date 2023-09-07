package mv

import (
	"github.com/labstack/echo/v4"
	"log"
	"strings"
)

const roleAdmin = "admin"

func RoleCheck(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		role := c.Request().Header.Get("User-Role")

		if strings.EqualFold(role, roleAdmin) {
			log.Println("Red button user detected!")
		}

		err := next(c)

		if err != nil {
			return err
		}

		return nil
	}
}
