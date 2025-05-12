package response

import (
	"encoding/json"
	"errors"
	"github.com/xoesae/cid-api/pkg/fault"
	"github.com/xoesae/cid-api/pkg/logger"
	"net/http"
)

type errorResponse struct {
	Error string `json:"error"`
}

func Json(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, err error) {
	logger.Get().Error(err.Error(), "err", err)

	var f *fault.Fault
	if errors.As(err, &f) {
		Json(w, f.HttpStatusCode, f)
		return
	}

	// Unhandled error -> internal server error
	Json(w, http.StatusInternalServerError, fault.NewFault("internal server error", nil, http.StatusInternalServerError))
}
