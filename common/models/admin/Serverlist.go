package admin

import (
	"github.com/astaxie/beego/orm"
)

type Serverlist struct {
	Id          int `orm:"column(id);pk;auto" json:"id"`
	IpAddr      string
	MachineRoom string `orm:varchar(30)`
	UniqueKey   string `orm:varchar(255)`
}

func init() {
	orm.RegisterModelWithPrefix("admin_", new(Serverlist))
}
func (u *Serverlist) TableName() string {
	return "serverlist"
}
func (this *Serverlist) getOrm() orm.Ormer {
	o := orm.NewOrm()
	o.Using("db_admin") // 默认使用 default，你可以指定为其他数据库
	return o
}