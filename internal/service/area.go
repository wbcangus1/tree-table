// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "tree-table/api/hello/v1"
)

type (
	IAreaService interface {
		// Create 创建区域
		Create(ctx context.Context, req *v1.AreaCreateReq) (id int64, err error)
		// Update 更新区域
		Update(ctx context.Context, req *v1.AreaUpdateReq) error
		// Delete 删除区域（递归删除子区域）
		Delete(ctx context.Context, id int) error
		// GetList 获取区域列表（分页）
		GetList(ctx context.Context, page int, size int) (list []*v1.Area, total int, err error)
		// GetTree 获取区域树结构
		GetTree(ctx context.Context, parentId ...int) ([]*v1.AreaTreeNode, error)
		// GetChildren 获取直接子区域
		GetChildren(ctx context.Context, parentId int) ([]*v1.Area, error)
		// GetNodePath 获取节点路径（从根到当前节点的完整路径）
		GetNodePath(ctx context.Context, id int) ([]*v1.Area, error)
		// AdjustSort 调整排序
		AdjustSort(ctx context.Context, id int, sort int) error
		BuildTree004854(areas []*v1.Area, parentId int) []*v1.AreaTreeNode
	}
)

var (
	localAreaService IAreaService
)

func AreaService() IAreaService {
	if localAreaService == nil {
		panic("implement not found for interface IAreaService, forgot register?")
	}
	return localAreaService
}

func RegisterAreaService(i IAreaService) {
	localAreaService = i
}
