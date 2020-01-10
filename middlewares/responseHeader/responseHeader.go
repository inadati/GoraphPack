package responseHeader

import (
	"net/http"
	"os"
)

var ALLOW_ORIGIN = os.Getenv("ACCESS_CONTROL_ALLOW_ORIGIN")

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", ALLOW_ORIGIN)
			w.Header().Set("Access-Control-Allow-Headers", "authorization,content-type")
			w.Header().Set("Access-Control-Allow-Methods", "OPTIONS,GET,POST")
			next.ServeHTTP(w, r)
		})
	}
}
