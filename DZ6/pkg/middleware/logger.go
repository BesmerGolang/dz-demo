package middleware

import (
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

// Logging принимает настроенный логгер и возвращает стандартную middleware.
func Logging(log *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			wrapper := &WrapperWriter{
				ResponseWriter: w,
				StatusCode:     http.StatusOK,
			}
			next.ServeHTTP(wrapper, r)
			duration := time.Since(start)

			entry := log.WithFields(logrus.Fields{
				"method":      r.Method,
				"path":        r.URL.Path,
				"status":      wrapper.StatusCode,
				"duration":    duration.String(),
				"remote_addr": r.RemoteAddr,
			})
			switch {
			case wrapper.StatusCode >= 500:
				entry.Error("HTTP request failed")
			case wrapper.StatusCode >= 400:
				entry.Warn("HTTP request client error")
			default:
				entry.Info("HTTP request")
			}
		})
	}
}
