package statistics_transport_http

import "github.com/jettyjunk/goland-todoapp/internal/core/domain"

type StatisticsDTOResponse struct {
	TasksCreated              int      `json:"tasks_created"`
	TasksCompleted            int      `json:"tasks_completed"`
	TasksCompletedRate        *float64 `json:"tasks_completed_rate"`
	TasksAverageCompletedTime *string  `json:"tasks_average_completed_time"`
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
