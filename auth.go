package main

import (
	"encoding/base64"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		// Check that the header begins with a prefix of Bearer
		if !strings.HasPrefix(authorization, "Bearer ") {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Pull out the token
		encodedToken := strings.TrimPrefix(authorization, "Bearer ")

		// Decode the token from base 64
		token, err := base64.StdEncoding.DecodeString(encodedToken)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		w.Header().Add("User", string(token))

		// We're just assuming a valid base64 token is a valid user id.
		//userID := string(token)
		//fmt.Println("userID:", userID)

		next.ServeHTTP(w, r)
	})
}
