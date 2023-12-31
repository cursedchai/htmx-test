package api

import (
	"fmt"
	"htmx-test/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (handler Handler) GetItem(c echo.Context) error {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
	}
	var i db.Item
	c.Bind(&i)

	i, err = db.GetItem(handler.DB, int(id), i.Type)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}
	return c.Render(http.StatusOK, "item", i)
}

func (handler Handler) PutItem(c echo.Context) error {
	var i db.Item
	err := c.Bind(&i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}

	recipe, err := db.UpdateItem(handler.DB, i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}
	return c.Render(http.StatusOK, "item", recipe)
}

func (handler Handler) EditItem(c echo.Context) error {
	id_str := c.Param("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
	}

	var i db.Item
	err = c.Bind(&i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("id param (%s) not a number", id_str))
	}
	i.Id = int(id)

	return c.Render(http.StatusOK, "edit_item", i)
}

func (handler Handler) DeleteItem(c echo.Context) error {
	var i db.Item
	err := c.Bind(&i)
	fmt.Printf("%s\n", i.Type)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}

	err = db.DeleteItem(handler.DB, i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("%b", err))
	}
	return c.NoContent(http.StatusOK)
}
