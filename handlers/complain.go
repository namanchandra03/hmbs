package handlers

import (
	"github.com/gin-gonic/gin"
	"hms/database/dbHelper"
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

	complainID, err := dbHelper.AddComplain(ctx.Request.Context(), complainInfo)
	if err != nil {
		utils.RespondErrWithJSON(ctx, http.StatusInternalServerError, "not able to register the complain", err)
		return
	}

	utils.RespondSuccessWithJSON(ctx, http.StatusCreated, complainID)
}

func EditComplain(ctx *gin.Context) {
	complainID := ctx.Param("complainID")
	updatedComplain := models.ComplainDetails{}
	err := ctx.ShouldBindJSON(&updatedComplain)
	if err != nil {
		utils.RespondErrWithJSON(ctx, http.StatusBadRequest, "unable to parse request body", err)
		return
	}

	err = dbHelper.UpdateComplain(ctx.Request.Context(), updatedComplain, complainID)
	if err != nil {
		utils.RespondErrWithJSON(ctx, http.StatusInternalServerError, "not able  to edit the complain", err)
		return
	}

	utils.RespondSuccessWithJSON(ctx, http.StatusOK, "successfully updated the complain")
}

func DeleteComplain(ctx *gin.Context) {
	complainID := ctx.Param("complainID")
	err := dbHelper.DeleteComplain(ctx.Request.Context(), complainID)
	if err != nil {
		utils.RespondErrWithJSON(ctx, http.StatusInternalServerError, "not able  to delete the complain", err)
		return
	}

	utils.RespondSuccessWithJSON(ctx, http.StatusOK, "successfully deleted the complain")
}

func GetComplain(ctx *gin.Context) {
	filters := utils.GetComplainFilter(ctx)
	complains, err := dbHelper.GetComplains(ctx.Request.Context(), filters)
	if err != nil {
		utils.RespondErrWithJSON(ctx, http.StatusInternalServerError, "no able to fetch complains", err)
		return
	}

	utils.RespondSuccessWithJSON(ctx, http.StatusOK, complains)
}
