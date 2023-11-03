package handlers

import "github.com/gin-gonic/gin"

func CheckError(ctx *gin.Context, code int, err error) bool {
	if err != nil {
		ctx.JSON(400, gin.H{
			"problem": map[string]any{
				"code": code,
				"desc": err.Error(),
			},
			"result": nil,
		})
		return true
	}
	return false
}
