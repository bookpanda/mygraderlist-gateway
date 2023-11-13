package health_check

import (
	"net/http"

	"github.com/bookpanda/mygraderlist-gateway/src/app/router"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HealthCheck(c *router.GinCtx) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"Health": "大丈夫",
	})
	return
}
