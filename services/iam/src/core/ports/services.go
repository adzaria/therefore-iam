package ports

type UsersService interface {
	Register(args RegisterArgs) (RegisterAnswer, error)
	// WhoAmI() error
}

type EmailVerificationService interface {
	Send() error
	Verify() error
}

type AuthenticationService interface {
	MagicLink() error
	Token() error
	RefreshToken() error
	RevokeToken() error
	VerifyToken() error
}