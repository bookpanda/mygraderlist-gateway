package health_check

import (
	"net/http"
)

type Handler struct {
}

type IContext interface {
	JSON(statusCode int, v interface{})
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthCheck(c IContext) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Health": "大丈夫",
	})
	return
}
