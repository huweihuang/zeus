package server

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
	"github.com/sirupsen/logrus"
)

// Use pool to avoid concurrent access for rand.Source
var entropyPool = sync.Pool{
	New: func() interface{} {
		return rand.New(rand.NewSource(time.Now().UnixNano()))
	},
}

// Generate Unique ID
// Currently using ULID, this maybe conflict with other process with very low possibility
func genUniqueID() string {
	entropy := entropyPool.Get().(*rand.Rand)
	defer entropyPool.Put(entropy)
	id := ulid.MustNew(ulid.Now(), entropy)
	return id.String()
}

func requestIDMiddleware(c *gin.Context) {
	reqID := genUniqueID()
	c.Set("req_id", reqID)
	c.Header("X-Request-ID", reqID)
}

func logMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		fields := logrus.Fields{
			"path":    path,
			"query":   query,
			"latency": formatlatency(latency),
			"ip":      clientIP,
			"method":  method,
			"code":    statusCode,
			"req_id":  c.GetString("req_id"),
		}

		if statusCode >= 500 {
			logger.WithFields(fields).Error()
		} else if statusCode >= 400 {
			logger.WithFields(fields).Warn()
		} else {
			logger.WithFields(fields).Info()
		}
	}
}

func formatlatency(latency time.Duration) int {
	// Convert to milliseconds
	return int(latency.Seconds() * 1000)
}
