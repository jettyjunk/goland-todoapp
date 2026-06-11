package users_transport_http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/jettyjunk/goland-todoapp/internal/core/domain"
	core_logger "github.com/jettyjunk/goland-todoapp/internal/core/logger"
	core_http_request "github.com/jettyjunk/goland-todoapp/internal/core/transport/http/request"
	core_http_response "github.com/jettyjunk/goland-todoapp/internal/core/transport/http/response"
	core_http_types "github.com/jettyjunk/goland-todoapp/internal/core/transport/http/types"
)

type PatchUserRequest struct {
	FullName    core_http_types.Nullable[string] `json:"full_name" swaggertype:"string" example:"Jetty Junk"`
	PhoneNumber core_http_types.Nullable[string] `json:"phone_number" swaggertype:"string" example:"+79998887766"`
}

func (r *PatchUserRequest) Validate() error {
	if r.FullName.Set {
		if r.FullName.Value == nil {
			return fmt.Errorf("`FullName` can't be NULL")
		}

		fullNameLen := len([]rune(*r.FullName.Value))
		if fullNameLen < 3 || fullNameLen > 100 {
			return fmt.Errorf("`FullName must be between 3 and 100 sybols`")
		}
	}

	if r.PhoneNumber.Set {
		if r.PhoneNumber.Value != nil {
			phoneNumberLen := len([]rune(*r.PhoneNumber.Value))
			if phoneNumberLen < 10 || phoneNumberLen > 15 {
				return fmt.Errorf("`PhoneNumber must be between 10 and 15 symbols`")
			}

			if !strings.HasPrefix(*r.PhoneNumber.Value, "+") {
				return fmt.Errorf("`PhoneNumber` must startswith '+' symbol")
			}
		}
	}

	return nil
}

type PatchUserResponse UserDTOResponse

// PatchUser 	godoc
// @Summary 	Изменения пользователя
// @Description Измение информации об уже существующем в системе пользователя
// @Description ### Логика обновления полей (Theree-state logic):
// @Description 1.**Поле не передано**: `phone_number` игнорируется, значение в БД не меняется
// @Description 2.**Переданно значение**: `phone_number` "`+71112223344`" - устанавливает новый `phone_number` в БД
// @Description 3.**Переданн null**: `"phone_number":null` - очищает поле в БД
// @Description Ограничение `full_name` не может быть как null
// @Tags 		users
// @Accept 		json
// @Produce 	json
// @Param 		id path int true  										 "ID изменяемого пользователя"
// @Param 		request body 			PatchUserRequest true 			 "PatchUser тело запроса"
// @Success 	200 		{object} 	PatchUserResponse 				 "Успешно изменённый пользователь"
// @Failure		400 		{object} 	core_http_response.ErrorResponse "Bad Request"
// @Failure		404 		{object} 	core_http_response.ErrorResponse "User not found"
// @Failure		409 		{object} 	core_http_response.ErrorResponse "Conflict"
// @Failure 	500 		{object} 	core_http_response.ErrorResponse "Internal server error"
// @Router 		/users/{id} [patch]
func (h *UsersHTTPHandler) PatchUser(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(log, w)

	userID, err := core_http_request.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to get userID path value")

		return
	}

	var request PatchUserRequest
	if err := core_http_request.DecodeAndValidateRequest(r, &request); err != nil {
		responseHandler.ErrorResponse(err, "failed to decode and validate HTTP request")
		return
	}

	userPatch := userPatchFromRequest(request)

	userDomain, err := h.usersService.PatchUser(ctx, userID, userPatch)
	if err != nil {
		responseHandler.ErrorResponse(err, "failed to patch user")

		return
	}

	response := PatchUserResponse(userDTOFromDomain(userDomain))

	responseHandler.JsonResponse(response, http.StatusOK)
}

func userPatchFromRequest(request PatchUserRequest) domain.UserPatch {
	return domain.NewUserPatch(request.FullName.ToDomain(), request.PhoneNumber.ToDomain())
}
