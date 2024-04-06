package main

import (
	"log"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(wrapped, r)

		addInfo := "["
		if len(wrapped.Header().Get("User")) > 0 {
			addInfo += "user:" + wrapped.Header().Get("User") + ", "
		}

		log.Println(wrapped.statusCode, r.Method, r.URL.Path, addInfo[:len(addInfo)-2]+"]", time.Since(start))
	})
}
