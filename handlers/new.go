package handlers

import (
	"auth-api/checks"
	"auth-api/globals"
	"auth-api/models"
	"auth-api/security"
	"errors"

	"github.com/gin-gonic/gin"
)

func NewUser(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	if code, err := checks.RunChecklist([]byte(username), usernameChecks); CheckError(ctx, code, err) {
		return
	}

	if code, err := checks.RunChecklist([]byte(password), passwordChecks); CheckError(ctx, code, err) {
		return
	}

	globals.Db.Save(&models.Member{
		Username: username,
		Password: security.Hash(password),
	})

	token, err := security.GenerateJwt(username)

	if err == nil {
		err = errors.New("")
	}

	ctx.JSON(200, gin.H{
		"problem": err.Error(),
		"result":  token,
	})
}
