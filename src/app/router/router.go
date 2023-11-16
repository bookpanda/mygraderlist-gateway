package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type FiberRouter struct {
	*fiber.App
	user    fiber.Router
	auth    fiber.Router
	course  fiber.Router
	problem fiber.Router
	like    fiber.Router
	emoji   fiber.Router
	rating  fiber.Router
}

type IGuard interface {
	Use(*FiberCtx)
}

func NewFiberRouter(authGuard IGuard) *FiberRouter {
	r := fiber.New(fiber.Config{
		StrictRouting: true,
		AppName:       "MyGraderList API",
	})

	r.Use(cors.New(cors.Config{AllowOrigins: "*"}))

	user := GroupWithAuthMiddleware(r, "/user", authGuard.Use)
	auth := GroupWithAuthMiddleware(r, "/auth", authGuard.Use)
	course := GroupWithAuthMiddleware(r, "/course", authGuard.Use)
	problem := GroupWithAuthMiddleware(r, "/problem", authGuard.Use)
	like := GroupWithAuthMiddleware(r, "/like", authGuard.Use)
	emoji := GroupWithAuthMiddleware(r, "/emoji", authGuard.Use)
	rating := GroupWithAuthMiddleware(r, "/rating", authGuard.Use)

	return &FiberRouter{r, user, auth, course, problem, like, emoji, rating}
}

func GroupWithAuthMiddleware(r *fiber.App, path string, middleware func(ctx *FiberCtx)) fiber.Router {
	return r.Group(path, func(c *fiber.Ctx) error {
		middleware(NewFiberCtx(c))
		return nil
	})
}
