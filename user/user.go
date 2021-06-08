package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetList godoc
// @Summary list!.
// @Description list!.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} []ResponseUser
// @Router /users [get]
func GetList(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}

// Create godoc
// @Summary create!.
// @Description create!.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} []ResponseUser
// @Router /users [put]
func Create(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}

type ResponseUser struct {
	ID        uint
	Name      string
	Something interface{}
}
