package userdomain

// メールアドレス
type Email struct{ UniqueEmail }

// 適切な形式のメールアドレス
type FormattedEmail struct{ string }

// 重複していないメールアドレス
type UniqueEmail struct{ FormattedEmail }

// メールアドレスを適切な形式であると保証する関数
type CheckEmailFormat func(string) (FormattedEmail, error)

// メールアドレスの重複を検証する関数
type CheckEmailUniqueness func(string) (UniqueEmail, error)

type ValidateEmail func(CheckEmailFormat, CheckEmailUniqueness) func(string) (Email, error)
