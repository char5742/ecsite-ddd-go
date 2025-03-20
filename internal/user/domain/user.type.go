package userdomain

import (
	identitytypes "github.com/char5742/ecsite-ddd-go/internal/identity/domain/types"
	sharetypes "github.com/char5742/ecsite-ddd-go/internal/share/domain/types"
)

// 性
type FirstName string

// 名
type LastName string

// パスワード
type Password string

// 郵便番号
type Zipcode string

// 都道府県
type Prefecture string

// 市区町村
type Municipalities string

// 住所
type Address string

// 電話番号
type Telephone string

// ユーザー
type User struct {
	ID             identitytypes.IdentityID `json:"id"`
	FirstName      FirstName                `json:"first_name"`
	LastName       LastName                 `json:"last_name"`
	Email          Email                    `json:"email"`
	Password       Password                 `json:"password"`
	Zipcode        Zipcode                  `json:"zipcode"`
	Prefecture     Prefecture               `json:"prefecture"`
	Municipalities Municipalities           `json:"municipalities"`
	Address        Address                  `json:"address"`
	Telephone      Telephone                `json:"telephone"`

	AuditInfo sharetypes.AuditInfo `json:"audit_info"`
}

// 検証されていないユーザー
type UnvalidatedUser struct {
	// ユーザーの一意識別子
	ID string `json:"id"`
	// ユーザーの性
	FirstName string `json:"first_name"`
	// ユーザーの名
	LastName string `json:"last_name"`
	// ユーザーのメールアドレス
	Email string `json:"email"`
	// ユーザーのパスワード
	Password string `json:"password"`
	// ユーザーの郵便番号
	Zipcode string `json:"zipcode"`
	// ユーザーの都道府県
	Prefecture string `json:"prefecture"`
	// ユーザーの市区町村
	Municipalities string `json:"municipalities"`
	// ユーザーの住所
	Address string `json:"address"`
	// ユーザーの電話番号
	Telephone string `json:"telephone"`
}

// 一次検証済みユーザー
type SelfValidatedUser struct {
	ID identitytypes.IdentityID
	FirstName
	LastName
	FormattedEmail
	Password
	Zipcode
	Prefecture
	Municipalities
	Address
	Telephone
}

// 検証済みユーザー
type ValidatedUser struct {
	ID identitytypes.IdentityID
	FirstName
	LastName
	UniqueEmail
	Password
	Zipcode
	Prefecture
	Municipalities
	Address
	Telephone
}

// 登録されたユーザー
type RegistedUser struct {
	ID identitytypes.IdentityID
	FirstName
	LastName
	Email
	Password
	Zipcode
	Prefecture
	Municipalities
	Address
	Telephone
}
