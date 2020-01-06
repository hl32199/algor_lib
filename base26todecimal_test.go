package algor_lib

import (
	"testing"
)

//26进制字符串转换为十进制字符串
func TestBase26toDecimal(t *testing.T) {
	if Base26toDecimal("z") != "25" {
		t.Fatal("wrong result")
	}

	if Base26toDecimal("zga") != "17056" {
		t.Fatal("wrong result")
	}
}

func TestIntMap_Add(t *testing.T) {
	i := IntMap{1, 2, 3, 4, 5}
	j := IntMap{8, 5, 6, 9, 0, 7}
	i = i.Add(j)
	if i.String() != "763979" {
		t.Fatalf("wrong result,i.string:%s,should:%s", i.String(), "763979")
	}
}

func TestIntMap_Multi(t *testing.T) {
	i := IntMap{8, 5, 6, 9, 0, 7}
	i = i.Multi(2361)
	if i.String() != "1675502538" {
		t.Fatalf("wrong result,i.string:%s,should:%s", i.String(), "1675502538")
	}
}

func TestIntMap_String(t *testing.T) {
	i := IntMap{1, 2, 3, 4, 5}
	if i.String() != "54321" {
		t.Fatal("wrong result")
	}

	i = IntMap{8, 5, 6, 9, 0, 7}
	if i.String() != "709658" {
		t.Fatal("wrong result")
	}
}
