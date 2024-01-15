package controllers

import (
	"net/http"

	"github.com/ponyjackal/go-gin-boilerplate/internal/domain/repositories"

	"github.com/gin-gonic/gin"
	"github.com/ponyjackal/go-gin-boilerplate/internal/domain/models"
)

func GetData(ctx *gin.Context) {
	var example []*models.Example
	repositories.Get(&example)
	ctx.JSON(http.StatusOK, &example)

}
