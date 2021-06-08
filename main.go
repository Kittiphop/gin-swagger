package main

import (
	"log"
	"net/http"

	_ "github.com/Kittiphop/gin-swagger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gin-swagger/user"
)

// @title Gin Swagger Example API
// @version 1.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func main() {
	// Gin instance
	r := gin.New()

	// Routes
	r.GET("/", Hello)

	r.GET("/users", user.GetList)
	r.PUT("/users", user.Create)

	url := ginSwagger.URL("http://localhost:3001/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// url := ginSwagger.URL("http://localhost:3001/swagger/doc.json") // The url pointing to API definition
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// Start server
	if err := r.Run(":3001"); err != nil {
		log.Fatal(err)
	}
}

// Hello godoc
// @Summary say hello!.
// @Description say hello!.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} ResponseHello
// @Router / [get]
func Hello(c *gin.Context) {
	res := map[string]interface{}{
		"data": "say hello!",
	}

	c.JSON(http.StatusOK, res)
}

type ResponseHello struct {
	SayHello  string
	Number    int
	Something interface{}
}
