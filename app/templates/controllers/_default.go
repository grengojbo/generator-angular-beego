// Copyright 2014 <%= autorName %>. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package controllers

import (
  "fmt"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/cache"
  "github.com/astaxie/beego/validation"
  "github.com/beego/i18n"
  "html/template"
  "<%= sname %>/models"
  "github.com/grengojbo/beego/modules/utils"
  "net/url"
  "reflect"
  "strconv"
  "strings"
  "time"
)

var langTypes []string // Languages that are supported.

func init() {
  // Initialize language type list.
  langTypes = strings.Split(beego.AppConfig.String("lang_types"), "|")

  // Load locale files according to language types.
  for _, lang := range langTypes {
    beego.Debug("Loading language: " + lang)
    if err := i18n.SetMessage(lang, "conf/"+"locale_"+lang+".ini"); err != nil {
      beego.Error("Fail to set message file:", err)
      return
    }
  }
}

// baseController represents base router for all other app routers.
// It implemented some methods for the same implementation;
// thus, it will be embedded into other routers.
type baseController struct {
  beego.Controller // Embed struct that has stub implementation of the interface.
  i18n.Locale      // For i18n usage when process data and render template.
  isLogin          bool
  userId           int
  user             models.User
  Cache            cache.Cache
}

// Prepare implemented Prepare() method for baseController.
// It's used for language option check and setting.
func (this *baseController) Prepare() {
  this.Data["PageStartTime"] = time.Now()
  c, err := cache.NewCache("memory", `{"interval":60}`)
  if err != nil {
    beego.Error("No Cache start")
  } else {
    this.Cache = c
  }
  sId := this.GetSession("_auth_user_id")
  djangoAuthModule := this.GetSession("_auth_user_backend")
  if sId == nil {
    this.userId = int(-1)
    this.SetSession("_auth_user_id", this.userId)
    sId = "-1"
  } else {
    beego.Debug("session ID:", sId)
    if str, ok := sId.(string); ok {
      this.userId, _ = strconv.Atoi(str)
    } else {
      this.userId = sId.(int)
      // this.userId, _ = strconv.Atoi(sId.(string))
    }
  }
  u, err := models.GetUser(this.userId)
  if err == nil {
    this.user = u
    this.isLogin = true
    beego.Debug("is Login:", this.user.Username)
  }
  beego.Debug("Session User: ", this.userId)
  if djangoAuthModule == nil {
    this.SetSession("_auth_user_backend", "django.contrib.auth.backends.ModelBackend")
  }
  // Reset language option.
  this.Lang = "" // This field is from i18n.Locale.

  // 1. Get language information from 'Accept-Language'.
  al := this.Ctx.Request.Header.Get("Accept-Language")
  if len(al) > 4 {
    al = al[:5] // Only compare first 5 letters.
    if i18n.IsExist(al) {
      this.Lang = al
    }
  }

  // 2. Default language is English.
  if len(this.Lang) == 0 {
    this.Lang = beego.AppConfig.String("lang_default")
  }

  // Set template level language option.
  this.Data["Lang"] = this.Lang
}

func (this *baseController) setLangCookie(lang string) {
  this.Ctx.SetCookie("lang", lang, 60*60*24*365, "/", nil, nil, false)
}

// check if not login then redirect
func (this *baseController) CheckLoginRedirect(args ...interface{}) bool {
  var redirect_to string
  code := 302
  needLogin := true
  for _, arg := range args {
    switch v := arg.(type) {
    case bool:
      needLogin = v
    case string:
      // custom redirect url
      redirect_to = v
    case int:
      // custom redirect url
      code = v
    }
  }

  // if need login then redirect
  if needLogin && !this.isLogin {
    if len(redirect_to) == 0 {
      req := this.Ctx.Request
      scheme := "http"
      if req.TLS != nil {
        scheme += "s"
      }
      redirect_to = fmt.Sprintf("%s://%s%s", scheme, req.Host, req.RequestURI)
    }
    redirect_to = "/login?to=" + url.QueryEscape(redirect_to)
    this.Redirect(redirect_to, code)
    return true
  }

  // if not need login then redirect
  if !needLogin && this.isLogin {
    if len(redirect_to) == 0 {
      redirect_to = "/"
    }
    this.Redirect(redirect_to, code)
    return true
  }
  return false
}

// read beego flash message
func (this *baseController) FlashRead(key string) (string, bool) {
  if data, ok := this.Data["flash"].(map[string]string); ok {
    value, ok := data[key]
    return value, ok
  }
  return "", false
}

// write beego flash message
func (this *baseController) FlashWrite(key string, value string) {
  flash := beego.NewFlash()
  flash.Data[key] = value
  flash.Store(&this.Controller)
}

func (this *baseController) validForm(form interface{}, names ...string) (bool, map[string]*validation.ValidationError) {
  // parse request params to form ptr struct
  utils.ParseForm(form, this.Input())

  // Put data back in case users input invalid data for any section.
  name := reflect.ValueOf(form).Elem().Type().Name()
  if len(names) > 0 {
    name = names[0]
  }
  this.Data[name] = form

  errName := name + "Error"

  // check form once
  if this.FormOnceNotMatch() {
    return false, nil
  }

  // Verify basic input.
  valid := validation.Validation{}
  if ok, _ := valid.Valid(form); !ok {
    errs := valid.ErrorMap()
    this.Data[errName] = &valid
    return false, errs
  }
  return true, nil
}

// check form once, void re-submit
func (this *baseController) FormOnceNotMatch() bool {
  notMatch := false
  recreat := false

  // get token from request param / header
  var value string
  if vus, ok := this.Input()["_once"]; ok && len(vus) > 0 {
    value = vus[0]
  } else {
    value = this.Ctx.Input.Header("X-Form-Once")
  }

  // exist in session
  if v, ok := this.GetSession("form_once").(string); ok && v != "" {
    // not match
    if value != v {
      notMatch = true
    } else {
      // if matched then re-creat once
      recreat = true
    }
  }

  this.FormOnceCreate(recreat)
  return notMatch
}

// create form once html
func (this *baseController) FormOnceCreate(args ...bool) {
  var value string
  var creat bool
  creat = len(args) > 0 && args[0]
  if !creat {
    if v, ok := this.GetSession("form_once").(string); ok && v != "" {
      value = v
    } else {
      creat = true
    }
  }
  if creat {
    value = utils.GetRandomString(10)
    this.SetSession("form_once", value)
  }
  this.Data["once_token"] = value
  this.Data["once_html"] = template.HTML(`<input type="hidden" name="_once" value="` + value + `">`)
}

// valid form and put errors to tempalte context
func (this *baseController) ValidForm(form interface{}, names ...string) bool {
  valid, _ := this.validForm(form, names...)
  return valid
}

// valid form and put errors to tempalte context
func (this *baseController) ValidFormSets(form interface{}, names ...string) bool {
  valid, errs := this.validForm(form, names...)
  this.setFormSets(form, errs, names...)
  return valid
}

func (this *baseController) SetFormSets(form interface{}, names ...string) *utils.FormSets {
  return this.setFormSets(form, nil, names...)
}

func (this *baseController) setFormSets(form interface{}, errs map[string]*validation.ValidationError, names ...string) *utils.FormSets {
  formSets := utils.NewFormSets(form, errs, this.Locale)
  name := reflect.ValueOf(form).Elem().Type().Name()
  if len(names) > 0 {
    name = names[0]
  }
  name += "Sets"
  this.Data[name] = formSets

  return formSets
}

func (this *baseController) TimesReachedTest(key string, times int) (int, bool) {
  var retries int
  if v := this.Cache.Get(key); v != nil {
    if d, ok := v.(int); ok {
      if d > times {
        return d, true
      }
      retries = d
    }
  }
  return retries, false
}

func (this *baseController) TimesReachedSet(key string, times int, reloadMinutes int) {
  this.Cache.Put(key, times+1, int64(reloadMinutes)*60)
}
