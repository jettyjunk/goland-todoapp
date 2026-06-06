package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	core_logger "github.com/jettyjunk/goland-todoapp/internal/core/logger"
	core_pgx_pool "github.com/jettyjunk/goland-todoapp/internal/core/repository/postgres/pool/pgx"
	core_http_middleware "github.com/jettyjunk/goland-todoapp/internal/core/transport/http/middleware"
	core_http_server "github.com/jettyjunk/goland-todoapp/internal/core/transport/http/server"
	tasks_postgres_repository "github.com/jettyjunk/goland-todoapp/internal/features/tasks/repository/postgres"
	tasks_service "github.com/jettyjunk/goland-todoapp/internal/features/tasks/service"
	tasks_transport_http "github.com/jettyjunk/goland-todoapp/internal/features/tasks/transport/http"
	users_postgres_repository "github.com/jettyjunk/goland-todoapp/internal/features/users/repository/postgres"
	users_service "github.com/jettyjunk/goland-todoapp/internal/features/users/service"
	users_transport_http "github.com/jettyjunk/goland-todoapp/internal/features/users/transport/http"
	"go.uber.org/zap"
)

var (
	timeZone = time.UTC
)

func main() {
	time.Local = timeZone

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	config := core_logger.NewConfigMast()
	logger, err := core_logger.NewLogger(config)
	if err != nil {
		fmt.Println("created logger:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Debug("initizaling postgres connection pool")
	pool, err := core_pgx_pool.NewPool(ctx, core_pgx_pool.NewConfigMust())
	if err != nil {
		logger.Fatal("failed to init postgres connection pool: %w", zap.Error(err))
	}
	defer pool.Close()

	logger.Debug("application time zone", zap.Any("zone", timeZone))

	logger.Debug("initizaling feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersSevice(usersRepository)
	userTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)

	logger.Debug("initizaling feature", zap.String("featire", "tasks"))
	tasksRepository := tasks_postgres_repository.NewTasksRepository(pool)
	tasksService := tasks_service.NewTasksService(tasksRepository)
	tasksTransportHTTP := tasks_transport_http.NewTasksHTTPHandler(tasksService)

	logger.Debug("initizaling HTTP server")
	httpServer := core_http_server.NewHTTPServer(
		core_http_server.NewConfigMust(),
		logger,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(logger),
		core_http_middleware.Trace(),
		core_http_middleware.Panic(),
	)
	apiVersionRouterV1 := core_http_server.NewAPIVersionRouter(core_http_server.ApiVersion1)
	apiVersionRouterV1.RegisterRoutes(userTransportHTTP.Routes()...)
	apiVersionRouterV1.RegisterRoutes(tasksTransportHTTP.Routes()...)
	httpServer.RegisterAPIRouter(apiVersionRouterV1)

	if err := httpServer.Run(ctx); err != nil {
		logger.Error("HTTP server run error", zap.Error(err))
	}
}
