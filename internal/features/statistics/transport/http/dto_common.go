package statistics_transport_http

import "github.com/jettyjunk/goland-todoapp/internal/core/domain"

type StatisticsDTOResponse struct {
	TasksCreated              int      `json:"tasks_created" example:"20"`
	TasksCompleted            int      `json:"tasks_completed" example:"10"`
	TasksCompletedRate        *float64 `json:"tasks_completed_rate" example:"50"`
	TasksAverageCompletedTime *string  `json:"tasks_average_completed_time" example:"2m10s"`
}

func toDTOFromDomain(statistics domain.Statistics) StatisticsDTOResponse {
	var avgTime *string

	if statistics.TasksAverageCompletedTime != nil {
		duration := statistics.TasksAverageCompletedTime.String()
		avgTime = &duration
	}

	return StatisticsDTOResponse{
		TasksCreated:              statistics.TasksCreated,
		TasksCompleted:            statistics.TasksCompleted,
		TasksCompletedRate:        statistics.TasksCompletedRate,
		TasksAverageCompletedTime: avgTime,
	}
}
