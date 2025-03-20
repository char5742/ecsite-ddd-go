package userdomain

import (
	"errors"

	shareutils "github.com/char5742/ecsite-ddd-go/internal/share/utils"
)

// 既にメールアドレスが使用されているか
type IsEmailTaken func(email FormattedEmail) (bool, error)

func NewValidatedUser(
	taken IsEmailTaken,
	user SelfValidatedUser,
) (*ValidatedUser, error) {
	// メールアドレスの重複チェック
	isTaken, err := taken(user.FormattedEmail)
	if err != nil {
		return nil, err
	}
	if isTaken {
		return nil, errors.New("このメールアドレスは既に使用されています")
	}
	return &ValidatedUser{
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		UniqueEmail:    UniqueEmail{user.FormattedEmail},
		Password:       user.Password,
		Zipcode:        user.Zipcode,
		Prefecture:     user.Prefecture,
		Municipalities: user.Municipalities,
		Address:        user.Address,
		Telephone:      user.Telephone,
	}, nil
}

// Userのドメインロジック

func NewSelfValidateUser(
	user UnvalidatedUser,
) (*SelfValidatedUser, error) {
	var errs shareutils.ValidationErrors

	firstName := user.FirstName
	lastName := user.LastName
	email := user.Email
	password := user.Password
	zipcode := user.Zipcode
	prefecture := user.Prefecture
	municipalities := user.Municipalities
	address := user.Address
	telephone := user.Telephone

	// ユーザー名のバリデーション
	validFirstName, err := validateFirstName(firstName)
	if err != nil {
		errs = errs.Add("firstName", err)
	}

	validLastName, err := validateLastName(lastName)
	if err != nil {
		errs = errs.Add("lastName", err)
	}

	validEmail, err := validateEmail(email)
	if err != nil {
		errs = errs.Add("email", err)
	}

	validPassword, err := validatePassword(password)
	if err != nil {
		errs = errs.Add("password", err)
	}

	validZipcode, err := validateZipcode(zipcode)
	if err != nil {
		errs = errs.Add("zipcode", err)
	}

	validPrefecture, err := validatePrefecture(prefecture)
	if err != nil {
		errs = errs.Add("prefecture", err)
	}

	validMunicipalities, err := validateMunicipalities(municipalities)
	if err != nil {
		errs = errs.Add("municipalities", err)
	}

	validAddress, err := validateAddress(address)
	if err != nil {
		errs = errs.Add("address", err)
	}

	validTelephone, err := validateTelephone(telephone)
	if err != nil {
		errs = errs.Add("telephone", err)
	}

	if !errs.IsEmpty() {
		return nil, errs
	}

	// ユーザーを生成する
	generatedUser := &SelfValidatedUser{
		FirstName:      *validFirstName,
		LastName:       *validLastName,
		FormattedEmail: *validEmail,
		Password:       *validPassword,
		Zipcode:        *validZipcode,
		Prefecture:     *validPrefecture,
		Municipalities: *validMunicipalities,
		Address:        *validAddress,
		Telephone:      *validTelephone,
	}

	return generatedUser, nil
}

func validateEmail(email string) (*FormattedEmail, error) {
	if email == "" {
		return nil, errors.New("メールアドレスを入力してください")
	}
	formatted := FormattedEmail{email}
	return &formatted, nil
}

func validatePassword(password string) (*Password, error) {
	if password == "" {
		return nil, errors.New("パスワードを入力してください")
	}

	if len(password) < 8 {
		return nil, errors.New("パスワードは8文字以上で入力してください")
	}

	formatted := Password(password)
	return &formatted, nil
}

func validateZipcode(zipcode string) (*Zipcode, error) {
	if zipcode == "" {
		return nil, errors.New("郵便番号を入力してください")
	}

	if len(zipcode) != 7 {
		return nil, errors.New("郵便番号は7桁で入力してください")
	}

	formatted := Zipcode(zipcode)
	return &formatted, nil
}

func validatePrefecture(prefecture string) (*Prefecture, error) {
	if prefecture == "" {
		return nil, errors.New("都道府県を入力してください")
	}

	formatted := Prefecture(prefecture)
	return &formatted, nil
}

func validateMunicipalities(municipalities string) (*Municipalities, error) {
	if municipalities == "" {
		return nil, errors.New("市区町村を入力してください")
	}

	formatted := Municipalities(municipalities)
	return &formatted, nil
}

func validateAddress(address string) (*Address, error) {
	if address == "" {
		return nil, errors.New("住所を入力してください")
	}

	formatted := Address(address)
	return &formatted, nil
}

func validateTelephone(telephone string) (*Telephone, error) {
	if telephone == "" {
		return nil, errors.New("電話番号を入力してください")
	}

	if len(telephone) < 10 || len(telephone) > 15 {
		return nil, errors.New("電話番号は10〜15桁で入力してください")
	}

	formatted := Telephone(telephone)
	return &formatted, nil
}

func validateFirstName(firstName string) (*FirstName, error) {
	if firstName == "" {
		return nil, errors.New("名前を入力してください")
	}

	formatted := FirstName(firstName)
	return &formatted, nil
}

func validateLastName(lastName string) (*LastName, error) {
	if lastName == "" {
		return nil, errors.New("姓を入力してください")
	}

	formatted := LastName(lastName)
	return &formatted, nil
}
