package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
	"usercenter/util"
)

const (
	DB    = "mysql"
	TABLE = "user"
)

func init() {
	orm.RegisterModel(new(User))
}

func Authenticate(username, password string) *User {

	return nil
}

type User struct {
	Id             int
	Nationality    string    `orm:"null;size(20)"`
	PhoneNum       string    `orm:"unique;null;size(20)"`
	Email          string    `orm:"unique;null;size(36)"`
	Username       string    `orm:"unique;size(20)"`
	Password       string    `orm:"size(20)"`
	RegTime        time.Time `orm:"auto_now_add;type(datatime)"`
	AuthPhone      bool      `orm:"default(false)"`
	AuthEmail      bool      `orm:"default(false)"`
	AuthGoogleAuth bool      `orm:"default(false)"`
}

func (this *User) TableEngine() string {
	return "INNODB AUTO_INCREMENT=100028"
}

func existPhone(phone string) bool {
	o := orm.NewOrm()
	return o.QueryTable(TABLE).Filter("phonenum", phone).Exist()
}

func existEmail(email string) bool {
	o := orm.NewOrm()
	return o.QueryTable(TABLE).Filter("email", email).Exist()
}

func RegisterByPhone(phone, password, nationality string) (User, error) {
	user := User{}
	user.PhoneNum = phone
	user.Password = password
	user.Username = phone
	user.Nationality = nationality
	if existPhone(phone) {
		str := fmt.Sprintf("phone %s already register", phone)
		return user, errors.New(str)
	}
	o := orm.NewOrm()
	id, err := o.Insert(user)
	if id == 0 {
		return user, err
	}
	user.Id = int(id)
	return user, nil
}

func Login(username, password string) (*User, error) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(TABLE).Filter("username", username).One(&user)
	if err == orm.ErrMultiRows {
		//TODO log
	} else if err == orm.ErrNoRows {
		return nil, errors.New("username not exist.")
	}
	if user.Password != password {
		return nil, errors.New("password error.")
	}
	return &user, nil
}

func RegistByEmail(email, password string) (*User, error) {
	if !util.CheckEmail(email) {
		return nil, errors.New("email invalid.")
	}
	if !util.CheckPassword(password) {
		return nil, errors.New("password invalid.")
	}
	user := User{}
	user.Email = email
	user.Password = password
	user.Username = email
	user.AuthEmail = true
	if existEmail(email) {
		str := fmt.Sprintf("email %s already register", email)
		return &user, errors.New(str)
	}
	o := orm.NewOrm()
	id, err := o.Insert(&user)
	if err != nil {
		return nil, err
	}
	user.Id = int(id)
	return &user, nil
}

func QueryUserById(id int) (*User, error) {
	o := orm.NewOrm()
	var user User
	err := o.QueryTable(TABLE).Filter("id", id).One(&user)
	if err == orm.ErrMultiRows {
		//TODO log
	} else if err == orm.ErrNoRows {
		return nil, errors.New("username not exist.")
	}
	return &user, nil
}
