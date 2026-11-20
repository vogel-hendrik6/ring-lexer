package middleware

import (
	"net/http"
	"time"
)

func Timeout(d time.Duration, next http.Handler) http.Handler {
	return http.TimeoutHandler(next, d, "request timeout")
}
