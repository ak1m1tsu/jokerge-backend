package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type ValidateUserReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ValidateUserCredentials валидирует данные пользователя для использования API
//
//	@Summary	валидация пользовательских данных
//	@Tags		auth
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	UserInfoItem
//	@Router		/api/auth [post]
func (e *Env) ValidateUserCredentials(ctx *fiber.Ctx) error {
	req := new(ValidateUserReq)
	if err := ctx.BodyParser(req); err != nil {
		zerolog.Ctx(ctx.Context()).Error().Err(err).Msg("failed to parse request body")
		return err
	}

	uinfo, ok, err := e.Service().ValidateUser(ctx.Context(), req.Email, req.Password)
	if err != nil {
		zerolog.Ctx(ctx.Context()).Error().Err(err).Msg("failed to validate user credentials")
		return err
	}

	if !ok {
		return ctx.Status(http.StatusForbidden).JSON(map[string]string{
			"error": "invalid credentials",
		})
	}

	return ctx.JSON(uinfo)
}
