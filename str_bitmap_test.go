package algor_lib

import (
	"strings"
	"testing"
)

func TestStrBitmap_SetBit(t *testing.T) {
	bm := GetStrBitMap("")
	if err := bm.SetBit(0); err == nil || !strings.Contains(err.Error(), "greater than 0") {
		t.Fatalf("expect err 'pos should be greater than 0',got %+v", err)
	}
	if err := bm.UnsetBit(0); err == nil || !strings.Contains(err.Error(), "greater than 0") {
		t.Fatalf("expect err 'pos should be greater than 0',got %+v", err)
	}

	items := [][]uint32{
		{1},
		{2, 5, 7},
		{9, 15, 20, 23},
		{10, 27, 100, 201, 222},
	}

	for _, item := range items {
		bm := GetStrBitMap("")
		for _, bit := range item {
			err := bm.SetBit(bit)
			if err != nil {
				t.Fatalf("call SetBit got err:%s,item:%+v,bit:%d", err, item, bit)
			}
		}

		for _, bit := range item {
			if !bm.GetBit(bit) {
				t.Fatalf("GetBit %d wrong,item:%+v", bit, item)
			}
		}
	}

	bm = GetStrBitMap("")
	for _, bit := range []uint32{10, 27, 100, 201, 222} {
		bm.SetBit(bit)
	}

	bm.UnsetBit(100)
	t.Logf("byte_count:%d,raw_str:%s", len([]byte(bm.GetStr())), bm.GetStr())
	if !checkBit(&bm, []uint32{10, 27, 201, 222}) {
		t.Fatalf("expect got bit:%+v", []uint32{10, 27, 201, 222})
	}
	if bm.GetBit(100) {
		t.Fatalf("should not has bit 100")
	}

	bm.UnsetBit(105)
	t.Logf("byte_count:%d,raw_str:%s", len([]byte(bm.GetStr())), bm.GetStr())
	if !checkBit(&bm, []uint32{10, 27, 201, 222}) {
		t.Fatalf("expect got bit:%+v", []uint32{10, 27, 201, 222})
	}
	if bm.GetBit(105) {
		t.Fatalf("should not has bit 105")
	}

	bm.UnsetBit(222)
	t.Logf("byte_count:%d,raw_str:%s", len([]byte(bm.GetStr())), bm.GetStr())
	if !checkBit(&bm, []uint32{10, 27, 201}) {
		t.Fatalf("expect got bit:%+v", []uint32{10, 27, 201, 222})
	}
	if bm.GetBit(222) {
		t.Fatalf("should not has bit 222")
	}

	bm.UnsetBit(223)
	bm.UnsetBit(221)
	bm.UnsetBit(202)
	bm.UnsetBit(200)
	bm.UnsetBit(400)
	t.Logf("byte_count:%d,raw_str:%s", len([]byte(bm.GetStr())), bm.GetStr())
	if !checkBit(&bm, []uint32{10, 27, 201}) {
		t.Fatalf("expect got bit:%+v", []uint32{10, 27, 201})
	}
	if bm.GetBit(223) {
		t.Fatalf("should not has bit 223")
	}
	if bm.GetBit(221) {
		t.Fatalf("should not has bit 221")
	}
	if bm.GetBit(202) {
		t.Fatalf("should not has bit 202")
	}
	if bm.GetBit(200) {
		t.Fatalf("should not has bit 200")
	}
	if bm.GetBit(400) {
		t.Fatalf("should not has bit 400")
	}

	bm.SetBit(1)
	bm.SetBit(7)
	bm.SetBit(8)
	bm.UnsetBit(8)
	if !checkBit(&bm, []uint32{1, 7, 10, 27, 201}) {
		t.Fatalf("expect got bit:%+v", []uint32{1, 7, 10, 27, 201})
	}
	if bm.GetBit(8) {
		t.Fatalf("should not has bit 8")
	}
	bm.UnsetBit(7)
	if !checkBit(&bm, []uint32{1, 10, 27, 201}) {
		t.Fatalf("expect got bit:%+v", []uint32{1, 7, 10, 27, 201})
	}
	if bm.GetBit(7) {
		t.Fatalf("should not has bit 7")
	}
	if bm.GetBit(8) {
		t.Fatalf("should not has bit 8")
	}
}

func checkBit(bm *StrBitmap, bits []uint32) bool {
	for _, bit := range bits {
		if !bm.GetBit(bit) {
			return false
		}
	}
	return true
}

func TestStrBitmap_GetBit(t *testing.T) {
	bm := GetStrBitMap("a0A")
	expBits := []uint32{1, 6, 7, 13, 14, 17, 23}
	unExpBits := []uint32{0, 2, 8, 9, 16, 24, 35}

	for _, bit := range expBits {
		if !bm.GetBit(bit) {
			t.Fatalf("GetBit %d wrong,should return true", bit)
		}
	}
	for _, bit := range unExpBits {
		if bm.GetBit(bit) {
			t.Fatalf("GetBit %d wrong,should return false", bit)
		}
	}
}

func TestStrBitmap_GetStr(t *testing.T) {
	bm := GetStrBitMap("")
	expBits := []uint32{1, 6, 7, 13, 14, 17, 23}

	for _, bit := range expBits {
		err := bm.SetBit(bit)
		if err != nil {
			t.Fatalf("call SetBit got err:%s,bit:%d", err, bit)
		}
	}

	if bm.GetStr() != "a0A" {
		t.Fatalf("call GetStr wrong,got:%s,expect:%s", bm.GetStr(), "a0A")
	}
}
