package v1

import (
	"net/http"

	api "github.com/bsnithin44/entiti/pkg/api"
	entities "github.com/bsnithin44/entiti/pkg/domain/entities"

	"github.com/labstack/echo/v4"
)

func CreateEntityTypeHandler(c echo.Context) error {

	Response := &api.CreatedResponse{}
	Response.Request.Uri = c.Request().URL.Path
	Response.Request.QueryString = c.QueryString()

	Request := new(entities.CreateEntityTypeStruct)
	if err := c.Bind(Request); err != nil {
		response := &api.Response{}
		response.Description = "Binding failed"
		return c.JSON(http.StatusBadRequest, response)
	}
	Response.Request.Body = Request
	service, err := entities.CreateEntityType(*Request)
	if err != nil {
		Response.Error.AdditionalInfo = err.AdditionalInfo
		Response.Error.Description = err.Description
		return c.JSONPretty(err.StatusCode, Response, "    ")
	} else {
		Response.Data.Id = service.Id
		return c.JSONPretty(http.StatusCreated, Response, "    ")
	}

}
func GetEntityTypeHandler(c echo.Context) error {
	Response := &api.Response{}
	Response.Uri = c.Request().URL.Path
	Response.QueryString = c.QueryString()

	entityTypeId := c.Param("entityTypeId")
	service, err := entities.GetEntityType(entityTypeId)
	if err != nil {
		Response.Error.AdditionalInfo = err.AdditionalInfo
		Response.Error.Description = err.Description
		return c.JSONPretty(err.StatusCode, Response, "    ")
	} else {
		Response.Data = service

		return c.JSONPretty(http.StatusOK, Response, "    ")
	}
}
func GetEntityTypesHandler(c echo.Context) error {
	Response := &api.Response{}
	Response.Uri = c.Request().URL.Path
	Response.QueryString = c.QueryString()

	isActive := c.QueryParam("isActive")
	services, err := entities.GetEntityTypes(isActive)
	if err != nil {
		Response.Error.AdditionalInfo = err.AdditionalInfo
		Response.Error.Description = err.Description
		return c.JSONPretty(err.StatusCode, Response, "    ")
	} else {
		Response.Data = services
		return c.JSONPretty(http.StatusOK, Response, "    ")
	}
}

func UpdateEntityTypeHandler(c echo.Context) error {
	Response := &api.CreatedResponse{}
	Response.Uri = c.Request().URL.Path
	Response.QueryString = c.QueryString()

	Request := new(entities.EntityTypeBase)
	entityTypeId := c.Param("entityTypeId")

	if err := c.Bind(Request); err != nil {
		response := &api.Response{}
		response.Description = "Binding failed"
		return c.JSON(http.StatusBadRequest, response)
	}

	service, err := entities.UpdateEntityType(entityTypeId, *Request)
	if err != nil {
		Response.Error.AdditionalInfo = err.AdditionalInfo
		Response.Error.Description = err.Description
		return c.JSONPretty(err.StatusCode, Response, "    ")
	} else {
		Response.Data.Id = service.Id
		return c.JSONPretty(http.StatusOK, Response, "    ")
	}
}

