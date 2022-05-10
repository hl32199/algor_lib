package algor_lib

import (
	"fmt"
)

type authNode struct {
	id       uint32
	parent   *authNode
	children []*authNode
}

type authItem struct {
	parent uint32
	id     uint32
}

func (n *authNode) getSortedChildIds() []uint32 {
	ids := make([]uint32, 0, len(n.children))
	for _, node := range n.children {
		ids = addToSortedSlice(ids, node.id)
	}

	return ids
}

func addToSortedSlice(list []uint32, item uint32) []uint32 {
	if len(list) == 0 || item >= list[len(list)-1] {
		list = append(list, item)
		return list
	}

	for i := 0; i < len(list); i++ {
		if item < list[i] {
			list = append(list[:i+1], list[i:]...)
			list[i] = item
			break
		}
	}
	return list
}

var m map[uint32]*authNode

//初始化权限map，key为权限id，值为权限节点，另外权限节点会以树状组织
func InitAuthMap(list []authItem) (map[uint32]*authNode, error) {
	var root = &authNode{
		id: 0,
	}
	var nodeMap = make(map[uint32]*authNode, len(list))
	nodeMap[0] = root
	for _, item := range list {
		nodeMap[item.id] = &authNode{
			id:       item.id,
			children: nil,
		}
	}

	for _, item := range list {
		node := nodeMap[item.id]
		parent, ok := nodeMap[item.parent]
		if !ok {
			return nil, fmt.Errorf("nonexistent parent,id=%d", item.parent)
		}
		node.parent = parent
		parent.children = append(parent.children, node)
	}

	return nodeMap, nil
}

//新增一项权限，origin为原权限集合，id为要新增的权限
func AddAuth(originAuth []uint32, id uint32) ([]uint32, error) {
	node, ok := m[id]
	if !ok {
		return nil, fmt.Errorf("nonexistent auth id:%d", id)
	}

	authMap := make(map[uint32]struct{}, len(originAuth)+3)
	for _, ori := range originAuth {
		authMap[ori] = struct{}{}
	}
	//目标节点已存在，直接返回原权限列表
	if _, ok := authMap[id]; ok {
		return originAuth, nil
	}

	//加入目标节点及其所有子孙节点
	needCheckNode := make([]*authNode, 0, len(node.children)*2+1)
	needCheckNode = append(needCheckNode, node)
	for len(needCheckNode) > 0 {
		item := needCheckNode[len(needCheckNode)-1]
		needCheckNode = needCheckNode[:len(needCheckNode)-1]
		authMap[item.id] = struct{}{}
		if len(item.children) > 0 {
			for _, cnode := range item.children {
				needCheckNode = append(needCheckNode, cnode)
			}
		}
	}

	//往上逐级检查父节点是否需要加入
	pnode := node.parent
	for {
		//已到根节点，退出
		if pnode.id == 0 {
			break
		}

		if _, ok := authMap[pnode.id]; ok {
			//不应该出现这种情况，如果有，说明之前给用户编辑权限已出错
			return nil, fmt.Errorf("子权限未加入而父权限已加入，parent:%d", pnode.id)
		}

		for _, n := range pnode.children {
			if _, ok := authMap[n.id]; !ok {
				//此处可直接退出循环，因为只要有一个节点未加入，父节点即不用加入，再往上也不用加入
				goto end
			}
		}
		//到这里说明当前节点所有子节点都在，则当前节点需要加入，接着检查上一级父节点
		authMap[pnode.id] = struct{}{}

	end:
		pnode = pnode.parent
	}

	newAuth := make([]uint32, 0, len(authMap))
	for id, _ := range authMap {
		newAuth = addToSortedSlice(newAuth, id)
	}
	return newAuth, nil
}

//删除一项权限，origin为原权限集合，id为要新增的权限
func DelAuth(originAuth []uint32, id uint32) ([]uint32, error) {
	node, ok := m[id]
	if !ok {
		return nil, fmt.Errorf("nonexistent auth id:%d", id)
	}

	authMap := make(map[uint32]struct{}, len(originAuth))
	for _, ori := range originAuth {
		authMap[ori] = struct{}{}
	}
	//目标节点已不存在，直接返回原权限列表
	if _, ok := authMap[id]; !ok {
		return originAuth, nil
	}

	//删除目标节点及其所有子孙节点
	needCheckNode := make([]*authNode, 0, len(node.children)*2+1)
	needCheckNode = append(needCheckNode, node)
	for len(needCheckNode) > 0 {
		item := needCheckNode[len(needCheckNode)-1]
		needCheckNode = needCheckNode[:len(needCheckNode)-1]
		delete(authMap, item.id)
		if len(item.children) > 0 {
			for _, cnode := range item.children {
				needCheckNode = append(needCheckNode, cnode)
			}
		}
	}

	//逐级删除所有父节点
	pnode := node.parent
	for {
		//已到根节点，退出
		if pnode.id == 0 {
			break
		}

		if _, ok := authMap[pnode.id]; !ok {
			//当前节点不存在，即可退出循环，因为往上的父节点肯定也不存在
			break
		}
		delete(authMap, pnode.id)
		pnode = pnode.parent
	}

	newAuth := make([]uint32, 0, len(authMap))
	for id, _ := range authMap {
		newAuth = addToSortedSlice(newAuth, id)
	}
	return newAuth, nil
}
