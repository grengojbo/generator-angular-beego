// Copyright 2014 <%= autorName %>. All rights reserved.
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.
package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	// "github.com/astaxie/beego/validation"
	// "strconv"
	"time"
)

type <%= classedName %> struct {
	ID      int64     `orm:"auto;pk"`
	Name    string    `orm:"size(255);null;index"`
	Created time.Time `orm:"auto_now_add;type(datetime);null"`
	Updated time.Time `orm:"auto_now;type(datetime);null"`
}

func (o *<%= classedName %>) TableName() string {
	return "<%= sname %>"
}

func Get<%= classedName %>(ObjectID int64) (object <%= classedName %>, err error) {
	o := orm.NewOrm()
	object = <%= classedName %>{ID: ObjectID}
	err = o.Read(&object)
	if err == orm.ErrNoRows {
		return object, errors.New("No result found.")
	} else if err == orm.ErrMissPK {
		return object, errors.New("No primary key found.")
	} else {
		return object, nil
	}
}

func Get<%= classedName %>List(sort string) (objects []orm.Params, count int64) {
	o := orm.NewOrm()
	p := new(<%= classedName %>)
	count, err := o.QueryTable(p).OrderBy(sort).Values(&objects)
	if err == nil {
		beego.Debug("Get<%= classedName %>List Result count: ", count)
	}
	return objects, count
}

func init() {
	orm.RegisterModel(new(<%= classedName %>))
}
