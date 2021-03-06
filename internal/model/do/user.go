// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta    `orm:"table:user, do:true"`
	Id        interface{} // 用户编号
	Avatar    interface{} // 头像
	Nickname  interface{} // 用户昵称
	Phone     interface{} // 手机号
	Password  interface{} // 密码
	Status    interface{} // 状态:0禁用1正常
	CreatedAt *gtime.Time // 创建时间
}
