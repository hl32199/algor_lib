package algor_lib

import "errors"

type Bitmap struct {
	intValue uint64
}

func GetBitMap(rawIntValue uint64) Bitmap {
	return Bitmap{intValue: rawIntValue}
}

//设置比特位
//pos：比特位，从1开始，没有0
func (m *Bitmap) SetBit(pos uint32) error {
	if m == nil {
		return errors.New("uninitialized Bitmap")
	}

	if m.posOutOfRange(pos) {
		return errors.New("bit position out of range[1,64]")
	}

	m.intValue = m.intValue | (1 << (pos - 1))
	return nil
}

//取消比特位
//pos：比特位，从1开始，没有0
func (m *Bitmap) UnsetBit(pos uint32) error {
	if m == nil {
		return errors.New("uninitialized Bitmap")
	}

	if m.posOutOfRange(pos) {
		return errors.New("bit position out of range[1,64]")
	}

	m.intValue = m.intValue & ^(1 << (pos - 1))
	return nil
}

//pos：比特位，从1开始，没有0
func (m Bitmap) GetBit(pos uint32) (bool, error) {
	if m.posOutOfRange(pos) {
		return false, errors.New("bit position out of range[1,64]")
	}

	return (m.intValue & (1 << (pos - 1))) > 0, nil
}

func (m *Bitmap) SetBitBatch(posList []uint32) error {
	if m == nil {
		return errors.New("uninitialized Bitmap")
	}

	v := m.intValue
	for _, pos := range posList {
		if m.posOutOfRange(pos) {
			return errors.New("bit position out of range[1,64]")
		}

		v = v | (1 << (pos - 1))
	}

	m.intValue = v
	return nil
}

func (m *Bitmap) UnsetBitBatch(posList []uint32) error {
	if m == nil {
		return errors.New("uninitialized Bitmap")
	}

	v := m.intValue
	for _, pos := range posList {
		if m.posOutOfRange(pos) {
			return errors.New("bit position out of range[1,64]")
		}

		v = v & ^(1 << (pos - 1))
	}

	m.intValue = v
	return nil
}

func (m Bitmap) GetIntV() uint64 {
	return m.intValue
}

func (m Bitmap) GetBits() []uint32 {
	bits := make([]uint32, 0, 64)
	v := m.intValue
	for i := uint32(1); v > 0; i++ {
		if (v & 1) == 1 {
			bits = append(bits, i)
		}
		v = v >> 1
	}
	return bits
}

func (m Bitmap) posOutOfRange(pos uint32) bool {
	return pos == 0 || pos > 64
}
