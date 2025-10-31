package rest

import (
	"News-portal/internal/newsportal"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

/*** News ***/

func (h *Router) GetAllNews(c echo.Context) error {
	var params queryParams
	if err := c.Bind(&params); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err)
	}

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

	news, err := h.manager.GetNewsByID(c.Request().Context(), newsId)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, NewNews(news))
}

func (h *Router) GetNewsSummaries(c echo.Context) error {
	var params queryParams
	if err := c.Bind(&params); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err)
	}

	list, err := h.manager.GetNewsByFilters(c.Request().Context(), params.NewFilter())
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	if len(list) == 0 {
		return newErrorResponse(c, http.StatusNoContent, errors.New("no news found"))
	}

	return c.JSON(http.StatusOK, NewNewsSummaries(list))
	//return nil
}

func (h *Router) GetNewsCount(c echo.Context) error {
	var params queryParams
	if err := c.Bind(&params); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err)
	}

	count, err := h.manager.GetNewsCount(c.Request().Context(), params.NewFilter())
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, count)
}

func (h *Router) AddNews(c echo.Context) error {
	var newItem *NewsInput
	if err := c.Bind(&newItem); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err)
	}

	res, err := h.manager.AddNews(c.Request().Context(), newsToManager(newItem))
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *Router) UpdateNews(c echo.Context) error {
	var input *NewsInput
	if err := c.Bind(&input); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err)
	}

	res, err := h.manager.UpdateNews(c.Request().Context(), newsToManager(input))
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return newNoContentResponse(c, http.StatusCreated, res)
}

func (h *Router) DeleteNews(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	res, err := h.manager.DeleteNews(c.Request().Context(), id)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return newNoContentResponse(c, http.StatusNoContent, res)
}

/*** Category ***/

func (h *Router) GetAllCategories(c echo.Context) error {
	categories, err := h.manager.GetAllCategory(c.Request().Context())
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, NewCategories(categories))
}

func (h *Router) AddCategory(c echo.Context) error {
	var newItem *CategoryInput
	if err := c.Bind(&newItem); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err)
	}

	res, err := h.manager.AddCategory(c.Request().Context(), categoryToManager(newItem))
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *Router) UpdateCategory(c echo.Context) error {
	var input *CategoryInput
	if err := c.Bind(&input); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err)
	}

	res, err := h.manager.UpdateCategory(c.Request().Context(), categoryToManager(input))
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}
	return newNoContentResponse(c, http.StatusCreated, res)
}

func (h *Router) DeleteCategory(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	res, err := h.manager.DeleteNews(c.Request().Context(), id)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return newNoContentResponse(c, http.StatusNoContent, res)
}

/*** Tag ***/

func (h *Router) GetAllTags(c echo.Context) error {
	tags, err := h.manager.GetTagsByFilters(c.Request().Context(), newsportal.Filters{})
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, NewTags(tags))
}

func (h *Router) AddTag(c echo.Context) error {
	var newItem *TagInput
	if err := c.Bind(&newItem); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err)
	}

	res, err := h.manager.AddTag(c.Request().Context(), tagToManager(newItem))
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *Router) UpdateTag(c echo.Context) error {
	var input *TagInput
	if err := c.Bind(&input); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, err)
	}

	res, err := h.manager.UpdateTag(c.Request().Context(), tagToManager(input))
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return newNoContentResponse(c, http.StatusCreated, res)
}

func (h *Router) DeleteTag(c echo.Context) error {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	res, err := h.manager.DeleteTag(c.Request().Context(), id)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, err)
	}

	return newNoContentResponse(c, http.StatusNoContent, res)
}
