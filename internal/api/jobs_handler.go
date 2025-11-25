package api

import (
	"net/http"

	"github.com/Facundoblanco10/go-pulse-core/internal/jobs"
	"github.com/gin-gonic/gin"
)

type JobHandler struct {
	svc *jobs.Service
}

func NewJobHandler(svc *jobs.Service) *JobHandler {
	return &JobHandler{svc: svc}
}

func (h *JobHandler) RegisterRoutes(r *gin.Engine) {
	r.POST("/jobs", h.createJob)
}

func (h *JobHandler) createJob(c *gin.Context) {
	var input jobs.CreateJobInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid payload",
		})
		return
	}

	job, err := h.svc.CreateJob(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not create job",
		})
		return
	}

	c.JSON(http.StatusCreated, job)
}
