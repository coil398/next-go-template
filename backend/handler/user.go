package handler

import (
	"backend/model"
	"backend/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
)

func registerV1UserHandlers(v1 *echo.Group, s *Server) {
	v1.GET("/user/:id", s.v1GetUser)
	v1.POST("/user", s.v1PostUser)

}

type v1GetUserRequest struct {
	ID int `param:"id"`
}

func (s *Server) v1GetUser(c echo.Context) error {
	var req v1GetUserRequest
	if err := c.Bind(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	user, err := usecase.V1GetUser(s.Store.User, req.ID)

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	if user == nil {
		return c.String(http.StatusBadGateway, "No user found")
	}

	return c.JSON(http.StatusOK, user)
}

type v1PostUserRequest struct {
	Name string `json:"name"`
}

func (s *Server) v1PostUser(c echo.Context) error {
	var req v1PostUserRequest
	if err := c.Bind(&req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
	}

	_, err := usecase.V1PostUser(s.Store.User, &model.User{Name: req.Name})

	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, req)
}
