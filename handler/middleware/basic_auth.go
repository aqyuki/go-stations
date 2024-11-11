package middleware

import "net/http"

func NewBasicAuthMiddleware(userID, password string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			parmUserID, parmPassword, ok := r.BasicAuth()
			if !ok || (parmUserID != userID || parmPassword != password) {
				w.Header().Set("WWW-Authenticate", `Basic realm="Authorization_Required"`)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}
