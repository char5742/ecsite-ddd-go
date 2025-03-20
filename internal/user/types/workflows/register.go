package userworkflows

import userdomain "github.com/char5742/ecsite-ddd-go/internal/user/domain"

// ユーザー登録
type RegisterUser func(ValidateUser, RegistUser) RegisterUserWorkflow

// ユーザー登録ワークフロー
type RegisterUserWorkflow func(user userdomain.UnvalidatedUser) ([]RegisterUserEvent, error)

// ユーザーの検証ステップ
type ValidateUser func(user userdomain.UnvalidatedUser) (*userdomain.ValidatedUser, error)

// ユーザーの登録ステップ
type RegistUser func(user userdomain.ValidatedUser) (*userdomain.RegistedUser, error)

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
