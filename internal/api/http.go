package api

import (
	"net/http"

	"github.com/Facundoblanco10/go-pulse-core/internal/jobs"
	"github.com/gin-gonic/gin"
)

func NewRouter(jobSvc *jobs.Service) *gin.Engine {
	router := gin.Default()

	// Health
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	// jobs
	jobHandler := NewJobHandler(jobSvc)
	jobHandler.RegisterRoutes(router)

	return router
}
