package router

import (
	"github.com/gofiber/fiber/v2"
)

func (r *FiberRouter) GetCourse(path string, h func(ctx *FiberCtx)) {
	r.course.Get(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}
