// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Area is the golang structure of table area for DAO operations like Where/Data.
type Area struct {
	g.Meta    `orm:"table:area, do:true"`
	Id        interface{} // 区域ID
	Name      interface{} // 名称
	ParentId  interface{} // 父区域ID
	NodePath  interface{} // 节点路径, 根节点id/../上级节点id/当前节点id
	Sort      interface{} // 同层级排序
	Level     interface{} // 节点所在层级
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 最后更新时间
}
