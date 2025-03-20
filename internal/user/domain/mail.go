package userdomain

func NewEmail(email string) Email {
	return Email{UniqueEmail{FormattedEmail{email}}}
}
