package auth

type LoginDto struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
}

type SignupDto struct {
	Fullname string `form:"Fullname" `
	Email    string `form:"Email" validate:"email"`
	Password string `form:"Password" `
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
	Email string `json:"email" validate:"required,email"`
}

type ResetPasswordDto struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type VerifyEmailDto struct {
	Email string `param:"Email"`
	Token string `param:"Token"`
}
