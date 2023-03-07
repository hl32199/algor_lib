package algor_lib

import "errors"

type StrBitmap struct {
	bytes []byte
}

func GetStrBitMap(rawStr string) StrBitmap {
	return StrBitmap{[]byte(rawStr)}
}

//设置比特位
//pos 比特位，从1开始
func (m *StrBitmap) SetBit(pos uint32) error {
	if m == nil {
		return errors.New("uninitialized Bitmap")
	}
	if pos == 0 {
		return errors.New("pos should be greater than 0")
	}

	realPos := pos - 1
	bytePos := getBytePos(realPos)
	bitPos := getBitPos(realPos)

	if bytePos >= uint32(len(m.bytes)) {
		m.bytes = append(m.bytes, make([]byte, bytePos-uint32(len(m.bytes))+1)...)
	}

	m.bytes[bytePos] = m.bytes[bytePos] | (1 << bitPos)
	return nil
}

//取消比特位
//pos 比特位，从1开始
func (m *StrBitmap) UnsetBit(pos uint32) error {
	if m == nil {
		return errors.New("uninitialized Bitmap")
	}
	if pos == 0 {
		return errors.New("pos should be greater than 0")
	}

	realPos := pos - 1
	bytePos := getBytePos(realPos)
	bitPos := getBitPos(realPos)

	if bytePos >= uint32(len(m.bytes)) {
		return nil
	}

	m.bytes[bytePos] = m.bytes[bytePos] & ^(1 << bitPos)
	return nil
}

//查询比特位
//pos 比特位，从1开始
func (m StrBitmap) GetBit(pos uint32) bool {
	if pos == 0 {
		return false
	}

	realPos := pos - 1
	bytePos := getBytePos(realPos)
	if bytePos >= uint32(len(m.bytes)) {
		return false
	}

	if m.bytes[bytePos]&(1<<getBitPos(realPos)) > 0 {
		return true
	}
	return false
}

//查询一个位所在的字节的索引
//realPos 要查询的比特位，realPos从0开始，pos从1开始，realPos = pos -1
//例如，11在第2个字节的第3位，所以bytePos=1
func getBytePos(realPos uint32) uint32 {
	return realPos >> 3
}

//查询一个位在所在的字节当中位的索引，范围为0-7
//realPos 要查询的比特位，realPos从0开始，pos从1开始，realPos = pos -1
//例如，11在第2个字节的第3位，所以bitPos=2
func getBitPos(realPos uint32) uint32 {
	return realPos & 7
}

func (m StrBitmap) GetStr() string {
	return string(m.bytes)
}
