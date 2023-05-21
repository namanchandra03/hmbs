package handlers

import (
	"github.com/gin-gonic/gin"
	"hms/models"
	"hms/utils"
	"net/http"
)

func AddComplain(ctx *gin.Context) {
	complainInfo := models.ComplainDetails{}
	err := ctx.ShouldBindJSON(&complainInfo)
	if err != nil {
		utils.RespondErrWithJSON(ctx, http.StatusBadRequest, "unable to parse request body", err)
		return
	}
}
