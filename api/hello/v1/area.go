package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 区域模型
type Area struct {
	Id          int         `json:"id"           description:"区域ID"`
	ParentId    int         `json:"parentId"     description:"父区域ID"`
	Name        string      `json:"name"         description:"名称"`
	Level       int8        `json:"level"        description:"层级"`
	Sort        int         `json:"sort"         description:"同层级排序"`
	Path        string      `json:"path"         description:"节点路径"`
	Description string      `json:"description"  description:"描述"`
	Tags        string      `json:"tags"         description:"标签Json数组"`
	CreatedAt   *gtime.Time `json:"createdAt"    description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"    description:"最后更新时间"`
}

// 区域树形节点
type AreaTreeNode struct {
	*Area
	Children []*AreaTreeNode `json:"children"`
}

// 创建区域请求
type AreaCreateReq struct {
	g.Meta      `path:"/create" tags:"创建区域" method:"get" summary:"区域"`
	ParentId    int    `json:"parentId" v:"required"   dc:"父区域ID"`
	Name        string `json:"name"     v:"required"   dc:"名称"`
	Description string `json:"description"             dc:"描述"`
	Tags        string `json:"tags"                    dc:"标签Json数组"`
}

// 创建区域响应
type AreaCreateRes struct {
	Id int64 `json:"id" dc:"新创建的区域ID"`
}

// 更新区域请求
type AreaUpdateReq struct {
	g.Meta      `path:"/update" tags:"更新区域" method:"get" summary:"区域"`
	Id          int    `json:"id"       v:"required"   dc:"区域ID"`
	Name        string `json:"name"                    dc:"名称"`
	Sort        int    `json:"sort"                    dc:"同层级排序"`
	Description string `json:"description"             dc:"描述"`
	Tags        string `json:"tags"                    dc:"标签Json数组"`
}

// 更新区域响应
type AreaUpdateRes struct {
	Success bool `json:"success" dc:"是否成功"`
}

// 删除区域请求
type AreaDeleteReq struct {
	g.Meta `path:"/del" tags:"删除区域" method:"get" summary:"区域"`
	Id     int `json:"id" v:"required" dc:"区域ID"`
}

// 删除区域响应
type AreaDeleteRes struct {
	Success bool `json:"success" dc:"是否成功"`
}

// 获取区域列表请求
type AreaGetListReq struct {
	g.Meta `path:"/list" tags:"获取区域列表" method:"get" summary:"区域"`
	Page   int `json:"page" d:"1"  v:"min:1" dc:"页码"`
	Size   int `json:"size" d:"10" v:"max:100" dc:"每页数量"`
}

// 获取区域列表响应
type AreaGetListRes struct {
	List  []*Area `json:"list"  dc:"区域列表"`
	Total int     `json:"total" dc:"总数量"`
}

// 获取区域树请求
type AreaGetTreeReq struct {
	g.Meta   `path:"/tree" tags:"获取区域树" method:"get" summary:"区域"`
	ParentId int `json:"parentId" d:"0" dc:"父区域ID，默认为0表示获取完整树"`
}

// 获取区域树响应
type AreaGetTreeRes struct {
	Tree []*AreaTreeNode `json:"tree" dc:"区域树"`
}

// 获取子区域请求
type AreaGetChildrenReq struct {
	g.Meta   `path:"/children" tags:"获取子区域" method:"get" summary:"区域"`
	ParentId int `json:"parentId" v:"required" dc:"父区域ID"`
}

// 获取子区域响应
type AreaGetChildrenRes struct {
	List []*Area `json:"list" dc:"子区域列表"`
}

// 获取节点路径请求
type AreaGetNodePathReq struct {
	g.Meta `path:"/path" tags:"获取节点路径" method:"get" summary:"区域"`
	Id     int `json:"id" v:"required" dc:"区域ID"`
}

// 获取节点路径响应
type AreaGetNodePathRes struct {
	Path []*Area `json:"path" dc:"路径节点列表"`
}

// 调整排序请求
type AreaAdjustSortReq struct {
	g.Meta `path:"/adjust" tags:"调整区域" method:"get" summary:"区域"`
	Id     int `json:"id"   v:"required" dc:"区域ID"`
	Sort   int `json:"sort" v:"required" dc:"新的排序值"`
}

// 调整排序响应
type AreaAdjustSortRes struct {
	Success bool `json:"success" dc:"是否成功"`
}
