package auth

import (
	"encoding/json"
	"main/internal/model"
	"main/pkg/engine/controller"
	"main/pkg/storage"
	view "main/server/admin/view/pages"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func login(ctx *controller.Context, Parameters *LoginDto) error {
	var UserMatch model.Users

	if Parameters.Email == "" || Parameters.Password == "" {
		return ctx.Renders(http.StatusBadRequest, view.Login())
	}

	if Result := storage.DB.Last(&UserMatch, &model.Users{Email: Parameters.Email}); Result.Error != nil || Result.RowsAffected < 1 {
		return ctx.Renders(http.StatusNotFound, view.Login())
	}

	err := bcrypt.CompareHashAndPassword([]byte(UserMatch.Password), []byte(Parameters.Password))
	if err != nil {
		return ctx.Renders(http.StatusUnauthorized, view.Login())
	}

	Expires := time.Now().Add(24 * time.Hour)

	AuthHtmxMeta, _ := json.Marshal(struct {
		Expires time.Time `json:"Expires"`
		Email   string    `json:"Email"`
		Token   string    `json:"Token"`
	}{
		Expires: Expires,
		Email:   UserMatch.Email,
		Token:   UserMatch.Token,
	})

	ctx.WriteCookie(controller.Cookie{Key: "user", Value: UserMatch.Email, Expires: Expires})
	ctx.WriteCookie(controller.Cookie{Key: "token", Value: UserMatch.Token, Expires: Expires})
	ctx.Response().Header().Add("HX-AUTH-META", string(AuthHtmxMeta))

	return ctx.Html(view.Login())
}

func ChangePassword(ctx *controller.Context, Parameters *LoginDto) error {
	if Parameters.Email == "" || Parameters.Password == "" {
		return ctx.String(http.StatusBadRequest, "atrakeb")
	}

	Hash, err := bcrypt.GenerateFromPassword([]byte(Parameters.Password), 12)
	if err != nil {
		return ctx.String(http.StatusBadRequest, "bycrypt!!")
	}

	Result := storage.DB.Table("users").Where(&model.Users{Email: Parameters.Email}).Update("password", string(Hash))
	if Result.Error != nil || Result.RowsAffected < 1 {
		return ctx.String(http.StatusBadRequest, "No rows affected")
	}
	return ctx.String(http.StatusOK, "success")
}
