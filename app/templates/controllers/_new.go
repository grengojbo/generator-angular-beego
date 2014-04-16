// Copyright 2014 <%= autorName %>. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package controllers

import (
	"github.com/astaxie/beego"
)

type <%= classedName %>Controller struct {
  baseController
}

func (this *<%= classedName %>Controller) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
  this.Data["Title"] = "<%= tTitle %>."
	this.TplNames = "<%= sname %>.tpl"

  // Setting properties.
  // this.Data["AppName"] = setting.AppName
  // this.Data["AppVer"] = setting.AppVer
  // this.Data["AppUrl"] = setting.AppUrl
  // this.Data["AppLogo"] = setting.AppLogo
}
