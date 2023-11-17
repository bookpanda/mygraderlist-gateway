package router

import "github.com/gofiber/fiber/v2"

func (r *FiberRouter) GetRating(path string, h func(ctx *FiberCtx)) {
	r.rating.Get(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) PostRating(path string, h func(ctx *FiberCtx)) {
	r.rating.Post(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) PutRating(path string, h func(ctx *FiberCtx)) {
	r.rating.Put(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}

func (r *FiberRouter) DeleteRating(path string, h func(ctx *FiberCtx)) {
	r.rating.Delete(path, func(c *fiber.Ctx) error {
		h(NewFiberCtx(c))
		return nil
	})
}
