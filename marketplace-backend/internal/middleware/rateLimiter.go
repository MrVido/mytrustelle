package middleware

import (
    "github.com/gin-gonic/gin"
    "golang.org/x/time/rate"
    "net/http"
    "sync"
    "time"
)

// RateLimiter stores the limiter for each client.
type RateLimiter struct {
    visitors map[string]*rate.Limiter
    mtx      sync.Mutex
    r        rate.Limit
    b        int
}

// NewRateLimiter creates a new RateLimiter.
func NewRateLimiter(r rate.Limit, b int) *RateLimiter {
    return &RateLimiter{
        visitors: make(map[string]*rate.Limiter),
        r:        r,
        b:        b,
    }
}

// AddVisitor checks the map for an existing rate limiter and adds a new one if none exists.
func (l *RateLimiter) AddVisitor(ip string) *rate.Limiter {
    l.mtx.Lock()
    defer l.mtx.Unlock()

    limiter, exists := l.visitors[ip]
    if !exists {
        limiter = rate.NewLimiter(l.r, l.b)
        l.visitors[ip] = limiter
    }
    return limiter
}

// GetVisitor returns the rate limiter for the current visitor's IP address.
func (l *RateLimiter) GetVisitor(ip string) *rate.Limiter {
    l.mtx.Lock()
    limiter, exists := l.visitors[ip]
    l.mtx.Unlock()

    if !exists {
        return l.AddVisitor(ip)
    }
    return limiter
}

// Middleware is the gin middleware for rate limiting.
func (l *RateLimiter) Middleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        limiter := l.GetVisitor(c.ClientIP())

        if !limiter.Allow() {
            c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
            return
        }

        c.Next()
    }
}
