package statistics_transport_http

import (
	"fmt"
	"net/http"
	"time"

	core_logger "github.com/jettyjunk/goland-todoapp/internal/core/logger"
	core_http_request "github.com/jettyjunk/goland-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/jettyjunk/goland-todoapp/internal/core/transport/http/response"
)

type GetStatisticsResponse StatisticsDTOResponse

func (h *StatisticsHTTPHandler) GetStatistics(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, from, to, err := getUserIDFromTOQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID/from/to query param")

		return
	}

	statistics, err := h.statisticsService.GetStatistics(ctx, userID, from, to)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get statistics")

		return
	}

	response := GetStatisticsResponse(toDTOFromDomain(statistics))

	responseHandler.JsonResponse(response, http.StatusOK)
}

func getUserIDFromTOQueryParams(r *http.Request) (*int, *time.Time, *time.Time, error) {
	const (
		userIDQueryParamKey = "user_id"
		fromQueryParamKey   = "from"
		toQueryParamKey     = "to"
	)

	userID, err := core_http_request.GetIntQueryParam(r, userIDQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'user_id' query param: %w'", err)
	}

	from, err := core_http_request.GetDateQueryParam(r, fromQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'from' quert param: %w", err)
	}

	to, err := core_http_request.GetDateQueryParam(r, toQueryParamKey)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("get 'to' quert param: %w", err)
	}

	return userID, from, to, nil
}
