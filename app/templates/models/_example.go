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

var (
	<%= classedName %>s map[string]*<%= classedName %>
)

type <%= classedName %> struct {
	ID      int64     `orm:"column(id);auto;pk"`
	Name    string    `orm:"size(255);null;index"`
	Created time.Time `orm:"auto_now_add;type(datetime);null"`
	Updated time.Time `orm:"auto_now;type(datetime);null"`
}

func (o *<%= classedName %>) TableName() string {
	return "<%= tableName %>"
}

func (o *<%= classedName %>) Insert() error {
	if _, err := orm.NewOrm().Insert(o); err != nil {
		return err
	}
	return nil
}

func (o *<%= classedName %>) Read(fields ...string) error {
	if err := orm.NewOrm().Read(o, fields...); err != nil {
		return err
	}
	return nil
}

func (o *<%= classedName %>) Update(fields ...string) error {
	fields = append(fields, "Updated")
	if _, err := orm.NewOrm().Update(o, fields...); err != nil {
		return err
	}
	return nil
}

func (o *<%= classedName %>) Delete() error {
	if _, err := orm.NewOrm().Delete(o); err != nil {
		return err
	}
	return nil
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

func GetAll<%= classedName %>(sort string) (objects []*<%= classedName %>, cnt int64, err error) {
	o := orm.NewOrm()

	objects = make([]*<%= classedName %>, 0)
	table := new(<%= classedName %>)
	qs := o.QueryTable(table)
	cnt, err = qs.OrderBy(sort).All(&objects)
	
	return objects, cnt, err
}

func Delete<%= classedName %>(id int64) error {
	ob := &<%= classedName %>{ID: id}
	return ob.Delete()
}

func init() {
	orm.RegisterModel(new(<%= classedName %>))
}
