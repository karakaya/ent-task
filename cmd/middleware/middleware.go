package middleware

import (
	"github.com/julienschmidt/httprouter"
	"github.com/rs/zerolog"
	"net/http"
)

type Middleware func(httprouter.Handle) httprouter.Handle

func SourceTypeCheckMiddleware(logger zerolog.Logger) Middleware {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
			sourceType := r.Header.Get("Source-Type")
			if sourceType != "game" && sourceType != "server" && sourceType != "payment" {
				logger.Warn().Msg("Missing Source-Type header")
				http.Error(w, "Missing Source-Type header", http.StatusBadRequest)
				return
			}

			logger.Info().Str("Source-Type", sourceType).Msg("Source-Type header found")
			next(w, r, ps)
		}
	}
}
