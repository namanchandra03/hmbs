package utils

import "github.com/gin-gonic/gin"

func RespondErrWithJSON(ctx *gin.Context, statusCode int, errMsg string, err error) {
	ctx.AbortWithStatusJSON(statusCode, gin.H{
		"err": err,
		"msg": errMsg,
	})
}
