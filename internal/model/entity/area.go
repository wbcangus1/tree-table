// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Area is the golang structure for table area.
type Area struct {
	Id        int         `json:"id"        orm:"id"         description:"区域ID"`                         // 区域ID
	Name      string      `json:"name"      orm:"name"       description:"名称"`                           // 名称
	ParentId  int         `json:"parentId"  orm:"parent_id"  description:"父区域ID"`                        // 父区域ID
	NodePath  string      `json:"nodePath"  orm:"node_path"  description:"节点路径, 根节点id/../上级节点id/当前节点id"` // 节点路径, 根节点id/../上级节点id/当前节点id
	Sort      int         `json:"sort"      orm:"sort"       description:"同层级排序"`                        // 同层级排序
	Level     int         `json:"level"     orm:"level"      description:"节点所在层级"`                       // 节点所在层级
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`                         // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"最后更新时间"`                       // 最后更新时间
}
