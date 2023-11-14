package auth

var ExcludePath = map[string]struct{}{
	"POST /auth/verify":       {},
	"POST /auth/refreshToken": {},
	"POST /auth/googleUrl":    {},
	"GET /problem/":           {},
	"GET /course/":            {},
	// "GET /like/:userId":       {},
	"GET /emoji/":  {},
	"GET /rating/": {},
}
