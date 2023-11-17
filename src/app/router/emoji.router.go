package router

import (
	"github.com/gofiber/fiber/v2"
)

func (r *FiberRouter) GetEmoji(path string, h func(ctx *FiberCtx)) {
	r.emoji.Get(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) PostEmoji(path string, h func(ctx *FiberCtx)) {
	r.emoji.Post(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) DeleteEmoji(path string, h func(ctx *FiberCtx)) {
	r.emoji.Delete(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}
