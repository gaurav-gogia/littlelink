package biz

import "net/http"

// SetHeaders function sets security headers for APIs
// For more info please visit: https://security.stackexchange.com/questions/147554/security-headers-for-a-web-api
// Mozilla Security Cheat Sheet: https://infosec.mozilla.org/guidelines/web_security#web-security-cheat-sheet
func SetHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	w.Header().Set("Content-Security-Policy", "script-src 'self'")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "Deny")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
}
