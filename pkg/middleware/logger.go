package middleware

import (
	"log"
	"net/http"
	"time"
)

func LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		// Log each Request and where it came from
		log.Printf("Request Received: %s\t%s\t%s",
			r.Method,
			r.RequestURI,
			r.RemoteAddr,
		)
		//************************
		// Wrap Handler
		//************************
		next.ServeHTTP(w, r)
		// Log runtime of Request
		log.Printf("Request Completed In: \t%s",
			time.Since(start),
		)
	})
}
