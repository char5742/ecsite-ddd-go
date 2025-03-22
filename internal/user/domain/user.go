package userdomain

import shareutils "github.com/char5742/ecsite-ddd-go/internal/share/utils"

var ToValidateUserImpl ToValidateUser = func(ext ExternalUserData) func(UnvalidatedUser) (*ValidatedUser, shareutils.DomainValidationResult) {
	validateEmail := toValidateEmailImpl(
		toFormattedEmailImpl,
		toUniqueEmailImpl,
		ext.ExternalEmailData,
	)
	return func(user UnvalidatedUser) (*ValidatedUser, shareutils.DomainValidationResult) {
		var result shareutils.DomainValidationResult
		email, res := validateEmail(user.Email)
		result = result.Merge(res)

		if (result != nil) && !result.IsComplete() {
			return nil, result
		}

		return &ValidatedUser{
			FirstName:      FirstName(user.FirstName),
			LastName:       LastName(user.LastName),
			Email:          email,
			Password:       Password(user.Password),
			Zipcode:        Zipcode(user.Zipcode),
			Prefecture:     Prefecture(user.Prefecture),
			Municipalities: Municipalities(user.Municipalities),
			Address:        Address(user.Address),
			Telephone:      Telephone(user.Telephone),
		}, result
	}

}
