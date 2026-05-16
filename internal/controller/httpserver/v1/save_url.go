package v1

import (
	"context"
	"errors"
	"log/slog"
	"net/http"

	"github.com/Kbnh/url_shortener/internal/domain"
	"github.com/Kbnh/url_shortener/internal/dto"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

type URLSaver interface {
	SaveURL(ctx context.Context, url, alias string) (int64, error)
}

func SaveURL(log *slog.Logger, uc URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("fn", "SaveURL"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req dto.SaveURLRequest

		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("render.DecodeJSON", slog.Any("error", err))
			render.JSON(w, r, dto.Error("invalid request body"))
			return
		}

		if err := validator.New().Struct(req); err != nil {
			log.Error("validator.New.Struct", slog.Any("error", err))
			render.JSON(w, r, dto.Error("validation failed"))
			return
		}

		res, err := uc.SaveURL(r.Context(), req.URL, req.Alias)
		if err != nil {
			if errors.Is(err, domain.ErrURLExists) {
				log.Info("url already exists", slog.String("url", req.URL))
				render.JSON(w, r, dto.Error("url already exists"))
				return
			}
			log.Error("usecase error", slog.Any("error", err))
			render.JSON(w, r, dto.Error("internal error"))
			return
		}

		log.Info("url saved", slog.Int64("id", res))
		render.JSON(w, r, dto.SaveURLResponse{
			Params: dto.OK(),
			ID:     res,
		})
	}
}
