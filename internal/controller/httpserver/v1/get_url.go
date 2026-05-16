package v1

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/Kbnh/url_shortener/internal/domain"
	"github.com/Kbnh/url_shortener/internal/dto"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type URLGetter interface {
	GetURL(ctx context.Context, alias string) (string, error)
}

func GetURL(log *slog.Logger, uc URLGetter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("fn", "GetURL"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req string

		req = chi.URLParam(r, "alias")
		if req == "" {
			log.Error("chi.URLParam", slog.Any("error", domain.ErrInvalidInput))

			render.JSON(w, r, dto.Error("invalid input"))

			return
		}

		log.Info("Request body decoded", slog.Any("request", req))

		res, err := uc.GetURL(r.Context(), req)
		if err != nil {
			if errors.Is(err, domain.ErrURLNotFound) {
				log.Info("url not found", slog.String("alias", req))

				render.JSON(w, r, dto.Error("not found"))

				return
			}
			log.Error("uc.GetURL", slog.Any("error", err))

			render.JSON(w, r, dto.Error("failed to get url"))

			return
		}

		log.Info("got url", slog.String("url", res))

		http.Redirect(w, r, res, http.StatusFound)
	}
}
