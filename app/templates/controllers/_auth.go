// Copyright 2014 <%= autorName %>. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	"<%= sname %>/models"
	"<%= _sname %>/modules/auth"
	"github.com/grengojbo/beego/modules/utils"
	"strings"
)

type LoginController struct {
	baseController
}

// Get implemented login page.
func (this *LoginController) Get() {
	this.Data["IsLoginPage"] = true
	this.TplNames = "auth/login.html"

	loginRedirect := strings.TrimSpace(this.GetString("to"))
	if utils.IsMatchHost(loginRedirect) == false {
		loginRedirect = "/"
	}

	// no need login
	if this.CheckLoginRedirect(false, loginRedirect) {
		return
	}

	if len(loginRedirect) > 0 {
		this.Ctx.SetCookie("login_to", loginRedirect, 0, "/")
	}

	form := auth.LoginForm{}
	this.SetFormSets(&form)
}

// Login implemented user login.
func (this *LoginController) Login() {
	this.Data["IsLoginPage"] = true
	this.TplNames = "auth/login.html"
	lmri, _ := beego.AppConfig.Int("LoginMaxRetries")

	// no need login
	if this.CheckLoginRedirect(false) {
		return
	}

	var user models.User
	var key string
	ajaxErrMsg := "auth.login_error_ajax"

	form := auth.LoginForm{}
	// valid form and put errors to template context
	if this.ValidFormSets(&form) == false {
		if this.IsAjax() {
			goto ajaxError
		}
		return
	}

	key = "auth.login." + form.UserName + this.Ctx.Input.IP()
	if times, ok := this.TimesReachedTest(key, lmri); ok {
		if this.IsAjax() {
			ajaxErrMsg = "auth.login_error_times_reached"
			goto ajaxError
		}
		this.Data["ErrorReached"] = true

	} else if auth.VerifyUser(&user, form.UserName, form.Password) {
		loginRedirect := this.LoginUser(&user, form.Remember)

		if this.IsAjax() {
			this.Data["json"] = map[string]interface{}{
				"success":  true,
				"message":  this.Tr("auth.login_success_ajax"),
				"redirect": loginRedirect,
			}
			this.ServeJson()
			return
		}

		this.Redirect(loginRedirect, 302)
		return
	} else {
		lfb, _ := beego.AppConfig.Int("LoginFailedBlocks")
		this.TimesReachedSet(key, times, lfb)
		if this.IsAjax() {
			goto ajaxError
		}
	}
	this.Data["Error"] = true
	return

ajaxError:
	this.Data["json"] = map[string]interface{}{
		"success": false,
		"message": this.Tr(ajaxErrMsg),
		"once":    this.Data["once_token"],
	}
	this.ServeJson()
}

// Logout implemented user logout page.
func (this *LoginController) Logout() {
	auth.LogoutUser(this.Ctx)

	// write flash message
	this.FlashWrite("HasLogout", "true")

	this.Redirect("/login", 302)
}

func (this *LoginController) LoginUser(user *models.User, remember bool) string {
	loginRedirect := strings.TrimSpace(this.Ctx.GetCookie("login_to"))
	if utils.IsMatchHost(loginRedirect) == false {
		loginRedirect = "/"
	} else {
		this.Ctx.SetCookie("login_to", "", -1, "/")
	}

	// login user
	auth.LoginUser(user, this.Ctx, remember)
	userLang := user.Lang
	if userLang > 0 {
		this.setLangCookie(i18n.GetLangByIndex(userLang))
	}

	return loginRedirect
}
