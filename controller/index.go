package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/matejl/cron-manager/model"
)

func Index(c *gin.Context) {

	jobs, err := model.GetAllCronjobs()
	if err != nil {
		c.HTML(http.StatusBadRequest, "Cronjobs could not be fetched", gin.H{})
	}
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"jobs": jobs,
	})

}