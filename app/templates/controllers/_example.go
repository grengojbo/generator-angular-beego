// Copyright 2014 <%= autorName %>. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package controllers

import (
  "github.com/astaxie/beego"
  "<%= _.slugify(baseName) %>/models"
  "strconv"
)

type <%= classedName %>Controller struct {
  baseController
}

// <%= classedName %>Controller metod Get
// GET http://localhost:8080/api/v3/<%= sname %> # All result
// GET http://localhost:8080/api/v3/<%= sname %>/<objectId> # One result
func (t *<%= classedName %>Controller) Get() {
  // page, _ := t.GetInt("page")
  // page_size, _ := t.GetInt("rows")
  sort := t.GetString("sort")
  order := t.GetString("order")
  if len(order) > 0 {
    beego.Debug("Order: ", order)
    if order == "desc" {
      sort = "-" + sort
    }
  } else {
    sort = "-created"
  }
  beego.Debug("API LiqPayController metod GET isLogin: ", t.isLogin)
  objectId := t.Ctx.Input.Params[":objectId"]
  // if t.isLogin {
  if len(objectId) > 0 {
    id, err := strconv.ParseInt(objectId, 10, 64)
    if err != nil {
      t.Ctx.Output.SetStatus(400)
      t.Ctx.Output.Body([]byte(err.Error()))
      return
    } else {
      ob, err := models.Get<%= classedName %>(id)
      if err != nil {
        mess := fmt.Sprintf("Is not row ID: %d", id)
        t.Data["json"] = &map[string]interface{}{"error": true, "message": mess}
      } else {
        t.Data["json"] = ob
      }
    }
  } else {
    ob, cnt, err := models.GetAll<%= classedName %>(sort)
    if err != nil {
      t.Ctx.Output.SetStatus(400)
      t.Ctx.Output.Body([]byte(err.Error()))
      return
    }
    t.Data["json"] = &map[string]interface{}{"count": cnt, "next": nil, "previous": nil, "results": &ob}
  }
  // } else {
  //  mes := "Is not Auntification :("
  //  t.Data["json"] = &map[string]interface{}{"error": true, "message": mes}
  // }
  t.ServeJson()
}

// <%= classedName %>Controller metod Post
// POST http://localhost:8080/api/v1/<%= sname %>/
func (this *<%= classedName %>Controller) Post() {
  var ob models.Object
  json.Unmarshal(this.Ctx.Input.RequestBody, &ob)
  objectid := models.AddOne(ob)
  this.Data["json"] = map[string]string{"ObjectId": objectid}
  this.ServeJson()
}

// <%= classedName %>Controller metod Put
// PUT http://localhost:8080/api/v1/<%= sname %>/
func (this *<%= classedName %>Controller) Put() {
  objectId := this.Ctx.Input.Params[":objectId"]
  var ob models.Object
  json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

  err := models.Update(objectId, ob.Score)
  if err != nil {
    this.Data["json"] = err
  } else {
    this.Data["json"] = "update success!"
  }
  this.ServeJson()
}

// <%= classedName %>Controller metod Delete
// API for delete object
// DELETE http://localhost:8080/api/v1/<%= sname %>/<objectId>
func (t *<%= classedName %>Controller) Delete() {
  if t.isLogin {
    id, err := strconv.ParseInt(t.Ctx.Input.Params[":objectId"], 10, 64)
    if err != nil {
      t.Ctx.Output.SetStatus(400)
      t.Ctx.Output.Body([]byte(err.Error()))
      return
    }
    err = models.Delete<%= classedName %>(id)
    if err != nil {
     t.Ctx.Output.SetStatus(400)
      t.Ctx.Output.Body([]byte(err.Error()))
      return
    }
    mes := "delete success!"
    t.Data["json"] = &map[string]interface{}{"error": false, "message": mes}
  } else {
    mes := "Is not Auntification :("
    t.Data["json"] = &map[string]interface{}{"error": true, "message": mes}
  }
  t.ServeJson()
}