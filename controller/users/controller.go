package controller

import (
	"bni-test/common"
	"bni-test/model"
	"net/http"

	"github.com/labstack/echo"
)

var (
	memory = map[string]model.Users{}
)

func SaveUser(c echo.Context) error {
	u := new(model.InsertUserRequest)

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := model.Users{
		Fullname:      u.Fullname,
		Address:       u.Address,
		BirthdayPlace: u.BirthdayPlace,
		Birthday:      u.Birthday,
		KTPnumber:     u.KTPnumber,
		PhoneNumber:   u.PhoneNumber,
	}

	memory[user.KTPnumber] = user

	return c.JSON(common.NewSuccessResponseWithoutData())
}

func GetUser(c echo.Context) error {
	ktp := c.Param("ktp")

	_, exist := memory[ktp]

	if !exist {
		return c.JSON(common.NewNotFoundResponse())
	}

	return c.JSON(common.NewSuccessResponse(memory[ktp]))
}

func GetAllUser(c echo.Context) error {

	users := []model.Users{}

	for _, item := range memory {
		users = append(users, item)
	}

	return c.JSON(common.NewSuccessResponse(users))
}

func UpdateUser(c echo.Context) error {
	ktp := c.Param("ktp")

	_, exist := memory[ktp]

	if !exist {
		return c.JSON(common.NewNotFoundResponse())
	}

	u := new(model.UpdatetUserRequest)

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	user := model.Users{
		Fullname:      u.Fullname,
		Address:       u.Address,
		BirthdayPlace: u.BirthdayPlace,
		Birthday:      u.Birthday,
		PhoneNumber:   u.PhoneNumber,
	}

	memory[ktp] = user

	return c.JSON(common.NewSuccessResponse(memory[ktp]))
	// return c.String(http.StatusOK, id)
}

func DeleteUser(c echo.Context) error {
	ktp := c.Param("ktp")

	_, exist := memory[ktp]

	if !exist {
		return c.JSON(common.NewNotFoundResponse())
	}

	history := memory[ktp]

	for key := range memory {
		if key == ktp {
			delete(memory, key)
		}
	}
	return c.JSON(common.NewSuccessDeleteResponse(history))
}
