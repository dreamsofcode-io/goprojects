package middleware

import "net/http"

// Middleware represents the type signature of a middleware
// function.
type Middleware func(http.Handler) http.Handler
