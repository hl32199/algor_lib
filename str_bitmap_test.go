package algor_lib

import (
	"strings"
	"testing"
)

func TestStrBitmap_SetBit(t *testing.T) {
	bm := GetBitMap("")
	if err := bm.SetBit(0, true); err == nil || !strings.Contains(err.Error(), "greater than 0") {
		t.Fatalf("expect err 'pos should be greater than 0',got %+v", err)
	}

	items := [][]uint32{
		{1},
		{2, 5, 7},
		{9, 15, 20, 23},
		{10, 27, 100, 201, 222},
	}

	for _, item := range items {
		bm := GetBitMap("")
		for _, bit := range item {
			err := bm.SetBit(bit, true)
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

	bm = GetBitMap("")
	for _, bit := range []uint32{10, 27, 100, 201, 222} {
		bm.SetBit(bit, true)
	}

	bm.SetBit(100, false)
	t.Logf("byte_count:%d,raw_str:%s", len([]byte(bm.GetStr())), bm.GetStr())
	if !checkBit(&bm, []uint32{10, 27, 201, 222}) {
		t.Fatalf("expect got bit:%+v", []uint32{10, 27, 201, 222})
	}
	if bm.GetBit(100) {
		t.Fatalf("should not has bit 100")
	}

	bm.SetBit(105, false)
	t.Logf("byte_count:%d,raw_str:%s", len([]byte(bm.GetStr())), bm.GetStr())
	if !checkBit(&bm, []uint32{10, 27, 201, 222}) {
		t.Fatalf("expect got bit:%+v", []uint32{10, 27, 201, 222})
	}
	if bm.GetBit(105) {
		t.Fatalf("should not has bit 105")
	}

	bm.SetBit(222, false)
	t.Logf("byte_count:%d,raw_str:%s", len([]byte(bm.GetStr())), bm.GetStr())
	if !checkBit(&bm, []uint32{10, 27, 201}) {
		t.Fatalf("expect got bit:%+v", []uint32{10, 27, 201, 222})
	}
	if bm.GetBit(222) {
		t.Fatalf("should not has bit 222")
	}

	bm.SetBit(223, false)
	bm.SetBit(221, false)
	bm.SetBit(202, false)
	bm.SetBit(200, false)
	bm.SetBit(400, false)
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

	bm.SetBit(1, true)
	bm.SetBit(7, true)
	bm.SetBit(8, true)
	bm.SetBit(8, false)
	if !checkBit(&bm, []uint32{1, 7, 10, 27, 201}) {
		t.Fatalf("expect got bit:%+v", []uint32{1, 7, 10, 27, 201})
	}
	if bm.GetBit(8) {
		t.Fatalf("should not has bit 8")
	}
	bm.SetBit(7, false)
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
	bm := GetBitMap("a0A")
	expBits := []uint32{1, 6, 7, 13, 14, 17, 23}

	for _, bit := range expBits {
		if !bm.GetBit(bit) {
			t.Fatalf("GetBit %d wrong", bit)
		}
	}
}

func TestStrBitmap_GetStr(t *testing.T) {
	bm := GetBitMap("")
	expBits := []uint32{1, 6, 7, 13, 14, 17, 23}

	for _, bit := range expBits {
		err := bm.SetBit(bit, true)
		if err != nil {
			t.Fatalf("call SetBit got err:%s,bit:%d", err, bit)
		}
	}

	if bm.GetStr() != "a0A" {
		t.Fatalf("call GetStr wrong,got:%s,expect:%s", bm.GetStr(), "a0A")
	}
}
