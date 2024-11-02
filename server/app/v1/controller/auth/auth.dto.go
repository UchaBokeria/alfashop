package auth

type LoginDto struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type SignupDto struct {
	Fullname string `form:"Fullname" `
	Email    string `form:"Email" validate:"email"`
	Password string `form:"Password" validate:"strongpwd"`
	Phone    string `form:"Phone" `
	Company  string `form:"Company" `
	// CityID   int    `form:"City" `
	TypeID int `form:"TypeID" `
}

type ChangePasswordDto struct {
	Email    string `form:"Email"`
	Token    string `form:"Token"`
	Password string `form:"Password"`
}

type ForgotPasswordDto struct {
	Email string `form:"email"`
}

type ResetPasswordDto struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type VerifyEmailTokenDto struct {
	Email string `param:"Email"`
	Token string `param:"Token"`
}
