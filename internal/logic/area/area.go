package area

import (
	"context"
	"fmt"
	"strings"
	v1 "tree-table/api/hello/v1"
	"tree-table/internal/service"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

// AreaService 区域服务
type sAreaService struct{}

func init() {
	service.RegisterAreaService(new(sAreaService))
}

// Create 创建区域
func (s *sAreaService) Create(ctx context.Context, req *v1.AreaCreateReq) (id int64, err error) {
	// 1. 开启事务
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 2. 检查父节点是否存在
		var parentArea *v1.Area
		if req.ParentId > 0 {
			err = tx.Model("area").Where("id", req.ParentId).Scan(&parentArea)
			if err != nil {
				return err
			}
			if parentArea == nil {
				return gerror.New("父区域不存在")
			}
		}

		// 3. 计算level值
		level := int8(1) // 默认为1级(根节点)
		if parentArea != nil {
			level = parentArea.Level + 1

		}

		// 4. 获取同级最大排序值
		var maxSort int
		_, err = tx.Model("area"). // 111
						Where("parent_id", req.ParentId).
						Order("sort DESC").
						Limit(1).
						Value(&maxSort, "sort")
		if err != nil {
			return err
		}
		sort := maxSort + 1

		// 5. 创建区域记录
		area := &v1.Area{
			ParentId:    req.ParentId,
			Name:        req.Name,
			Level:       level,
			Sort:        sort,
			Description: req.Description,
			Tags:        req.Tags,
		}

		result, err := tx.Model("area").Data(area).Insert()
		if err != nil {
			return err
		}

		// 6. 获取新插入的ID
		lastId, err := result.LastInsertId()
		if err != nil {
			return err
		}

		// 7. 更新path值
		path := fmt.Sprintf("%d", lastId)
		if parentArea != nil {
			path = fmt.Sprintf("%s,%d", parentArea.Path, lastId)
		}

		// 8. 更新记录的path字段
		_, err = tx.Model("area").
			Where("id", lastId).
			Data(g.Map{"path": path}).
			Update()
		if err != nil {
			return err
		}

		id = lastId
		return nil
	})

	return id, err
}

// Update 更新区域
func (s *sAreaService) Update(ctx context.Context, req *v1.AreaUpdateReq) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 检查区域是否存在
		var area *v1.Area
		err := tx.Model("area").Where("id", req.Id).Scan(&area)
		if err != nil {
			return err
		}
		if area == nil {
			return gerror.New("区域不存在")
		}

		// 2. 准备更新数据
		updateData := g.Map{}
		if req.Name != "" {
			updateData["name"] = req.Name
		}
		if req.Description != "" {
			updateData["description"] = req.Description
		}
		if req.Tags != "" {
			updateData["tags"] = req.Tags
		}
		if req.Sort > 0 {
			updateData["sort"] = req.Sort
		}

		// 3. 执行更新
		if len(updateData) > 0 {
			_, err = tx.Model("area").Where("id", req.Id).Data(updateData).Update()
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// Delete 删除区域（递归删除子区域）
func (s *sAreaService) Delete(ctx context.Context, id int) error {
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 1. 检查区域是否存在
		var area *v1.Area
		err := tx.Model("area").Where("id", id).Scan(&area)
		if err != nil {
			return err
		}
		if area == nil {
			return gerror.New("区域不存在")
		}

		// 2. 查询所有需要删除的子区域ID（使用path字段高效查询）
		var ids []int
		err = tx.Model("area").
			WhereLike("path", fmt.Sprintf("%s%%", area.Path)).
			Fields("id").
			Scan(&ids)
		if err != nil {
			return err
		}

		// 3. 批量删除所有子区域和当前区域
		if len(ids) > 0 {
			_, err = tx.Model("area").WhereIn("id", ids).Delete()
			if err != nil {
				return err
			}
		}

		return nil
	})
}

// GetList 获取区域列表（分页）
func (s *sAreaService) GetList(ctx context.Context, page, size int) (list []*v1.Area, total int, err error) {
	model := g.DB().Model("area")
	total, err = model.Count()
	if err != nil {
		return nil, 0, err
	}

	err = model.Page(page, size).Order("level ASC, sort ASC").Scan(&list)
	return list, total, err
}

// GetTree 获取区域树结构
func (s *sAreaService) GetTree(ctx context.Context, parentId ...int) ([]*v1.AreaTreeNode, error) {
	// 1. 查询所有区域
	var areas []*v1.Area
	model := g.DB().Model("area").Order("sort ASC")

	rootId := 0
	if len(parentId) > 0 && parentId[0] > 0 {
		rootId = parentId[0]
	}

	err := model.Scan(&areas)
	if err != nil {
		return nil, err
	}

	// 2. 构建树形结构
	return s.buildTree(areas, rootId), nil
}

// GetChildren 获取直接子区域
func (s *sAreaService) GetChildren(ctx context.Context, parentId int) ([]*v1.Area, error) {
	var areas []*v1.Area
	err := g.DB().Model("area").
		Where("parent_id", parentId).
		Order("sort ASC").
		Scan(&areas)
	return areas, err
}

// GetNodePath 获取节点路径（从根到当前节点的完整路径）
func (s *sAreaService) GetNodePath(ctx context.Context, id int) ([]*v1.Area, error) {
	// 1. 查询当前节点
	var area *v1.Area
	err := g.DB().Model("area").Where("id", id).Scan(&area)
	if err != nil {
		return nil, err
	}
	if area == nil {
		return nil, gerror.New("区域不存在")
	}

	// 2. 解析路径
	pathIds := strings.Split(area.Path, ",")
	var ids []int
	for _, idStr := range pathIds {
		ids = append(ids, gconv.Int(idStr))
	}

	// 3. 查询路径上的所有节点
	var pathNodes []*v1.Area
	err = g.DB().Model("area").
		WhereIn("id", ids).
		Order("level ASC").
		Scan(&pathNodes)
	return pathNodes, err
}

// AdjustSort 调整排序
func (s *sAreaService) AdjustSort(ctx context.Context, id, sort int) error {
	// 1. 检查区域是否存在
	var area *v1.Area
	err := g.DB().Model("area").Where("id", id).Scan(&area)
	if err != nil {
		return err
	}
	if area == nil {
		return gerror.New("区域不存在")
	}

	// 2. 更新排序
	_, err = g.DB().Model("area").
		Where("id", id).
		Data(g.Map{"sort": sort}).
		Update()
	return err
}

// 构建树形结构
func (s *sAreaService) buildTree(areas []*v1.Area, parentId int) []*v1.AreaTreeNode {
	var nodes []*v1.AreaTreeNode
	for _, area := range areas {
		if area.ParentId == parentId {
			node := &v1.AreaTreeNode{
				Area:     area,
				Children: s.buildTree(areas, area.Id),
			}
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func (s *sAreaService) BuildTree004854(areas []*v1.Area, parentId int) []*v1.AreaTreeNode {
	var nodes []*v1.AreaTreeNode
	fmt.Println(666)
	for _, area := range areas {
		if area.ParentId == parentId {
			node := &v1.AreaTreeNode{
				Area:     area,
				Children: s.buildTree(areas, area.Id),
			}
			nodes = append(nodes, node)
		}
	}
	return nodes
}
