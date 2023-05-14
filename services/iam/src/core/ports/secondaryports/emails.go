package secondaryports

import (
	"iam/src/core/types"
)

type Email interface {
	WelcomeNewUser(email types.Email, username types.Username) error
	VerifyEmail(email types.Email, username types.Username, verificationCode types.Code) error
}
