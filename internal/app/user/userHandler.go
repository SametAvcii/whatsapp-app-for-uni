package user

import (
	"net/http"
	"whatsapp-app/dto/request"
	userService "whatsapp-app/internal/service/user"
	utils "whatsapp-app/internal/utils"
	"whatsapp-app/internal/utils/response"
	"whatsapp-app/internal/utils/validate"

	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	Register(ctx echo.Context) error
	Login(ctx echo.Context) error
	//SendEmail(ctx echo.Context) error
	VerifyEmail(ctx echo.Context) error
}

type UserHandler struct {
	service userService.IUserService
	utils   utils.IUtils
}

func NewUserHandler(service userService.IUserService, utils utils.IUtils) IUserHandler {
	return &UserHandler{
		service: service,
		utils:   utils,
	}
}

func (h *UserHandler) Register(c echo.Context) error {
	var request request.UserRegisterDTO
	if validate.Validator(&c, &request) != nil {
		return nil
	}
	ctx := c.Request().Context()

	user, err := h.service.Register(ctx, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response(http.StatusBadRequest, err.Error()))

	}
	return c.JSON(http.StatusOK, response.Response(http.StatusOK, user))
}

func (h *UserHandler) Login(c echo.Context) error {
	var request request.UserLoginDTO
	if validate.Validator(&c, &request) != nil {
		return nil
	}
	ctx := c.Request().Context()

	user, err := h.service.Login(ctx, request)

	if user.Email == "" {
		return c.JSON(http.StatusBadRequest, response.Response(http.StatusBadRequest, err.Error()))
	}

	if !user.Verified {
		return c.JSON(http.StatusForbidden, response.Response(http.StatusForbidden, user.Email))
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response(http.StatusBadRequest, err.Error()))
	}
	return c.JSON(http.StatusOK, response.Response(http.StatusOK, user))
}

/*func (h *UserHandler) SendEmail(c echo.Context) error {
	var request request.UserVerifyDTO
	if validate.Validator(&c, &request) != nil {
		return nil
	}
	ctx := c.Request().Context()

	message, err := h.service.SendVerifyEmail(ctx, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response(http.StatusBadRequest, err.Error()))

	}
	return c.JSON(http.StatusOK, response.Response(http.StatusOK, message))
}*/

func (h *UserHandler) VerifyEmail(c echo.Context) error {
	var request request.UserVerifyEmailDTO
	if validate.Validator(&c, &request) != nil {
		return nil
	}
	ctx := c.Request().Context()

	message, err := h.service.VerifyUserEmail(ctx, request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Response(http.StatusBadRequest, err.Error()))

	}
	return c.JSON(http.StatusOK, response.Response(http.StatusOK, message))
}
