package controllers

import (
	"net/http"

	"github.com/devhijazi/api/internal/models"
	"github.com/devhijazi/api/pkg/database"
	"github.com/gin-gonic/gin"
)

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")

	var user models.User

	database.GetDb().Where("id = ?", id).First(&user)

	ctx.JSON(http.StatusOK, user)
}

func CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	database.GetDb().Create(&user)

	ctx.JSON(http.StatusOK, &user)
}

func UpdateUser(ctx *gin.Context) {
}

func DeleteUser(ctx *gin.Context) {
}
