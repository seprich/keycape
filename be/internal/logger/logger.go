package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/seprich/keycape/internal/util"
	"log/slog"
	"net/http"
	"os"
	"time"
)

var Logger *slog.Logger = loadLogger()

// Logging middlewares for gin

var GinLogger gin.HandlerFunc = generateGinLogger(Logger)
var GinRecovery = generateGinRecovery(Logger)

func loadLogger() *slog.Logger {
	return slog.New(slog.NewJSONHandler(os.Stderr, nil))
}

func generateGinLogger(l *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		attributes := collectRequestLogItems(c.Request)
		durationStr := fmt.Sprintf("%d μs", time.Since(startTime).Microseconds())
		attributes = append(attributes, slog.String("latency", durationStr))
		attributes = append(attributes, slog.Int("status", c.Writer.Status()))
		attributes = append(attributes, slog.String("client_ip", c.ClientIP()))
		attributes = append(attributes, collectErrorData(c))
		l.LogAttrs(c.Request.Context(), slog.LevelInfo, "request completed", attributes...)
	}
}

func generateGinRecovery(l *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				attributes := collectRequestLogItems(c.Request)
				attributes = append(attributes, slog.Any("error", err))
				attributes = append(attributes, slog.String("client_ip", c.ClientIP()))
				attributes = append(attributes, slog.String("stack_trace", util.CaptureStackTrace()))
				l.LogAttrs(c.Request.Context(), slog.LevelError, "panic recovered", attributes...)

				code := http.StatusInternalServerError
				c.AbortWithStatusJSON(code, gin.H{"error": http.StatusText(code)})
			}
		}()
		c.Next()
	}
}

var sensitiveHeaders = []string{"Authorization", "Cookie", "Client-Certificate", "Tls-Client-Certificate"}

func collectRequestLogItems(r *http.Request) []slog.Attr {
	var result []slog.Attr
	result = append(result, slog.String("path", r.URL.Path))
	result = append(result, slog.String("method", r.Method))
	result = append(result, slog.String("protocol", r.Proto))
	var headers []any
	for key, values := range r.Header {
		if lo.Contains(sensitiveHeaders, key) {
			headers = append(headers, slog.String(key, "<redacted>"))
		} else {
			for _, value := range values {
				headers = append(headers, slog.String(key, value))
			}
		}
	}
	result = append(result, slog.Group("headers", headers...))
	return result
}

func collectErrorData(c *gin.Context) slog.Attr {
	var result []any
	if r, e := c.Get("error_code"); e {
		result = append(result, slog.Any("code", r))
	}
	if r, e := c.Get("error_details"); e {
		result = append(result, slog.Any("details", r))
	}
	if r, e := c.Get("error_ref"); e {
		result = append(result, slog.Any("ref", r))
	}
	return slog.Group("error", result...)
}
