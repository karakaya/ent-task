package middleware

import (
	"entain-golang-task/pkg/utils"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"mime"
	"net/http"
)

type Middleware func(httprouter.Handle) httprouter.Handle

var validSourceTypes = map[string]bool{
	"game":    true,
	"server":  true,
	"payment": true,
}

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

func ContentTypeCheckMiddleware(logger zerolog.Logger) Middleware {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			contentType := r.Header.Get("Content-Type")
			mediaType, _, err := mime.ParseMediaType(contentType)
			if err != nil || mediaType != "application/json" {
				logger.Warn().Msg("Invalid or missing Content-Type header")
				utils.WriteJSONError(logger, w, http.StatusUnsupportedMediaType, utils.ErrInvalidContentType)
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

func Chain(h httprouter.Handle, middlewares ...Middleware) httprouter.Handle {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
