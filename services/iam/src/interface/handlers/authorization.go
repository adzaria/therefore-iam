package handlers

import (
	"encoding/json"
	"fmt"
	"iam/pkg/apperrors"
	"iam/pkg/httphelpers"
	"iam/src/core/ports/primaryports"
	"net/http"
)

type AuthorizationHandler struct {
	authorizationService primaryports.AuthorizationService
}

func NewAuthorizationHandler(authorizationService primaryports.AuthorizationService) *AuthorizationHandler {
	return &AuthorizationHandler{
		authorizationService: authorizationService,
	}
}

func (h *AuthorizationHandler) Refresh(w http.ResponseWriter, r *http.Request) {
	var args primaryports.RefreshAccessTokenArgs

	err := json.NewDecoder(r.Body).Decode(&args)

	if err != nil {
		fmt.Println(err)
		httphelpers.WriteError(http.StatusBadRequest, "error", apperrors.ServerError)(w, r)
		return
	}

	answer, err := h.authorizationService.RefreshAccessToken(args)

	if err != nil {
		switch err.Error() {
		case apperrors.InvalidRefreshToken:
			httphelpers.WriteError(http.StatusForbidden, "error", apperrors.InvalidAccessToken)(w, r)
		default:
			fmt.Println(err)
			httphelpers.WriteError(http.StatusInternalServerError, "error", apperrors.ServerError)(w, r)
		}
		return
	}

	httphelpers.WriteSuccess(http.StatusAccepted, "Refresh token accepted", answer)(w, r)
}
