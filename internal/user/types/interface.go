package usertypes

import (
	"context"

	identitytypes "github.com/char5742/ecsite-ddd-go/internal/identity/domain/types"
	shareinterfaces "github.com/char5742/ecsite-ddd-go/internal/share/domain/interfaces"
	sharetypes "github.com/char5742/ecsite-ddd-go/internal/share/domain/types"
	userdomain "github.com/char5742/ecsite-ddd-go/internal/user/domain"
	userworkflows "github.com/char5742/ecsite-ddd-go/internal/user/types/workflows"
)

type UserCommand = sharetypes.Command[userCommand]

type userCommand interface {
	isUserCommand()
}

// ユーザー登録コマンド
type RegisterUserCommand struct {
	userdomain.UnvalidatedUser
}

func (RegisterUserCommand) isUserCommand() {}

// ユーザー登録を行うユースケース
type RegisterUserUsecase func(userworkflows.RegisterUser) func(sharetypes.Command[RegisterUserCommand]) ([]shareinterfaces.Event, error)

// UserQuery はユーザー関連の問い合わせを表す型です。
// クエリは読み取り操作を表し、システムの状態を変更しません。
type UserQuery = sharetypes.Query[userQuery]

// userQuery はユーザークエリのマーカーインターフェースです。
// このインターフェースを実装する型はユーザークエリとして扱われます。
type userQuery interface {
	isUserQuery()
}

// GetUserInfoQuery はユーザー情報を取得するクエリです。
// 指定されたIDを持つユーザーの情報を取得します。
type GetUserInfoQuery struct {
	// ID は情報を取得するユーザーの一意識別子です
	ID string `json:"id"`
}

func (GetUserInfoQuery) isUserQuery() {}

// ユーザー情報取得ユースケース
type GetUserInfoUsecase func(FetchUserAggregate, userworkflows.GetUserInfo) func(context.Context, sharetypes.Query[GetUserInfoQuery]) (*userworkflows.UserInfo, error)

// ユーザー集約から取得
type FetchUserAggregate func(
	userID identitytypes.IdentityID,
) (userdomain.User, error)
