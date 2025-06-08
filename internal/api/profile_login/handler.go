package profile_login

import (
	"context"
	"github.com/AlekSi/pointer"
	"github.com/yasonofriychuk/real-estate-insight/internal/generated/api"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/errors"
	"github.com/yasonofriychuk/real-estate-insight/internal/infrastructure/logger"
	"net/http"
)

type ProfileLoginHandler struct {
	log            logger.Log
	jwtAuth        jwtAuth
	profileStorage profileStorage
}

func New(log logger.Log, jwtAuth jwtAuth, profileStorage profileStorage) *ProfileLoginHandler {
	return &ProfileLoginHandler{
		log:            log,
		jwtAuth:        jwtAuth,
		profileStorage: profileStorage,
	}
}

func (h *ProfileLoginHandler) UserLogin(ctx context.Context, request *api.UserLoginReq) (api.UserLoginRes, error) {
	profileId, err := h.profileStorage.Login(ctx, request.Email, request.Password)
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("userLogin failed getting profileId")
		return pointer.To(api.UserLoginInternalServerError(
			errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		)), nil
	}

	if profileId == nil {
		return pointer.To(api.UserLoginUnauthorized(
			errors.BuildError(http.StatusUnauthorized, "Ошибка авторизации"),
		)), nil
	}

	token, err := h.jwtAuth.CreatePermanentToken(*profileId)
	if err != nil {
		h.log.WithContext(ctx).WithError(err).Error("userLogin failed getting token")
		return pointer.To(api.UserLoginInternalServerError(
			errors.BuildError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)),
		)), nil
	}

	return &api.UserLoginOK{
		ProfileID: *profileId,
		Token:     token,
	}, nil
}
