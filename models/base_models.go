package models

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julycw/Jeremiah/models/mgo_models"
	"net/url"
	"time"
)

var Orm orm.Ormer

func init() {
	orm.DefaultTimeLoc = time.Local
	var connstr string
	connstr = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		beego.AppConfig.String("user"),
		beego.AppConfig.String("passwd"),
		beego.AppConfig.String("host"),
		beego.AppConfig.String("port"),
		beego.AppConfig.String("db")) + "&loc=" + url.QueryEscape("Local")
	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Avaliable))
	orm.RegisterModel(new(FixOrder))
	orm.RegisterModel(new(FixOrderDetail))
	orm.RegisterModel(new(FixOrderPair))
	orm.RegisterModel(new(FollowUser))
	orm.RegisterModel(new(Log))
	orm.RegisterModel(new(Role))
	orm.RegisterModel(new(Message))
	orm.RegisterModel(new(Note))
	orm.RegisterModel(new(Option))
	orm.RegisterModel(new(News))
	orm.RegisterModel(new(Dictionary))
	orm.RegisterModel(new(Product))

	mgo_models.DialDB("127.0.0.1", "27017")
	orm.RegisterDataBase("default", "mysql", connstr)

	Orm = orm.NewOrm()
	Orm.Using("defualt")
}
