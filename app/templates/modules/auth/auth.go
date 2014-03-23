// Copyright 2014 Oleg Dolya. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package auth

import (
	// "encoding/hex"
	"fmt"
	"github.com/astaxie/beego/context"
	"strings"
	// "time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/session"

	"github.com/grengojbo/beego/modules/utils"
	// "github.com/beego/wetalk/setting"
	"myapp/models"
)

// check if exist user by username or email
func HasUser(user *models.User, username string) bool {
	var err error
	qs := orm.NewOrm()
	if strings.IndexRune(username, '@') == -1 {
		user.Username = username
		err = qs.Read(user, "Username")
	} else {
		user.Email = username
		err = qs.Read(user, "Email")
	}
	if err == nil {
		return true
	}
	return false
}

// set a new password to user
func SaveNewPassword(user *models.User, password string) error {
	salt := models.GetUserSalt()
	user.Password = fmt.Sprintf("%s$%s", salt, utils.EncodePassword(password, salt))
	return user.Update("Password", "Rands", "Updated")
}

// get login redirect url from cookie
func GetLoginRedirect(ctx *context.Context) string {
	loginRedirect := strings.TrimSpace(ctx.GetCookie("login_to"))
	if utils.IsMatchHost(loginRedirect) == false {
		loginRedirect = "/"
	} else {
		ctx.SetCookie("login_to", "", -1, "/")
	}
	return loginRedirect
}

// login user
func LoginUser(user *models.User, ctx *context.Context, remember bool) {
	// werid way of beego session regenerate id...
	ctx.Input.CruSession.SessionRelease(ctx.ResponseWriter)
	ctx.Input.CruSession = beego.GlobalSessions.SessionRegenerateId(ctx.ResponseWriter, ctx.Request)
	ctx.Input.CruSession.Set("_auth_user_id", user.Id)

	if remember {
		WriteRememberCookie(user, ctx)
	}
}

func WriteRememberCookie(user *models.User, ctx *context.Context) {
	secret := utils.EncodeMd5(beego.AppConfig.String("secret") + user.Password)
	lrd, _ := beego.AppConfig.Int("LoginRememberDays")
	days := 86400 * lrd
	ctx.SetCookie(beego.AppConfig.String("CookieUserName"), user.Username, days)
	ctx.SetSecureCookie(secret, beego.AppConfig.String("CookieRememberName"), user.Username, days)
}

func DeleteRememberCookie(ctx *context.Context) {
	ctx.SetCookie(beego.AppConfig.String("CookieUserName"), "", -1)
	ctx.SetCookie(beego.AppConfig.String("CookieRememberName"), "", -1)
}

func LoginUserFromRememberCookie(user *models.User, ctx *context.Context) (success bool) {
	userName := ctx.GetCookie(beego.AppConfig.String("CookieUserName"))
	if len(userName) == 0 {
		return false
	}

	defer func() {
		if !success {
			DeleteRememberCookie(ctx)
		}
	}()

	user.Username = userName
	if err := user.Read("UserName"); err != nil {
		return false
	}

	secret := utils.EncodeMd5(beego.AppConfig.String("secret") + user.Password)
	value, _ := ctx.GetSecureCookie(secret, beego.AppConfig.String("CookieRememberName"))
	if value != userName {
		return false
	}

	LoginUser(user, ctx, true)

	return true
}

// logout user
func LogoutUser(ctx *context.Context) {
	DeleteRememberCookie(ctx)
	ctx.Input.CruSession.Delete("_auth_user_id")
	ctx.Input.CruSession.Flush()
	beego.GlobalSessions.SessionDestroy(ctx.ResponseWriter, ctx.Request)
}

func GetUserIdFromSession(sess session.SessionStore) int {
	if id, ok := sess.Get("_auth_user_id").(int); ok && id > 0 {
		return id
	}
	return 0
}

// get user if key exist in session
func GetUserFromSession(user *models.User, sess session.SessionStore) bool {
	id := GetUserIdFromSession(sess)
	if id > 0 {
		u := models.User{Id: id}
		if u.Read() == nil {
			*user = u
			return true
		}
	}

	return false
}

// verify username/email and password
func VerifyUser(user *models.User, username, password string) (success bool) {
	// search user by username or email
	if HasUser(user, username) == false {
		return
	}

	if VerifyPassword(password, user.Password) {
		// success
		success = true

		// re-save discuz password
		if len(user.Password) == 39 {
			if err := SaveNewPassword(user, password); err != nil {
				beego.Error("SaveNewPassword err: ", err.Error())
			}
		}
	}
	return
}

// compare raw password and encoded password
func VerifyPassword(rawPwd, encodedPwd string) bool {

	// for discuz accounts
	if len(encodedPwd) == 39 {
		salt := encodedPwd[:6]
		encoded := encodedPwd[7:]
		return encoded == utils.EncodeMd5(utils.EncodeMd5(rawPwd)+salt)
	}

	// split
	var salt, encoded string
	if len(encodedPwd) > 11 {
		salt = encodedPwd[:10]
		encoded = encodedPwd[11:]
	}

	return utils.EncodePassword(rawPwd, salt) == encoded
}
