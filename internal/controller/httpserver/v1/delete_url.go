package v1

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/Kbnh/url_shortener/internal/domain"
	"github.com/Kbnh/url_shortener/internal/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type URLDeleter interface {
	DeleteURL(ctx context.Context, id int64) error
}

func DeleteURL(log *slog.Logger, uc URLDeleter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("fn", "DeleteURL"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req int64

		req, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			log.Error("strconv.ParseInt", slog.Any("error", err))
			render.JSON(w, r, dto.Error("invalid input"))
			return
		}

		log.Info("Request body decoded", slog.Int64("request", req))

		if err = uc.DeleteURL(r.Context(), req); err != nil {
			if errors.Is(err, domain.ErrURLNotFound) {
				log.Info("id not found", slog.Int64("id", req))
				w.WriteHeader(http.StatusNoContent)
				return
			}
			log.Error("uc.DeleteURL", slog.Any("error", err))
			render.JSON(w, r, dto.Error("failed to delete url"))
			return
		}

		log.Info("url deleted")
		w.WriteHeader(http.StatusNoContent)

	}
}
