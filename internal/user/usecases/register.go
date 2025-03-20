package userusecases

import (
	shareinterfaces "github.com/char5742/ecsite-ddd-go/internal/share/domain/interfaces"
	sharetypes "github.com/char5742/ecsite-ddd-go/internal/share/domain/types"
	userdomain "github.com/char5742/ecsite-ddd-go/internal/user/domain"
	usertypes "github.com/char5742/ecsite-ddd-go/internal/user/types"
	userworkflows "github.com/char5742/ecsite-ddd-go/internal/user/types/workflows"
)

func NewRegisterUserUsecase(
	taken userdomain.IsEmailTaken,
	flow userworkflows.RegisterUser,
) func(sharetypes.Command[usertypes.RegisterUserCommand]) ([]shareinterfaces.Event, error) {

	return func(cmd sharetypes.Command[usertypes.RegisterUserCommand]) ([]shareinterfaces.Event, error) {
		// ユーザー登録のビジネスロジックを実装する
		// ここでは、ユーザー登録の処理を行い、イベントを生成して返す
		// 例えば、ユーザー登録成功イベントなど

		validateUser := func(user userdomain.UnvalidatedUser) (*userdomain.ValidatedUser, error) {
			selfValidatedUser, err := userdomain.NewSelfValidateUser(user)
			if err != nil {
				return nil, err
			}
			validatedUser, err := userdomain.NewValidatedUser(
				taken,
				*selfValidatedUser)
			if err != nil {
				return nil, err
			}
			return validatedUser, nil
		}
		registUser := func(user userdomain.ValidatedUser) (*userdomain.RegistedUser, error) {

			return &userdomain.RegistedUser{
				ID:             user.ID,
				FirstName:      user.FirstName,
				LastName:       user.LastName,
				Email:          userdomain.Email{user.UniqueEmail},
				Password:       user.Password,
				Zipcode:        user.Zipcode,
				Prefecture:     user.Prefecture,
				Municipalities: user.Municipalities,
				Address:        user.Address,
				Telephone:      user.Telephone,
			}, nil
		}

		events, err := flow(validateUser, registUser)(cmd.Data.UnvalidatedUser)
		if err != nil {
			return nil, err
		}

		slice := make([]shareinterfaces.Event, 0, len(events))
		for i, event := range events {
			slice[i] = event
		}
		return slice, nil
	}

}
