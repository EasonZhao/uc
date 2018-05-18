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
	orm.RegisterModel(new(User), new(GAAuthenticator))
}

func Authenticate(username, password string) *User {

	return nil
}

type GAAuthenticator struct {
	Id      int
	User    *User  `orm:"reverse(one)"`
	Sercet  string `orm:"null;size(40)"`
	Account string `orm:"null;size(36)"`
	Authed  bool   `orm:"default(false)"`
}

func (this *GAAuthenticator) TableName() string {
	return "ga_auth"
}

type User struct {
	Id          int
	Nationality string           `orm:"null;size(20)"`
	PhoneNum    string           `orm:"unique;null;size(20)"`
	Email       string           `orm:"unique;null;size(36)"`
	Username    string           `orm:"unique;size(36)"`
	Password    string           `orm:"size(20)"`
	RegTime     time.Time        `orm:"auto_now_add;type(datatime)"`
	AuthPhone   bool             `orm:"default(false)"`
	AuthEmail   bool             `orm:"default(false)"`
	GAAuth      *GAAuthenticator `orm:"null;rel(one);on_delete(set_null)"`
}

func (this *User) TableEngine() string {
	return "INNODB AUTO_INCREMENT=100028"
}

func (this *User) AuthGA(sercet string) error {
	o := orm.NewOrm()
	this.GAAuth.Authed = true
	this.GAAuth.Account = this.Username
	this.GAAuth.Sercet = sercet
	if _, err := o.Update(this.GAAuth); err != nil {
		return err
	}
	return nil
}

func NewUser() *User {
	u := new(User)
	u.GAAuth = new(GAAuthenticator)
	return u
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
	user := NewUser()
	user.Email = email
	user.Password = password
	user.Username = email
	user.AuthEmail = true
	if existEmail(email) {
		str := fmt.Sprintf("email %s already register", email)
		return user, errors.New(str)
	}
	o := orm.NewOrm()
	o.Begin()
	if _, err := o.Insert(user.GAAuth); err != nil {
		o.Rollback()
		return nil, err
	}
	if _, err := o.Insert(user); err != nil {
		o.Rollback()
		return nil, err
	}
	if err := o.Commit(); err != nil {
		o.Rollback()
		return nil, err
	}
	return user, nil
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
	if _, err := o.LoadRelated(&user, "GAAuth"); err != nil {
		return nil, err
	}
	return &user, nil
}
