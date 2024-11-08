package middleware

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

func Logging(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		detail := logDetail{
			Timestamp: time.Now(),
			Latency:   time.Since(start).Milliseconds(),
			Path:      r.URL.Path,
			OS:        UnboxOSInfo(r.Context()),
		}

		if err := json.NewEncoder(os.Stdout).Encode(detail); err != nil {
			log.Printf("failed to encode log detail, err=%v\n", err)
		}
	}
	return http.HandlerFunc(fn)
}

type logDetail struct {
	Timestamp time.Time `json:"timestamp"`
	Latency   int64     `json:"latency"`
	Path      string    `json:"path"`
	OS        string    `json:"os"`
}
