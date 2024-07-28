package endpoints

import (
	"log"
	"net/http"
)

const (
	ContentSecurityPolicyVal = "default-src 'self'; style-src 'self' fonts.googleapis.com; font-src fonts.gstatic.com"
	ReferrerPolicyVal        = "origin-when-cross-origin"
	XContentTypeOptionsVal   = "nosniff"
	XFrameOptionsVal         = "deny"
	XXSSProtectionVal        = "0"
)

var (
	safeHeaders = map[string]string{
		"Content-Security-Policy": ContentSecurityPolicyVal,
		"Referrer-Policy":         ReferrerPolicyVal,
		"X-Content-Type-Options":  XContentTypeOptionsVal,
		"X-Frame-Options":         XFrameOptionsVal,
		"X-XSS-Protection":        XXSSProtectionVal,
	}
)

func SecureHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for key, value := range safeHeaders {
			w.Header().Set(key, value)
		}
		next.ServeHTTP(w, r)
	})
}

func LogHandler(next http.Handler, log *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.RequestURI())
		next.ServeHTTP(w, r)
	})
}
