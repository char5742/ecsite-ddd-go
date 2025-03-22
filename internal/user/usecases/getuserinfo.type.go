package userusecases

import (
	"context"

	identitytypes "github.com/char5742/ecsite-ddd-go/internal/identity/domain/types"
	userdomain "github.com/char5742/ecsite-ddd-go/internal/user/domain"
	userworkflows "github.com/char5742/ecsite-ddd-go/internal/user/workflows"
)

// ユーザー情報取得ユースケース
type GetUserInfoUsecase func(LoadhUserAggregate, userworkflows.GetUserInfoWorkflow) func(context.Context,
	struct{ ID string }) (*userworkflows.UserInfo, error)

// ユーザー集約から取得
type LoadhUserAggregate func(
	userID identitytypes.IdentityID,
) (userdomain.User, error)
