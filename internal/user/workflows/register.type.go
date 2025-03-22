package userworkflows

import (
	sharetypes "github.com/char5742/ecsite-ddd-go/internal/share/domain/types"
	shareutils "github.com/char5742/ecsite-ddd-go/internal/share/utils"
	userdomain "github.com/char5742/ecsite-ddd-go/internal/user/domain"
)

// ユーザー登録コマンド
type RegisterUserCommand sharetypes.Command[RegisterUserCommandData]
type RegisterUserCommandData struct {
	userdomain.UnvalidatedUser
	userdomain.ExternalUserData
}

// ユーザー登録ワークフロー
type RegisterUserWorkflow func(RegisterUserCommand) ([]RegisterUserEvent, shareutils.DomainValidationResult, error)

// ユーザー登録
type RegisterUser func(ValidateUser, RegistUser) RegisterUserWorkflow

// ユーザーのローカル検証ステップ
type ValidateUser func(userdomain.UnvalidatedUser, userdomain.ExternalUserData) (*userdomain.ValidatedUser, shareutils.DomainValidationResult)

// ユーザーの登録ステップ
type RegistUser func(userdomain.ValidatedUser) (*userdomain.RegistedUser, error)

type RegisterUserEvent interface {
	registerUserEvent()
	IsEvent()
}

// ユーザー登録イベント
type UserRegistered struct {
	RegistedUser userdomain.RegistedUser
}

func (UserRegistered) registerUserEvent() {}
func (UserRegistered) IsEvent()           {}
