package middleware

import (
	"log"
	"net/http"
	"time"
)

func Log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("%s %s %s", time.Now(), req.Method, req.RequestURI)
		next.ServeHTTP(w, req)
	})
}
