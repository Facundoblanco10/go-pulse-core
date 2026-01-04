package api

import (
	"errors"
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
	r.PATCH("/jobs/:id/cancel", h.cancelJob)
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

func (h *JobHandler) cancelJob(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing job ID",
		})
		return
	}

	err := h.svc.CancelJob(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, jobs.ErrJobNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "job not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "could not cancel job",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "job canceled successfully",
	})
}
