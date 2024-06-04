package item

import (
	"cakho.com/tudye/domain/product"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ItemHttpDelivery struct {
	api      *echo.Group
	itemRepo ItemRepository
}

func (i *ItemHttpDelivery) Get(c echo.Context) error {
	//generate code about get all item
	items, err := i.itemRepo.FindAll(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, items)
}

func (i *ItemHttpDelivery) Create(c echo.Context) error {
	//generate code about create item
	item := new(product.Item)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err)

	}

	err := i.itemRepo.Create(c.Request().Context(), *item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)

	}

	return c.JSON(http.StatusOK, item)

}

func (i *ItemHttpDelivery) Update(c echo.Context) error {
	item := new(product.Item)
	if err := c.Bind(item); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	err := i.itemRepo.Update(c.Request().Context(), *item)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, item)
}

func (i *ItemHttpDelivery) Delete(c echo.Context) error {
	//generate code about delete item
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id is required")
	}

	err := i.itemRepo.Delete(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "item deleted")

}

func (i *ItemHttpDelivery) GetById(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return c.JSON(http.StatusBadRequest, "id is required")
	}

	item, err := i.itemRepo.FindById(c.Request().Context(), id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, item)
}

func NewItemHttpDelivery(api *echo.Group, itemRepo ItemRepository) *ItemHttpDelivery {
	http := &ItemHttpDelivery{api: api, itemRepo: itemRepo}
	api.GET("/all", http.Get)
	api.POST("", http.Create)
	api.PUT("", http.Update)
	api.DELETE("/:id", http.Delete)
	api.GET("/:id", http.GetById)

	return http
}
