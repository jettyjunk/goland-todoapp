package core_http_response

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	core_errors "github.com/jettyjunk/goland-todoapp/internal/core/errors"
	core_logger "github.com/jettyjunk/goland-todoapp/internal/core/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *core_logger.Logger
	w   http.ResponseWriter
}

func NewHTTPResponseHandler(log *core_logger.Logger, w http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		w:   w,
	}
}

func (h *HTTPResponseHandler) HTMLResponse(html []byte) {
	h.w.WriteHeader(http.StatusOK)

	h.w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := h.w.Write(html); err != nil {
		h.log.Error("write HTML HTTP response", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) NoContentResponse() {
	h.w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPResponseHandler) JsonResponse(reponseBody any, statusCode int) {
	h.w.WriteHeader(statusCode)

	if err := json.NewEncoder(h.w).Encode(reponseBody); err != nil {
		h.log.Error("write HTTP response", zap.Error(err))
	}
}

func (h *HTTPResponseHandler) ErrorResponse(err error, msg string) {
	var (
		statusCode int
		logFunc    func(string, ...zap.Field)
	)

	switch {
	case errors.Is(err, core_errors.ErrInvalidArgument):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn
	case errors.Is(err, core_errors.ErrNotFounde):
		statusCode = http.StatusNotFound
		logFunc = h.log.Debug
	case errors.Is(err, core_errors.ErrConflict):
		statusCode = http.StatusBadRequest
		logFunc = h.log.Warn

	default:
		statusCode = http.StatusInternalServerError
		logFunc = h.log.Error
	}

	logFunc(msg, zap.Error(err))
	h.errorResponse(statusCode, msg, err)
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	statusCode := http.StatusInternalServerError
	err := fmt.Errorf("unexpected panic: %v", p)

	h.log.Error(msg, zap.Error(err))
	h.errorResponse(statusCode, msg, err)
}

func (h *HTTPResponseHandler) errorResponse(statusCode int, msg string, err error) {
	response := ErrorResponse{
		Error:   err.Error(),
		Message: msg,
	}
	h.JsonResponse(response, statusCode)
}
