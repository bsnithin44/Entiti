package main

import (
	"net/http"

	"github.com/bsnithin44/entiti/cmd/router"
	"github.com/labstack/echo/v4"
	v1 "github.com/bsnithin44/entiti/pkg/api/v1"
)

func init() {

	// Load DB instance using a singleton
	// database.GetDbSession()

}

func main() {

	r := router.New()

	r.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service is Up! ")
	})

	groupV1 := r.Group("/entities/api/v1")

	// entities
	groupV1.POST("/entity-type", v1.CreateEntityTypeHandler)
	groupV1.GET("/entity-type", v1.GetEntityTypesHandler)
	groupV1.GET("/entity-type/:entityTypeId", v1.GetEntityTypeHandler)
	groupV1.PUT("/entity-type/:entityTypeId", v1.UpdateEntityTypeHandler)

	r.Logger.Fatal(r.Start("0.0.0.0:" + "8000"))
}
