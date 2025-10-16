package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/quochao170402/notification-service/internal/core"
	"github.com/quochao170402/notification-service/internal/repository"
)

type NotificationHandler struct {
	repo repository.NotificationRepository
}

// NewNotificationHandler returns a handler with injected repository.
func NewNotificationHandler(r repository.NotificationRepository) *NotificationHandler {
	return &NotificationHandler{repo: r}
}

func RegisterTaskRoutes(rg *gin.RouterGroup, repo repository.NotificationRepository) {
	h := NewNotificationHandler(repo)

	rg.POST("", h.Create)
	rg.GET("", h.GetAll)
	rg.GET("/:id", h.GetByID)
	rg.PUT("/:id/status", h.UpdateStatus)
	rg.DELETE("/:id", h.Delete)
}

// Create adds a new notification.
func (h *NotificationHandler) Create(c *gin.Context) {
	var req core.Notification
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Default status if not provided
	if req.Status == "" {
		req.Status = core.StatusPending
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	id, err := h.repo.Create(ctx, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

// GetAll returns all notifications.
func (h *NotificationHandler) GetAll(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	list, err := h.repo.GetAll(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, list)
}

// GetByID retrieves a single notification.
func (h *NotificationHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	notif, err := h.repo.GetByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, notif)
}

// UpdateStatus sets the notification status (e.g., SENT, FAILED).
func (h *NotificationHandler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Status core.Status `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.repo.UpdateStatus(ctx, id, body.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"updated": id})
}

// Delete removes a notification.
func (h *NotificationHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(c.Request.Context(), 5*time.Second)
	defer cancel()

	if err := h.repo.Delete(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"deleted": id})
}
