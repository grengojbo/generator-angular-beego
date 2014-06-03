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

// <%= classedName %> model
type <%= classedName %> struct {
	ID      int64     `orm:"column(id);auto;pk"`
	Name    string    `orm:"size(255);null;index"`
	Created time.Time `orm:"auto_now_add;type(datetime);null"`
	Updated time.Time `orm:"auto_now;type(datetime);null"`
}

// Set ORM Table name
func (o *<%= classedName %>) TableName() string {
	return "<%= tableName %>"
}

// Insert new record to <%= classedName %> model
func (o *<%= classedName %>) Insert() error {
	if _, err := orm.NewOrm().Insert(o); err != nil {
		return err
	}
	return nil
}

// Read only record from <%= classedName %> model
// ob := models.<%= classedName %>{ID: 1, Name: "NoName"}
// if err := ob.Read("ID", "Name"); err != nil {
// 	beego.Error("<%= classedName %> Read:", err.Error())
// }
// beego.Debug("<%= classedName %> Read:", ob.Created)
func (o *<%= classedName %>) Read(fields ...string) error {
	if err := orm.NewOrm().Read(o, fields...); err != nil {
		return err
	}
	return nil
}

// Update records from <%= classedName %> model
// ob.Status = 1
// if err := ob.Update("Status"); err != nil {
//   beego.Error("No update id:", ob.ID, err.Error())
// }
func (o *<%= classedName %>) Update(fields ...string) error {
	// fields = append(fields, "Updated")
	if _, err := orm.NewOrm().Update(o, fields...); err != nil {
		return err
	}
	return nil
}

// Delete records from <%= classedName %> model
// ob := models.<%= classedName %>{ID: 1}
// if err := pay.Delete(); err != nill {
// 	beego.Error("<%= classedName %> Delete", err.Error())
// }
func (o *<%= classedName %>) Delete() error {
	if _, err := orm.NewOrm().Delete(o); err != nil {
		return err
	}
	return nil
}

// Get<%= classedName %> returns one record from the database
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

// Get<%= classedName %>List returns all records from the database, sorted in the field sort
// Return []orm.Params
func Get<%= classedName %>List(sort string) (objects []orm.Params, count int64) {
	o := orm.NewOrm()
	p := new(<%= classedName %>)
	count, err := o.QueryTable(p).OrderBy(sort).Values(&objects)
	if err == nil {
		beego.Debug("Get<%= classedName %>List Result count: ", count)
	}
	return objects, count
}

// GetAll<%= classedName %>List returns all records from the database, sorted in the field sort
// Return:
// 					objects []*<%= classedName %>
// 					cnt     Count rows
// 					err     Error
func GetAll<%= classedName %>(sort string) (objects []*<%= classedName %>, cnt int64, err error) {
	o := orm.NewOrm()

	objects = make([]*<%= classedName %>, 0)
	table := new(<%= classedName %>)
	qs := o.QueryTable(table)
	cnt, err = qs.OrderBy(sort).All(&objects)
	
	return objects, cnt, err
}

// Delete record from ID int64
func Delete<%= classedName %>(id int64) error {
	ob := &<%= classedName %>{ID: id}
	return ob.Delete()
}

func init() {
	orm.RegisterModel(new(<%= classedName %>))
}
