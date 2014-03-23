// Copyright 2014 <%= autorName %>. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package routers

import (
	"<%= baseName %>/controllers"
	"github.com/astaxie/beego"
)

func init(){
  beego.RESTRouter("/api/v1/example", &controllers.<%= modelName %>ServiceController{})
	beego.Router("/", &controllers.MainController{})
  // addController
}
