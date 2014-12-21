package main

import (
	"strings"

	"github.com/gin-gonic/gin"
)

type Err struct {
	Status  int    `json:"status"`
	Message string `json:"error"`
}

func ErrorResponse(code int, message string, c *gin.Context) {
	c.JSON(code, Err{code, message})
}

func SuccessResponse(message string, c *gin.Context) {
	c.JSON(200, map[string]string{"message": message})
}

func GetFrontends(c *gin.Context) {
	frontends, err := hipache.Frontends()

	if err != nil {
		ErrorResponse(400, err.Error(), c)
		return
	}

	c.JSON(200, frontends)
}

func GetBackends(c *gin.Context) {
	fe := c.Params.ByName("fe")
	exists, err := hipache.FrontendExists(fe)

	if err != nil {
		ErrorResponse(400, err.Error(), c)
		return
	}

	if !exists {
		ErrorResponse(404, "Frontend does not exist", c)
		return
	}

	backends, err := hipache.Backends(c.Params.ByName("fe"))

	if err != nil {
		ErrorResponse(400, err.Error(), c)
		return
	}

	c.JSON(200, backends)
}

func DeleteFrontend(c *gin.Context) {
	fe := c.Params.ByName("fe")
	exists, err := hipache.FrontendExists(fe)

	if err != nil {
		ErrorResponse(400, err.Error(), c)
		return
	}

	if !exists {
		ErrorResponse(404, "Frontend does not exist", c)
		return
	}

	err = hipache.RemoveFrontend(fe)

	if err != nil {
		ErrorResponse(400, err.Error(), c)
		return
	}

	SuccessResponse("Frontend has been deleted", c)
}

func FlushFrontends(c *gin.Context) {
	err := hipache.Flush()

	if err != nil {
		ErrorResponse(400, err.Error(), c)
		return
	}

	c.JSON(200, map[string]string{"message": "All frontends are now flushed"})
}

func CreateFrontend(c *gin.Context) {
	host := c.Request.FormValue("host")
	backends := c.Request.FormValue("backends")

	if host == "" {
		ErrorResponse(400, "Missing parameter 'host'", c)
		return
	}

	exists, _ := hipache.FrontendExists(host)

	if exists {
		ErrorResponse(400, "Frontend already exists", c)
		return
	}

	err := hipache.AddFrontend(host)

	if err != nil {
		ErrorResponse(400, err.Error(), c)
		return
	}

	if backends != "" {
		for _, be := range strings.Split(backends, ",") {
			// TODO: Some error checking!
			hipache.AddBackend(host, be)
		}
	}

	SuccessResponse("Frontend has been created", c)
}

func CreateBackend(c *gin.Context) {
	fe := c.Params.ByName("fe")
	exists, _ := hipache.FrontendExists(fe)

	if !exists {
		ErrorResponse(404, "Frontend does not exist", c)
		return
	}

	backends := c.Request.FormValue("backends")

	if backends == "" {
		ErrorResponse(400, "Missing 'backends' parameter", c)
		return
	}

	for _, be := range strings.Split(backends, ",") {
		hipache.AddBackend(fe, be)
	}

	SuccessResponse("Backends created", c)
}

func DeleteBackend(c *gin.Context) {
	fe := c.Params.ByName("fe")
	be := c.Request.FormValue("backend")

	if be == "" {
		ErrorResponse(400, "Missing 'backend' parameter", c)
		return
	}

	exists, _ := hipache.FrontendExists(fe)
	if !exists {
		ErrorResponse(400, "Frontend does not exist", c)
		return
	}

	err := hipache.RemoveBackend(fe, be)

	if err != nil {
		ErrorResponse(400, err.Error(), c)
		return
	}

	SuccessResponse("Backend removed", c)
}
