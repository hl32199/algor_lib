package algor_lib

import (
	"reflect"
	"testing"
)

func (n *authNode) getSortedChildIds() []uint32 {
	ids := make([]uint32, 0, len(n.children))
	for _, node := range n.children {
		ids = addToSortedSlice(ids, node.id)
	}

	return ids
}

func TestInitAuthMap(t *testing.T) {
	list := []authItem{
		{23, 231}, {1, 11}, {23, 233}, {0, 3}, {2, 22}, {2, 23},
		{21, 211}, {23, 232}, {2, 21}, {0, 1},
	}
	//ids := [...]uint32{1,2,3,11,21,22,23,231,211,232,233}
	m, err := InitAuthMap(list)
	if err == nil {
		t.Fatalf("expect err,got nil")
	}
	list = append(list, authItem{0, 2})
	m, err = InitAuthMap(list)
	if err != nil {
		t.Fatalf("expect nil,got %s", err)
	}
	//t.Log(*m[0])
	for _, item := range list {
		node, ok := m[item.id]
		if !ok {
			t.Fatalf("item not in map,id:%d", item.id)
		}
		if node.parent.id != item.parent {
			t.Fatalf("wrong node parent,id:%d,parent:%d", item.id, node.parent.id)
		}
	}

	var ids, exp []uint32
	ids = m[0].getSortedChildIds()
	exp = []uint32{1, 2, 3}
	if !reflect.DeepEqual(ids, exp) {
		t.Fatalf("expect %v,got %v", exp, ids)
	}
	ids = m[1].getSortedChildIds()
	exp = []uint32{11}
	if !reflect.DeepEqual(ids, exp) {
		t.Fatalf("expect %v,got %v", exp, ids)
	}
	ids = m[2].getSortedChildIds()
	exp = []uint32{21, 22, 23}
	if !reflect.DeepEqual(ids, exp) {
		t.Fatalf("expect %v,got %v", exp, ids)
	}
	ids = m[21].getSortedChildIds()
	exp = []uint32{211}
	if !reflect.DeepEqual(ids, exp) {
		t.Fatalf("expect %v,got %v", exp, ids)
	}
	ids = m[22].getSortedChildIds()
	exp = []uint32{}
	if !reflect.DeepEqual(ids, exp) {
		t.Fatalf("expect %v,got %v", exp, ids)
	}
	ids = m[23].getSortedChildIds()
	exp = []uint32{231, 232, 233}
	if !reflect.DeepEqual(ids, exp) {
		t.Fatalf("expect %v,got %v", exp, ids)
	}
	ids = m[3].getSortedChildIds()
	exp = []uint32{}
	if !reflect.DeepEqual(ids, exp) {
		t.Fatalf("expect %v,got %v", exp, ids)
	}
	ids = m[232].getSortedChildIds()
	exp = []uint32{}
	if !reflect.DeepEqual(ids, exp) {
		t.Fatalf("expect %v,got %v", exp, ids)
	}
}

func TestAddDelAuth(t *testing.T) {
	list := []authItem{
		{23, 231}, {1, 11}, {23, 233}, {0, 3}, {2, 22}, {2, 23},
		{21, 211}, {23, 232}, {2, 21}, {0, 1}, {0, 2},
	}
	m, _ = InitAuthMap(list)

	var ori []uint32
	ori, err := AddAuth(ori, 1)
	if err != nil {
		t.Fatalf("expect nil,got %s", err)
	}
	var exp []uint32
	exp = []uint32{1, 11}
	if !reflect.DeepEqual(ori, exp) {
		t.Fatalf("expect %v,got %v", exp, ori)
	}
	ori, _ = AddAuth(ori, 3)
	exp = []uint32{1, 3, 11}
	if !reflect.DeepEqual(ori, exp) {
		t.Fatalf("expect %v,got %v", exp, ori)
	}
	ori, _ = AddAuth(ori, 211)
	exp = []uint32{1, 3, 11, 21, 211}
	if !reflect.DeepEqual(ori, exp) {
		t.Fatalf("expect %v,got %v", exp, ori)
	}
	ori, _ = AddAuth(ori, 22)
	exp = []uint32{1, 3, 11, 21, 22, 211}
	if !reflect.DeepEqual(ori, exp) {
		t.Fatalf("expect %v,got %v", exp, ori)
	}
	ori, _ = DelAuth(ori, 3)
	exp = []uint32{1, 11, 21, 22, 211}
	if !reflect.DeepEqual(ori, exp) {
		t.Fatalf("expect %v,got %v", exp, ori)
	}
	ori, _ = DelAuth(ori, 11)
	exp = []uint32{21, 22, 211}
	if !reflect.DeepEqual(ori, exp) {
		t.Fatalf("expect %v,got %v", exp, ori)
	}
	ori, _ = AddAuth(ori, 23)
	exp = []uint32{2, 21, 22, 23, 211, 231, 232, 233}
	if !reflect.DeepEqual(ori, exp) {
		t.Fatalf("expect %v,got %v", exp, ori)
	}
	ori, _ = DelAuth(ori, 232)
	exp = []uint32{21, 22, 211, 231, 233}
	if !reflect.DeepEqual(ori, exp) {
		t.Fatalf("expect %v,got %v", exp, ori)
	}
}
