package auth

import (
	"encoding/json"
	"main/internal/model"
	"main/pkg/engine/controller"
	"main/pkg/helpers"
	mailer "main/pkg/services/mail"
	"main/pkg/storage"
	view "main/server/app/v1/view/pages/auth"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func LoginPage(ctx *controller.Context) error {
	return ctx.Html(view.Login())
}

func SignupPage(ctx *controller.Context) error {
	return ctx.Html(view.Signup())
}

func ForgotPage(ctx *controller.Context) error {
	return ctx.Html(view.ForgotPassword())
}

func Login(ctx *controller.Context, Parameters *LoginDto) error {
	if Parameters.Email == "" || Parameters.Password == "" {
		return ctx.String(http.StatusBadRequest, "parameters are not provided :: Email")
	}

	var user model.Users
	if res := storage.DB.Where(&model.Users{
		Email: Parameters.Email,
	}).First(&user); res.Error != nil || res.RowsAffected < 1 {
		return ctx.String(http.StatusNotFound, "user not found")
	} else if !user.EmailVerified {
		return ctx.Html(view.VerifyEmail())
	}

	if wrpwd := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Parameters.Password)); wrpwd != nil {
		return ctx.Renders(http.StatusUnauthorized, view.Login())
	}

	Expires := time.Now().Add(24 * time.Hour)
	AuthHtmxMeta, _ := json.Marshal(struct {
		Expires time.Time `json:"Expires"`
		Email   string    `json:"Email"`
		Token   string    `json:"Token"`
	}{
		Expires: Expires,
		Email:   user.Email,
		Token:   user.Token,
	})

	ctx.WriteCookie(controller.Cookie{Key: "user", Value: user.Email, Expires: Expires})
	ctx.WriteCookie(controller.Cookie{Key: "token", Value: user.Token, Expires: Expires})
	ctx.Response().Header().Add("HX-AUTH-META", string(AuthHtmxMeta))
	ctx.Response().Header().Add("HX-Refresh", "true")

	return ctx.Html(view.Login())
}

func Signup(ctx *controller.Context, Parameters *SignupDto) error {
	if Parameters.TypeID != 2 && Parameters.TypeID != 3 {
		return ctx.String(http.StatusBadRequest, "parameters are not provided :: TypeID")
	} else if Parameters.Email == "" || Parameters.Password == "" || Parameters.Fullname == "" || Parameters.Phone == "" {
		return ctx.String(http.StatusBadRequest, "parameters are not provided :: Email")
	} else if res := storage.DB.Where(&model.Users{Email: Parameters.Email}).First(&model.Users{}); res.Error == nil || res.RowsAffected > 0 {
		return ctx.String(http.StatusConflict, "user already exists")
	}

	token := helpers.GenerateToken(155)
	if _, email := mailer.Send(mailer.Config{
		To:      Parameters.Email,
		Subject: "თქვენი მომხმარებელი გაიარეთ ანგარიში",
		Body:    "მომხმარებელი გაიარეთ ანგარიშის შესრულება მისამართით: https://alface.app/auth/verify/email/" + Parameters.Email + "/" + token,
	}); email != nil {
		return ctx.String(http.StatusInternalServerError, "mailer error")
	}

	Hash, err := bcrypt.GenerateFromPassword([]byte(Parameters.Password), 12)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bycrypt!!")
	}

	var Company *string
	if Parameters.TypeID == 3 && Parameters.Company != "" {
		Company = &Parameters.Company
	}

	if res := storage.DB.Create(&model.Users{
		Fullname: Parameters.Fullname,
		Email:    Parameters.Email,
		Phone:    Parameters.Phone,
		TypeID:   Parameters.TypeID,
		Password: string(Hash),
		Company:  Company,
		Token:    token,
	}); res.Error != nil {
		return ctx.String(http.StatusInternalServerError, "database error ->"+res.Error.Error())
	}

	return ctx.Html(view.VerifyEmail())
}

func VerifyEmail(ctx *controller.Context, Parameters *VerifyEmailTokenDto) error {
	if _, welcome := mailer.Send(mailer.Config{
		To:      Parameters.Email,
		Subject: "Alfashop",
		Body:    "მოგესალმებით, თქვენ წარმატებით გაიარეთ ვერიფიკაცია",
	}); welcome != nil {
		return ctx.String(http.StatusInternalServerError, "mailer error")
	}

	if verify := storage.DB.Model(&model.Users{}).
		Where(&model.Users{
			Email: Parameters.Email,
			Token: Parameters.Token,
		}).
		Updates(model.Users{
			EmailVerified:   true,
			EmailVerifyedAt: time.Now(),
		}); verify.Error != nil || verify.RowsAffected < 1 {
		return ctx.String(http.StatusInternalServerError, "database error")
	}

	return ctx.Redirect(http.StatusMovedPermanently, "/auth/login")
}

func ForgotPassword(ctx *controller.Context, Parameters *ForgotPasswordDto) error {
	if Parameters.Email == "" {
		return ctx.String(http.StatusBadRequest, "parameters are not provided :: Email")
	}

	var user model.Users
	res := storage.DB.Where(&model.Users{Email: Parameters.Email}).First(&user)
	if res.Error != nil || res.RowsAffected < 1 {
		return ctx.String(http.StatusNotFound, "user not found")
	}

	Token := helpers.GenerateToken(155)
	user.ResetToken = helpers.ToPointer(Token)
	user.ResetTokenExpiresAt = time.Now().Add(30 * time.Minute)
	if res := storage.DB.Save(&user); res.Error != nil {
		return ctx.String(http.StatusInternalServerError, "database error -> "+res.Error.Error())
	}

	if _, forgotEmail := mailer.Send(mailer.Config{
		To:      Parameters.Email,
		Subject: "Alfashop",
		Body:    "მომხმარებელი გაიარეთ ანგარიშის შესრულება მისამართით: https://alface.app/auth/verify/password/" + Parameters.Email + "/" + Token,
	}); forgotEmail != nil {
		return ctx.String(http.StatusInternalServerError, "mailer error")
	}

	return ctx.Html(view.ForgotPasswordSuccess())
}

func VerifyPassword(ctx *controller.Context, Parameters *VerifyEmailTokenDto) error {
	if Parameters.Email == "" || Parameters.Token == "" {
		return ctx.String(http.StatusBadRequest, "parameters are not provided :: Email || Token")
	}

	var user model.Users
	if res := storage.DB.Where(&model.Users{
		Email:      Parameters.Email,
		ResetToken: &Parameters.Token,
	}).First(&user); res.Error != nil || res.RowsAffected < 1 {
		return ctx.String(http.StatusNotFound, "session not found")
	} else if user.ResetTokenExpiresAt.Before(time.Now()) {
		return ctx.String(http.StatusNotFound, "session expired")
	}

	return ctx.Html(view.ChangePassword(user.Email, *user.ResetToken))
}

func ChangePassword(ctx *controller.Context, Parameters *ChangePasswordDto) error {
	if Parameters.Email == "" || Parameters.Password == "" || Parameters.Token == "" {
		return ctx.String(
			http.StatusBadRequest,
			"parameters are not provided :: Email || Password || Token",
		)
	}

	var user model.Users
	res := storage.DB.Where(&model.Users{
		Email:      Parameters.Email,
		ResetToken: &Parameters.Token,
	}).First(&user)

	if res.Error != nil || res.RowsAffected < 1 {
		return ctx.String(http.StatusNotFound, "user not found")
	} else if user.ResetTokenExpiresAt.Before(time.Now()) {
		return ctx.String(http.StatusNotFound, "session expired")
	}

	Hash, err := bcrypt.GenerateFromPassword([]byte(Parameters.Password), 12)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bycrypt!!")
	}

	user.Password = string(Hash)
	if res := storage.DB.Save(&user); res.Error != nil {
		return ctx.String(http.StatusInternalServerError, "database error -> "+res.Error.Error())
	}

	return ctx.Redirect(http.StatusMovedPermanently, "/auth/login")
}

func Logout(ctx *controller.Context) error {
	ctx.RemoveCookie("user")
	ctx.RemoveCookie("token")
	return ctx.Html(view.Logout())
}
