package auth

type Authenticator interface {
	Authenticate() bool
	UserID() uint
	IsAdmin() bool
}
