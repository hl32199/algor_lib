package algor_lib

import (
	"strconv"
	"testing"
)

func TestBitmap_GetBitGetIntv(t *testing.T) {
	bm := GetBitMap(0xf0f)
	if bm.GetIntV() != 0xf0f {
		t.Fatalf("GetIntV wrong,expect %d,got %d", 0xf0f, bm.GetIntV())
	}

	if res, err := bm.GetBit(65); err == nil || err.Error() != "bit position out of range[1,64]" {
		t.Fatalf("expect error:'bit position out of range[1,64]',got err:%+v,res:%+v", err, res)
	}

	items := []bool{true, true, true, true, false, false, false, false, true, true, true, true, false}
	for i, exp := range items {
		if res, err := bm.GetBit(uint32(i + 1)); err != nil || res != exp {
			t.Fatalf("unexpected res,got err:%+v,res:%+v;expect res:%+v", err, res, exp)
		}
	}

	bm = GetBitMap(0)
	if bm.GetIntV() != 0 {
		t.Fatalf("GetIntV wrong,expect %d,got %d", 0, bm.GetIntV())
	}
	items = []bool{false, false, false, false, false, false, false, false, false, false, false, false, false}
	for i, exp := range items {
		if res, err := bm.GetBit(uint32(i + 1)); err != nil || res != exp {
			t.Fatalf("unexpected res,got err:%+v,res:%+v;expect res:%+v", err, res, exp)
		}
	}
}

func TestBitmap_SetBit(t *testing.T) {
	bm := GetBitMap(0xa)
	items := []struct {
		bit  uint32
		intV uint64
	}{
		{2, 0xa}, {4, 0xa}, {1, 0xb}, {4, 0xb}, {6, 0x2b}, {9, 0x12b}, {2, 0x12b},
		{6, 0x12b}, {7, 0x16b},
	}
	for _, item := range items {
		err := bm.SetBit(item.bit)
		if err != nil {
			t.Fatalf("unexpected err:%s", err.Error())
		}
		if bm.GetIntV() != item.intV {
			t.Fatalf("expected %d,got %d", item.intV, bm.GetIntV())
		}
	}
	bm = GetBitMap(0)
	items = []struct {
		bit  uint32
		intV uint64
	}{
		{2, 0x2}, {4, 0xa}, {1, 0xb}, {4, 0xb}, {6, 0x2b}, {9, 0x12b}, {2, 0x12b},
		{6, 0x12b}, {7, 0x16b},
	}
	for _, item := range items {
		err := bm.SetBit(item.bit)
		if err != nil {
			t.Fatalf("unexpected err:%s", err.Error())
		}
		if bm.GetIntV() != item.intV {
			t.Fatalf("expected %d,got %d", item.intV, bm.GetIntV())
		}
	}

	unsetItems := []struct {
		bit  uint32
		intV uint64
	}{
		{6, 0x14b}, {1, 0x14a}, {9, 0x4a}, {4, 0x42}, {1, 0x42}, {9, 0x42}, {2, 0x40},
		{7, 0x0}, {6, 0x0},
	}
	for _, item := range unsetItems {
		err := bm.UnsetBit(item.bit)
		if err != nil {
			t.Fatalf("unexpected err:%s", err.Error())
		}
		if bm.GetIntV() != item.intV {
			t.Fatalf("expected %d,got %d", item.intV, bm.GetIntV())
		}
	}

	if err := bm.SetBit(68); err == nil || err.Error() != "bit position out of range[1,64]" {
		t.Fatalf("expect err:'bit position out of range[1,64]', got %+v", err)
	}
	var bm2 *Bitmap
	if err := bm2.SetBit(68); err == nil || err.Error() != "uninitialized Bitmap" {
		t.Fatalf("expect err:'uninitialized Bitmap', got %+v", err)
	}
}

func TestBitmap_SetBitBatch(t *testing.T) {
	items := [][]uint32{
		{1, 3, 9}, {10, 12}, {11, 3, 10}, {22, 1},
	}
	bm1 := GetBitMap(0)
	bm2 := GetBitMap(0)
	for _, item := range items {
		_ = bm1.SetBitBatch(item)
		for _, bit := range item {
			_ = bm2.SetBit(bit)
		}
		if bm1.GetIntV() != bm2.GetIntV() {
			t.Fatalf("expect %d,got %d", bm2.GetIntV(), bm1.GetIntV())
		}
	}

	unsetItems := [][]uint32{
		{9, 10, 22}, {1, 3, 10}, {11, 12, 3}, {9, 10, 3, 1, 22},
	}
	for _, item := range unsetItems {
		_ = bm1.UnsetBitBatch(item)
		for _, bit := range item {
			_ = bm2.UnsetBit(bit)
		}
		if bm1.GetIntV() != bm2.GetIntV() {
			t.Fatalf("expect %d,got %d", bm2.GetIntV(), bm1.GetIntV())
		}
	}

	var bm3 *Bitmap
	if err := bm3.SetBitBatch([]uint32{1, 4, 90}); err == nil || err.Error() != "uninitialized Bitmap" {
		t.Fatalf("expect err,got %+v", err)
	}
	if err := bm3.UnsetBitBatch([]uint32{1, 4, 90}); err == nil || err.Error() != "uninitialized Bitmap" {
		t.Fatalf("expect err,got %+v", err)
	}
	bm4 := GetBitMap(1)
	if err := bm4.SetBitBatch([]uint32{1, 4, 90}); err == nil || err.Error() != "bit position out of range[1,64]" {
		t.Fatalf("expect err,got %+v", err)
	}
	if err := bm4.UnsetBitBatch([]uint32{1, 4, 90}); err == nil || err.Error() != "bit position out of range[1,64]" {
		t.Fatalf("expect err,got %+v", err)
	}
}

func TestBitmap_GetBits(t *testing.T) {
	items := map[uint64]string{
		0xf0f: "1,2,3,4,9,10,11,12,",
		0xa00: "10,12,",
		0:     "",
	}
	for intv, item := range items {
		bm := GetBitMap(intv)
		bits := bm.GetBits()
		res := ""
		for _, bit := range bits {
			res += strconv.Itoa(int(bit))
			res += ","
		}
		if item != res {
			t.Fatalf("expect %s,got %s", item, res)
		}
	}
}
