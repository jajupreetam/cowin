package middleware

import (
	slog "github.com/go-eden/slf4go"
	"net/http"
	"runtime/debug"
)

func LogRequestURI(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Debug(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				defer panic(err)
				slog.Error(err)
				debug.PrintStack()
				//RespondWithStatus(w, CreateApiError(err), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
