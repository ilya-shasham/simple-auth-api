package handlers

import (
	"auth-api/checks"
	"auth-api/dbutils"
	"auth-api/globals"
	"auth-api/models"
	"auth-api/security"
	"errors"

	"github.com/gin-gonic/gin"
)

func UpdatePassword(ctx *gin.Context) {
	token := ctx.PostForm("token")
	new_password := ctx.PostForm("new_password")
	claims, err := security.CheckJwt(token)

	if CheckError(ctx, 0, err) {
		return
	}

	username := claims["username"].(string)
	member, exists := dbutils.RowExists[models.Member](
		"username = ?",
		func(m *models.Member) bool {
			return m.Username != ""
		},
		username,
	)

	if !exists {
		CheckError(ctx, 1, errors.New("username not found"))
		return
	}

	if code, err := checks.RunChecklist([]byte(new_password), passwordChecks); CheckError(ctx, code, err) {
		return
	}

	member.Password = security.Hash(new_password)

	globals.Db.Save(member)

	ctx.JSON(200, gin.H{
		"problem": nil,
		"result":  nil,
	})
}
