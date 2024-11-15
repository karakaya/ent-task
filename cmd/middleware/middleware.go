package middleware

import (
	"entain-golang-task/pkg/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"net/http"
)

type Middleware func(httprouter.Handle) httprouter.Handle

var validSourceTypes = map[string]bool{
	"game":    true,
	"server":  true,
	"payment": true,
}

//ignored content-type json check

func SourceTypeCheckMiddleware(logger zerolog.Logger) Middleware {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			sourceType := r.Header.Get("Source-Type")
			if _, ok := validSourceTypes[sourceType]; !ok {
				logger.Warn().Msg("Missing Source-Type header")
				utils.WriteJSONError(logger, w, http.StatusBadRequest, utils.ErrMissingSourceType)

				return
			}
			next(w, r, ps)
		}
	}
}

func ErrorHandlingMiddleware(logger zerolog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				logger.Error().Interface("error", err).Msg("Unhandled error occurred")
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
