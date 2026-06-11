package users_transport_http

import (
	"net/http"

	"github.com/jettyjunk/goland-todoapp/internal/core/domain"
	core_logger "github.com/jettyjunk/goland-todoapp/internal/core/logger"
	core_http_request "github.com/jettyjunk/goland-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/jettyjunk/goland-todoapp/internal/core/transport/http/response"
)

type CreateUserRequest struct {
	FullName    string  `json:"full_name" validate:"required,min=3,max=100" example:"Jetty Junk"`
	PhoneNumber *string `json:"phone_number" validate:"omitempty,e164,min=10,max=15" example:"+79998887766"`
}

type CreateUserResponse UserDTOResponse

// CreateUser 	godoc
// @Summary 	Создать пользователя
// @Description Создать нового пользователя в системе
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		request body 		CreateUserRequest true "CreateUser тело запроса"
// @Success 	201 	{object} 	CreateUserResponse "Успешно созданный пользователь"
// @Failure		400 	{object} 	core_http_response.ErrorResponse "Bad Request"
// @Failure 	500 	{object} 	core_http_response.ErrorResponse "Internal server error"
// @Router 		/users 	[post]
func (h *UsersHTTPHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	var request CreateUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userDomain := domaingFromDTO(request)
	userDomain, err := h.usersService.CreateUser(ctx, userDomain)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to created user")
		return
	}

	response := CreateUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JsonResponse(response, http.StatusCreated)
}

func domaingFromDTO(dto CreateUserRequest) domain.User {
	return domain.NewUserUnitialized(dto.FullName, dto.PhoneNumber)
}
