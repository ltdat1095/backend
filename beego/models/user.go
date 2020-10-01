package models

import (
	"reflect"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

// User modal Struct
type User struct {
	Id            int
	Email         string `orm:"size(64);unique" valid:"Required;Email"`
	Password      string `orm:"size(64);valid:"Required;MinSize(6)"`
	Age           int
	FirstName     string    `orm:"size(100)"`
	LastName      string    `orm:"size(100)"`
	LastLoginTime time.Time `orm:"type(datetime);null" form:"-"`
	Created       time.Time `orm:"type(datetime);auto_now_add"`
	Updated       time.Time `orm:"type(datetime);auto_now_add"`
}

func Users() orm.QuerySeter {
	return orm.NewOrm().QueryTable("user").OrderBy("-Id")
}

// GetUsers order by ID
func GetUsers() ([]*User, error) {
	o := orm.NewOrm()

	var users []*User
	if _, err := o.QueryTable("user").OrderBy("-Id").All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (m *User) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *User) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) ReadOrCreate(field string, fields ...string) (bool, int64, error) {
	return orm.NewOrm().ReadOrCreate(m, field, fields...)
}

func (m *User) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *User) IsEmpty() bool {
	return reflect.DeepEqual(m, User{})
}
