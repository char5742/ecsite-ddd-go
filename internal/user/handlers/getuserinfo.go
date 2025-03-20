package userhandlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	sharetypes "github.com/char5742/ecsite-ddd-go/internal/share/domain/types"
	userdomain "github.com/char5742/ecsite-ddd-go/internal/user/domain"
	userpersistence "github.com/char5742/ecsite-ddd-go/internal/user/persistence"
	usertypes "github.com/char5742/ecsite-ddd-go/internal/user/types"
	userworkflows "github.com/char5742/ecsite-ddd-go/internal/user/types/workflows"
	userusecases "github.com/char5742/ecsite-ddd-go/internal/user/usecases"
)

type GetUserInfoHandler struct{}

func (h GetUserInfoHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	query, err := h.req2query(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fetchUserAggregate := userpersistence.NewFetchUserAggregate(ctx)
	workflow := userusecases.NewGetUserInfoUsecase(
		fetchUserAggregate, func(uti userworkflows.UserToInfo) userworkflows.GetUserInfoWorkflow {
			return func(user userdomain.User) (*userworkflows.UserInfo, error) {
				userInfo := uti(user)
				return &userInfo, nil
			}
		})

	userInfo, err := workflow(ctx, query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(userInfo); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

func (h GetUserInfoHandler) req2query(r *http.Request) (sharetypes.Query[usertypes.GetUserInfoQuery], error) {
	var query usertypes.GetUserInfoQuery
	if err := json.NewDecoder(r.Body).Decode(&query); err != nil {
		return sharetypes.Query[usertypes.GetUserInfoQuery]{}, fmt.Errorf("bad request: %w", err)
	}
	if err := h.validate(query); err != nil {
		return sharetypes.Query[usertypes.GetUserInfoQuery]{}, err
	}
	return sharetypes.Query[usertypes.GetUserInfoQuery]{Data: query}, nil
}

func (GetUserInfoHandler) validate(raw usertypes.GetUserInfoQuery) error {
	return nil
}
