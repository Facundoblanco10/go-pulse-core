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
	r.GET("/jobs", h.listJobs)
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

func (h *JobHandler) listJobs(c *gin.Context) {
	jobs, err := h.svc.ListJobs(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to retrieve jobs",
		})
		return
	}

	c.JSON(http.StatusOK, jobs)
}
