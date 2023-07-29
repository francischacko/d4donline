package middleware

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gofiber/fiber/v2"
)

// RateLimiter is a middleware function that enforces rate limiting.
func RateLimiter(maxRequestsPerCPU int) fiber.Handler {
	fmt.Println("middleware got called")
	// Calculate the total rate limit based on the number of CPUs
	maxRequests := runtime.NumCPU() * maxRequestsPerCPU

	var (
		tokens     = maxRequests
		lastUpdate = time.Now()
	)

	return func(c *fiber.Ctx) error {
		// Calculate the time elapsed since the last token update
		elapsed := time.Since(lastUpdate)

		// Calculate how many tokens were added since the last update
		tokensToAdd := int(elapsed.Seconds()) * maxRequestsPerCPU

		// Update the token count with the new tokens
		tokens = min(tokens+tokensToAdd, maxRequests)

		// If no tokens are available, return a Too Many Requests response
		if tokens <= 0 {
			c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error": "Too many requests",
			})
			return nil
		}

		// Consume one token for the current request
		tokens--

		// Update the last token update time
		lastUpdate = time.Now()

		// Call the next handler in the chain
		return c.Next()
	}
}

// Helper function to return the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
