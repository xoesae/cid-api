package handler

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/xoesae/cid-api/internal/domain/service"
	"github.com/xoesae/cid-api/internal/infra/http/response"
	"github.com/xoesae/cid-api/pkg/logger"
	"net/http"
	"strconv"
)

type Cid struct {
	Code string `json:"code"`
}

type CidHandler struct {
	CidService service.CidService
}

func (h *CidHandler) HandleGetCidPaginated(w http.ResponseWriter, r *http.Request) {
	logger.Get().Debug("HandleGetCidPaginated")

	ctx := context.Background()

	page := 1
	pageSize := 10

	query := r.URL.Query()

	if p, err := strconv.Atoi(query.Get("page")); err == nil && p > 0 {
		page = p
	}
	if ps, err := strconv.Atoi(query.Get("page_size")); err == nil && ps > 0 {
		pageSize = ps
	}

	search := query.Get("search")

	paginated, err := h.CidService.GetAllPaginated(ctx, pageSize, page, search)
	if err != nil {
		logger.Get().Error(err.Error())

		response.Error(w, err)
		return
	}

	logger.Get().Debug("Returning from GetCidPaginated", "result", paginated)

	response.Json(w, http.StatusOK, paginated)
}

func (h *CidHandler) HandleGetCidByCode(w http.ResponseWriter, r *http.Request) {
	logger.Get().Debug("HandleGetCidByCode")

	ctx := context.Background()

	code := chi.URLParam(r, "code")

	result, err := h.CidService.GetByCode(ctx, code)
	if err != nil {
		logger.Get().Error(err.Error())
		response.Error(w, err)
		return
	}

	response.Json(w, http.StatusOK, result)
}
