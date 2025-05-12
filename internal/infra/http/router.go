package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/xoesae/cid-api/internal/infra/http/handler"
)

func RegisterCidRoutes(r chi.Router, h *handler.CidHandler) {
	r.Route("/api/v1/cid", func(r chi.Router) {
		r.Get("/", h.HandleGetCidPaginated)
		r.Get("/{code}", h.HandleGetCidByCode)
	})
}
