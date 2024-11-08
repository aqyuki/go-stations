package middleware

import (
	"log"
	"net/http"
)

func Recovery(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				// NOTE: Slackなどにアラートを上げるなどの処理を必要に応じて追加する
				// NOTE: trace idなどをログに付与すると原因解明などに役立つ(別のmiddlewareで実装するかも)
				log.Printf("panic occurred: %v", err)
			}
		}()
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
