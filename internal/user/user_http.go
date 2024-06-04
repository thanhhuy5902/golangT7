package user

import (
	"cakho.com/tudye/domain/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHttpDelivery struct {
	api      *echo.Group
	userRepo UserRepository
}

func (u *UserHttpDelivery) Get(c echo.Context) error {
	//generate code about get all user
	users, err := u.userRepo.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, users)
}

func (u *UserHttpDelivery) Create(c echo.Context) error {
	//generate code about create user
	user := new(user.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}

	err := u.userRepo.Create(c.Request().Context(), *user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHttpDelivery) Update(c echo.Context) error {
	user := new(user.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := u.userRepo.Update(c.Request().Context(), *user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHttpDelivery) Delete(c echo.Context) error {
	//generate code about delete user
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id is required")
	}
	err := u.userRepo.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, "deleted")
}

func (u *UserHttpDelivery) FindById(c echo.Context) error {
	//generate code about find user by id
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id is required")
	}

	user, err := u.userRepo.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, user)
}

func (u *UserHttpDelivery) FindAll(c echo.Context) error {
	//generate code about find all user
	users, err := u.userRepo.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, users)
}

func NewUserHttpDelivery(api *echo.Group, userRepo UserRepository) *UserHttpDelivery {
	http := &UserHttpDelivery{
		api:      api,
		userRepo: userRepo,
	}
	api.GET("", http.Get)
	api.POST("", http.Create)
	api.PUT("", http.Update)
	api.DELETE("/:id", http.Delete)
	api.GET("/:id", http.FindById)
	api.GET("", http.FindAll)
	return http
}
