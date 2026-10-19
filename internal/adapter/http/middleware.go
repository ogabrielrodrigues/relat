package http

import (
	"fmt"
	"net/http"
	"slices"
	"time"

	"relat/internal/platform/httperr"
	"relat/internal/platform/logger"
)

type middleware func(next http.Handler) http.Handler

func chain(handler http.Handler, middlewares ...middleware) http.Handler {
	for _, midd := range slices.Backward(middlewares) {
		handler = midd(handler)
	}

	return handler
}

func logMiddleware(next http.Handler) http.Handler {
	log, err := logger.NewLogger()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err != nil {
			SendError(w, httperr.ErrInternalServerError())
			return
		}

		t := time.Now()

		next.ServeHTTP(w, r)

		log.InfoContext(
			r.Context(),
			"new request",
			"route",
			fmt.Sprintf("%s %s", r.Method, r.URL.Path),
			"addr",
			r.RemoteAddr,
			"response",
			time.Since(t),
		)
	})
}
