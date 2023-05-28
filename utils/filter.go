package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"hms/models"
	"strconv"
)

var defaultLimit = 10

func GetGenricFilters(ctx *gin.Context) *models.GenericFilters {
	limit := ctx.Query("limit ")
	page := ctx.Query("page")
	filters := &models.GenericFilters{}
	if limit != "" {
		limit, err := strconv.Atoi(limit)
		if err != nil {
			logrus.Errorf("not able to parse query param: %s", err)
			return filters
		}

		filters.Limit = limit
	} else {
		filters.Limit = defaultLimit
	}

	if page != "" {
		page, err := strconv.Atoi(limit)
		if err != nil {
			logrus.Errorf("not able to parse query param: %s", err)
			return filters
		}

		filters.Page = page
	}

	return filters
}

func GetComplainFilter(ctx *gin.Context) *models.ComplainFilters {
	filters := models.ComplainFilters{}
	filters.GenericFilters = GetGenricFilters(ctx)
	filters.ComplainID = ctx.Query("complainID")
	filters.HostelID = ctx.Query("hostelID")
	filters.StudentID = ctx.Query("studentID")

	return &filters
}
