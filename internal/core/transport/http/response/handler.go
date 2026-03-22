package core_http_response

import (
	"encoding/json"
	"fmt"
	"net/http"

	core_logger "github.com/fenzit/golang-todoapp/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResposeHandler struct {
	log *core_logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponseHandler(
	log *core_logger.Logger,
	rw http.ResponseWriter,
) *HTTPResposeHandler {
	return &HTTPResposeHandler{
		log: log,
		rw:  rw,
	}
}

func (h *HTTPResposeHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("Unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.rw.WriteHeader(statusCode)

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}
	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("write http response", zap.Error(err))
	}
}
