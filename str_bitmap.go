package algor_lib

import "errors"

type StrBitmap struct {
	bytes []byte
}

func GetBitMap(rawStr string) StrBitmap {
	return StrBitmap{[]byte(rawStr)}
}

func (m *StrBitmap) SetBit(pos uint32, hasValue bool) error {
	if m == nil {
		return errors.New("uninitialized Bitmap")
	}
	if pos == 0 {
		return errors.New("pos should be greater than 0")
	}

	bytePos, bitPos := getBytePosBitPos(pos)

	if bytePos >= uint32(len(m.bytes)) {
		if hasValue {
			m.bytes = append(m.bytes, make([]byte, bytePos-uint32(len(m.bytes))+1)...)
		} else {
			return nil
		}
	}

	if hasValue {
		m.bytes[bytePos] = m.bytes[bytePos] | (1 << bitPos)
	} else {
		m.bytes[bytePos] = m.bytes[bytePos] & ^(1 << bitPos)
	}

	return nil
}

func (m StrBitmap) GetBit(pos uint32) bool {
	bytePos, bitPos := getBytePosBitPos(pos)
	if bytePos >= uint32(len(m.bytes)) {
		return false
	}

	if m.bytes[bytePos]&(1<<bitPos) > 0 {
		return true
	}
	return false
}

//pos 要查询的比特位
//bytePos pos所在的字节的索引
//bitPos pos在所在字节的比特位索引
//例如，11在第2个字节的第3位，所以bytePos=1，bitPos=2
func getBytePosBitPos(pos uint32) (bytePos uint32, bitPos uint32) {
	bytePos = pos / 8
	if pos%8 == 0 {
		bytePos--
	}
	bitPos = (pos + 7) % 8
	return
}

func (m StrBitmap) GetStr() string {
	return string(m.bytes)
}
