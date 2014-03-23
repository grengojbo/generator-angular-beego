// Copyright 2014 <%= autorName %>. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package main

import (
  "fmt"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/orm"
  "github.com/beego/i18n"
  _ "github.com/go-sql-driver/mysql"
  _ "<%= baseName %>/routers"
  "time"
)

func init() {
  orm.RegisterDriver("mysql", orm.DR_MySQL)
  // param 4 (optional):  set maximum idle connections
  // param 4 (optional):  set maximum connections (go >= 1.2)
  maxIdle := 30
  maxConn := 30
  conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", beego.AppConfig.String("mysqluser"), beego.AppConfig.String("mysqlpass"), beego.AppConfig.String("mysqlurls"), beego.AppConfig.String("mysqlport"), beego.AppConfig.String("mysqldb"))
  // orm.RegisterDataBase("default", "mysql", conn, maxIdle, maxConn)
  orm.RegisterDataBase("default", "mysql", conn, maxIdle, maxConn)
}

//    Objects
//  URL         HTTP Verb       Functionality
//  /object       POST          Creating Objects
//  /object/<objectId>  GET           Retrieving Objects
//  /object/<objectId>  PUT           Updating Objects
//  /object       GET           Queries
//  /object/<objectId>  DELETE          Deleting Objects

func main() {
  beego.SetStaticPath("/media", "media")
  if beego.AppConfig.String("runmode") == "dev" {
    beego.SetLevel(beego.LevelDebug)
    orm.Debug = true
  } else {
    beego.SetLevel(beego.LevelWarning)
  }
  beego.AddFuncMap("i18n", i18n.Tr)
  orm.DefaultTimeLoc = time.UTC
  orm.RunCommand()
  beego.Run()
}
