// Copyright 2014 Oleg Dolya. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package controllers

type MainController struct {
	baseController
	// beego.Controller
}

func (this *MainController) Get() {
	this.Data["Website"] = "beego.me"
	this.Data["AppUrl"] = '/'
	this.Data["IsHome"] = true
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
