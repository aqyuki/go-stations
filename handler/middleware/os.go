package middleware

import (
	"context"
	"net/http"

	"github.com/mileusna/useragent"
)

type contextKey string

const osKey = contextKey("os")

func BoxOSInfo(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		ua := useragent.Parse(r.UserAgent())
		ctx := context.WithValue(r.Context(), osKey, ua.OS)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	}
	return http.HandlerFunc(fn)
}

func UnboxOSInfo(ctx context.Context) string {
	os, ok := ctx.Value(osKey).(string)
	if !ok {
		return ""
	}
	return os
}
