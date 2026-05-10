package v1

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/Kbnh/url_shortener/internal/domain"
	"github.com/Kbnh/url_shortener/internal/dto"
	"github.com/Kbnh/url_shortener/internal/seeder"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/go-playground/validator/v10"
)

const aliasLenght = 6

type Request struct {
	URL   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	Params dto.Params
	Alias  string `json:"alias,omitempty"`
}

type URLSaver interface {
	SaveURL(urlToSave, alias string) (int64, error)
}

func SaveURL(log *slog.Logger, urlSaver URLSaver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log = log.With(
			slog.String("fn", "SaveURL"),
			slog.String("request_id", middleware.GetReqID(r.Context())),
		)

		var req Request

		if err := render.DecodeJSON(r.Body, &req); err != nil {
			log.Error("render.DecodeJSON", err.Error())

			render.JSON(w, r, dto.Error("failed request decode"))

			return
		}

		log.Info("Request body decoded", slog.Any("request", req))

		if err := validator.New().Struct(req); err != nil {
			log.Error("validator.New.Struct", err.Error())

			render.JSON(w, r, dto.Error("failed request validation"))

			return
		}

		alias := req.Alias
		if alias == "" {
			alias = seeder.NewRandomString(aliasLenght)
		}

		id, err := urlSaver.SaveURL(req.URL, alias)
		if err != nil {
			if errors.Is(err, domain.ErrURLExists) {
				log.Info("url already exists", slog.String("url", req.URL))

				render.JSON(w, r, dto.Error("url already exists"))

				return
			}
			log.Error("urlSaver.SaveURL", err.Error())

			render.JSON(w, r, dto.Error("failed to save url"))

			return
		}

		log.Info("url added", slog.Int64("id", id))

		render.JSON(w, r, Response{
			Params: dto.OK(),
			Alias:  alias,
		})
	}
}
