package userusecases

import (
	"context"

	sharetypes "github.com/char5742/ecsite-ddd-go/internal/share/domain/types"
	userdomain "github.com/char5742/ecsite-ddd-go/internal/user/domain"
	usertypes "github.com/char5742/ecsite-ddd-go/internal/user/types"
	userworkflows "github.com/char5742/ecsite-ddd-go/internal/user/types/workflows"
)

func NewGetUserInfoUsecase(fetch usertypes.FetchUserAggregate, getUserInfo userworkflows.GetUserInfo) func(context.Context, sharetypes.Query[usertypes.GetUserInfoQuery]) (*userworkflows.UserInfo, error) {

	return func(ctx context.Context, query sharetypes.Query[usertypes.GetUserInfoQuery]) (*userworkflows.UserInfo, error) {
		user, err := fetch(query.Data.ID)
		if err != nil {
			return nil, err
		}

		return getUserInfo(NewUsetToInfo())(user)
	}
}

func NewUsetToInfo() userworkflows.UserToInfo {
	return func(user userdomain.User) userworkflows.UserInfo {
		return userworkflows.UserInfo{
			ID:             user.ID,
			FirstName:      user.FirstName,
			LastName:       user.LastName,
			Email:          user.Email,
			Zipcode:        user.Zipcode,
			Prefecture:     user.Prefecture,
			Municipalities: user.Municipalities,
			Address:        user.Address,
			Telephone:      user.Telephone,
			AuditInfo:      user.AuditInfo,
		}
	}
}
