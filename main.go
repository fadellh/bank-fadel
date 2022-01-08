package main

import (
	usersController "bni-test/controller/users"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", GetHello)
	e.POST("/users", usersController.SaveUser)
	e.GET("/users/:ktp", usersController.GetUser)
	e.GET("/users", usersController.GetAllUser)
	e.PUT("/users/:ktp", usersController.UpdateUser)
	e.DELETE("/users/:ktp", usersController.DeleteUser)

	e.Logger.Fatal(e.Start(":1323"))

}

func GetHello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
