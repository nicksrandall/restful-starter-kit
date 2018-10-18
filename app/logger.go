package app

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

type correlationIdType int

const (
	requestIdKey correlationIdType = iota
	loggerIdKey
)

var logger *zap.Logger
var sugar *zap.SugaredLogger

func init() {
	logger, _ = zap.NewProduction()
}

// WithRqId returns a context which knows its request ID
func WithRqId(ctx context.Context, rqId string) context.Context {
	return context.WithValue(ctx, requestIdKey, rqId)
}

// Logger returns a zap logger with as much context as possible
func Logger(ctx context.Context) *zap.Logger {
	newLogger := logger
	if ctx != nil {
		if ctxRqId := middleware.GetReqID(ctx); ctxRqId != "" {
			newLogger = newLogger.With(zap.String("rqId", ctxRqId))
		}
		// add other things like sessionID here
	}
	return newLogger
}

type loggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *loggingResponseWriter {
	// WriteHeader(int) is not called if our response implicitly returns 200 OK, so
	// we default to that status code.
	return &loggingResponseWriter{w, http.StatusOK}
}

func (lrw *loggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

func LoggerMiddleware(logger *zap.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			// some evil middlewares modify this values
			path := r.URL.Path
			query := r.URL.RawQuery

			lrw := NewLoggingResponseWriter(w)
			next.ServeHTTP(lrw, r)

			end := time.Now()
			latency := end.Sub(start)

			logger.Info(path,
				zap.Int("status", lrw.statusCode),
				zap.String("method", r.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", r.RemoteAddr),
				zap.String("user-agent", r.UserAgent()),
				zap.String("time", end.Format(time.RFC3339)),
				zap.Duration("latency", latency),
			)
		}
		return http.HandlerFunc(fn)
	}
}
