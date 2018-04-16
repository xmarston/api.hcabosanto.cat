package main

import (
	"net/http"
	"fmt"
	"strings"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Forwarded-For") == AllowedXForwardedFor {
		// Create return string
		var request []string

		for name, headers := range r.Header {
			name = strings.ToLower(name)
			for _, h := range headers {
				request = append(request, fmt.Sprintf("%v: %v", name, h))
			}
		}

		w.Write([]byte(strings.Join(request, "\n")))
	} else {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Forbidden"))
	}
}
