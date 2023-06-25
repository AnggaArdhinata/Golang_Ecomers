package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mrwaggel/golimiter"
)

const limitPerMinute = 100
var limiter = golimiter.New(limitPerMinute, time.Minute)
func RateLimit(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		IP := c.RealIP()
		// Check if the IP is limited
		if limiter.IsLimited(IP) {
			return c.String(http.StatusTooManyRequests, fmt.Sprintf("to many request from %s", IP))
		}
		// Increment the value for the IP
		limiter.Increment(c.RealIP())
		// Continue default operation of Echo
		return next(c)
	}
}
