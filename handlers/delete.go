package handlers

import (
	"auth-api/dbutils"
	"auth-api/globals"
	"auth-api/models"
	"auth-api/security"
	"errors"

	"github.com/gin-gonic/gin"
)

func DeleteMember(ctx *gin.Context) {
	token := ctx.PostForm("token")
	claims, err := security.CheckJwt(token)

	if CheckError(ctx, 0, err) {
		return
	}

	username := claims["username"].(string)
	_, exists := dbutils.RowExists[models.Member](
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

	globals.Db.Exec("delete from members where username = ?", username)

	ctx.JSON(200, gin.H{
		"problem": nil,
		"result":  nil,
	})
}
