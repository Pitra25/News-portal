package rest

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (h *Router) GetAllNews(c echo.Context) error {
	values := c.QueryParams()
	params := getParams(values)

	list, err := h.manager.GetNewsByFilters(c.Request().Context(), params.NewFilter())
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	if len(list) == 0 {
		return newErrorResponse(c, http.StatusNoContent, errors.New("no list found"))
	}

	return c.JSON(http.StatusOK, NewNewsList(list))
}

func (h *Router) GetNewsById(c echo.Context) error {
	newsIdStr := c.Param("id")

	newsId, err := strconv.Atoi(newsIdStr)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	news, err := h.manager.GetNewsById(c.Request().Context(), newsId)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, NewNews(news))
}

func (h *Router) GetAllShortNews(c echo.Context) error {
	values := c.QueryParams()
	params := getParams(values)

	list, err := h.manager.GetNewsByFilters(c.Request().Context(), params.NewFilter())
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	if len(list) == 0 {
		return newErrorResponse(c, http.StatusNoContent, errors.New("no news found"))
	}

	return c.JSON(http.StatusOK, NewNewsSummaries(list))
}

func (h *Router) GetNewsCount(c echo.Context) error {
	values := c.QueryParams()
	params := getParams(values)

	count, err := h.manager.GetNewsCount(c.Request().Context(), params.NewFilter())
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, count)
}

func (h *Router) GetAllCategories(c echo.Context) error {
	categories, err := h.manager.GetAllCategory(c.Request().Context())
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, NewCategories(categories))
}

func (h *Router) GetAllTags(c echo.Context) error {
	tags, err := h.manager.GetAllTag(c.Request().Context())
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, NewTags(tags))
}
