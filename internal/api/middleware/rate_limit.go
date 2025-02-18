package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type rateLimiter struct {
	sync.Mutex
	requests map[string][]time.Time
}

var limiter = &rateLimiter{
	requests: make(map[string][]time.Time),
}

func RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		limiter.Lock()
		defer limiter.Unlock()

		now := time.Now()

		// Clean old requests
		if requests, exists := limiter.requests[ip]; exists {
			var valid []time.Time
			for _, t := range requests {
				if now.Sub(t) < time.Minute {
					valid = append(valid, t)
				}
			}
			limiter.requests[ip] = valid
		}

		// Check rate limit
		if len(limiter.requests[ip]) >= 100 { // 100 requests per minute
			c.JSON(429, gin.H{"error": "Rate limit exceeded"})
			c.Abort()
			return
		}

		// Add current request
		limiter.requests[ip] = append(limiter.requests[ip], now)

		c.Next()
	}
}
