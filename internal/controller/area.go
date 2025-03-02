package controller

// AreaController 区域控制器

import (
	"context"
	v1 "tree-table/api/hello/v1"
	"tree-table/internal/service"
)

// AreaController 区域控制器
type cAreaController struct{}

var AreaController = &cAreaController{}

// Create 创建区域
func (c *cAreaController) Create(ctx context.Context, req *v1.AreaCreateReq) (res *v1.AreaCreateRes, err error) {
	id, err := service.AreaService().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.AreaCreateRes{Id: id}, nil
}

// Update 更新区域
func (c *cAreaController) Update(ctx context.Context, req *v1.AreaUpdateReq) (res *v1.AreaUpdateRes, err error) {
	err = service.AreaService().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.AreaUpdateRes{Success: true}, nil
}

// Delete 删除区域
func (c *cAreaController) Delete(ctx context.Context, req *v1.AreaDeleteReq) (res *v1.AreaDeleteRes, err error) {
	err = service.AreaService.Delete(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.AreaDeleteRes{Success: true}, nil
}

// GetList 获取区域列表
func (c *cAreaController) GetList(ctx context.Context, req *v1.AreaGetListReq) (res *v1.AreaGetListRes, err error) {
	list, total, err := service.AreaService.GetList(ctx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}
	return &v1.AreaGetListRes{
		List:  list,
		Total: total,
	}, nil
}

// GetTree 获取区域树
func (c *cAreaController) GetTree(ctx context.Context, req *v1.AreaGetTreeReq) (res *v1.AreaGetTreeRes, err error) {
	tree, err := service.AreaService.GetTree(ctx, req.ParentId)
	if err != nil {
		return nil, err
	}
	return &v1.AreaGetTreeRes{Tree: tree}, nil
}

// GetChildren 获取子区域
func (c *cAreaController) GetChildren(ctx context.Context, req *v1.AreaGetChildrenReq) (res *v1.AreaGetChildrenRes, err error) {
	list, err := service.AreaService.GetChildren(ctx, req.ParentId)
	if err != nil {
		return nil, err
	}
	return &v1.AreaGetChildrenRes{List: list}, nil
}

// GetNodePath 获取节点路径
func (c *cAreaController) GetNodePath(ctx context.Context, req *v1.AreaGetNodePathReq) (res *v1.AreaGetNodePathRes, err error) {
	path, err := service.AreaService.GetNodePath(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.AreaGetNodePathRes{Path: path}, nil
}

// AdjustSort 调整排序
func (c *cAreaController) AdjustSort(ctx context.Context, req *v1.AreaAdjustSortReq) (res *v1.AreaAdjustSortRes, err error) {
	err = service.AreaService.AdjustSort(ctx, req.Id, req.Sort)
	if err != nil {
		return nil, err
	}
	return &v1.AreaAdjustSortRes{Success: true}, nil
}
