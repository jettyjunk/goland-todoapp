REST API для управления задачами (todo-приложение), написанное на Go.

## Технологии

- **Go 1.25** — основной язык
- **PostgreSQL 18** — база данных
- **Docker / Docker Compose** — контейнеризация
- **pgx / pgxpool** — драйвер PostgreSQL
- **golang-migrate** — миграции базы данных
- **zap** — логирование
- **Swagger** — документация API

## Архитектура

Проект построен на трёхслойной архитектуре:

```
Transport (HTTP) → Service (бизнес логика) → Repository (база данных)
```

## Требования

- Docker и Docker Compose
- Go 1.25+
- Make

## Запуск

### 1. Клонировать репозиторий

```bash
git clone https://github.com/jettyjunk/goland-todoapp.git
cd goland-todoapp
```

### 2. Создать файл `.env`

```env
HTTP_ADDR=:5050
HTTP_SHUTDOWN_TIMEOUT=30s

ALLOWED_ORIGINS=http://localhost:5050,null

POSTGRES_USER=admin
POSTGRES_PASSWORD=123
POSTGRES_DB=test-db
POSTGRES_TIMEOUT=10s

LOGGER_LEVEL=DEBUG

TIME_ZONE=UTC
```

### 3. Запустить базу данных

```bash
make env-up
```

### 4. Применить миграции

```bash
make migrate-up
```

### 5. Запустить приложение

```bash
make todoapp-run
```

Приложение будет доступно на `http://localhost:5050`

Swagger документация: `http://localhost:5050/swagger/`

## API эндпоинты

### Пользователи

| Метод | Путь | Описание |
|-------|------|----------|
| POST | /api/v1/users | Создать пользователя |
| GET | /api/v1/users | Список пользователей |
| GET | /api/v1/users/{id} | Получить пользователя |
| PATCH | /api/v1/users/{id} | Обновить пользователя |
| DELETE | /api/v1/users/{id} | Удалить пользователя |

### Задачи

| Метод | Путь | Описание |
|-------|------|----------|
| POST | /api/v1/tasks | Создать задачу |
| GET | /api/v1/tasks | Список задач |
| GET | /api/v1/tasks/{id} | Получить задачу |
| PATCH | /api/v1/tasks/{id} | Обновить задачу |
| DELETE | /api/v1/tasks/{id} | Удалить задачу |

### Статистика

| Метод | Путь | Описание |
|-------|------|----------|
| GET | /api/v1/statistics | Статистика по задачам |

## Команды Makefile

```bash
make env-up            # запустить базу данных
make env-down          # остановить базу данных
make env-cleanup       # очистить данные базы
make migrate-up        # применить миграции
make migrate-down      # откатить миграции
make todoapp-run       # запустить приложение локально
make todoapp-deploy    # задеплоить через Docker
make todoapp-undeploy  # остановить задеплоенное приложение
make swagger-gen       # сгенерировать Swagger документацию
make swagger-build     # пересобрать swagger контейнер
make logs-cleanup      # очистить логи
make ps                # статус контейнеров
```