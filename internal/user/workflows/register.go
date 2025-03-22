package userworkflows

import (
	shareerrs "github.com/char5742/ecsite-ddd-go/internal/share/domain/errs"
	userdomain "github.com/char5742/ecsite-ddd-go/internal/user/domain"
)

var NewRegisterUserWorkflow RegisterUser = func(
	validateUser ValidateUser,
	registUser RegistUser,
) RegisterUserWorkflow {
	return func(cmd RegisterUserCommand) ([]RegisterUserEvent, shareerrs.DomainValidationResult, error) {
		validated, res := validateUser(cmd.Data.UnvalidatedUser, cmd.Data.ExternalUserData)
		if !res.IsComplete() {
			return nil, res, nil
		}
		registed, err := registUser(*validated)
		if err != nil {
			return nil, nil, nil
		}
		return []RegisterUserEvent{&UserRegistered{RegistedUser: *registed}}, nil, nil
	}
}

var ValidateUserImpl ValidateUser = func(uu userdomain.UnvalidatedUser, ext userdomain.ExternalUserData) (*userdomain.ValidatedUser, shareerrs.DomainValidationResult) {
	var result shareerrs.DomainValidationResult
	toValidateUser := userdomain.ToValidateUserImpl(ext)
	validated, res := toValidateUser(uu)
	result.Merge(res)

	return validated, result
}

var RegistUserImpl RegistUser = func(user userdomain.ValidatedUser) (*userdomain.RegistedUser, error) {

	return &userdomain.RegistedUser{}, nil
}
