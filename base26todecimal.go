package algor_lib

import (
	"bytes"
	"strconv"
)

//26进制字符串转换为十进制字符串
func Base26toDecimal(source string) string {
	temp := make(IntMap, 5, 5)
	temp[0] = 1
	sourceLen := len(source)
	result := make(IntMap, 10, 10)
	for i := sourceLen - 1; i >= 0; i-- {
		num := uint(source[i] - 97)

		result = result.Add(temp.Multi(num))
		temp = temp.Multi(26)
	}

	length := len(result)
	resBuf := bytes.NewBuffer([]byte{})
	var maxInd int
	for i := length - 1; i >= 0; i-- {
		if result[i] > 0 {
			maxInd = i
			break
		}
	}

	for i := maxInd; i >= 0; i-- {
		s := strconv.Itoa(int(result[i]))
		resBuf.WriteString(s)
	}
	return resBuf.String()
}

//数组形式表示的十进制数
type IntMap []uint8

//intMap的加法
func (ii IntMap) Add(j IntMap) IntMap {
	length := len(j)
	length2 := len(ii)
	var maxLen int
	if length > length2 {
		maxLen = length
	} else {
		maxLen = length2
	}
	i := make(IntMap, maxLen+1, maxLen+1)
	copy(i, ii)

	//fmt.Println(i)
	for k := 0; k < length; k++ {
		tmp := j[k] + i[k]
		if tmp < 10 {
			i[k] = tmp
		} else {
			i[k] = tmp - 10
			i[k+1] = i[k+1] + 1
		}
		//fmt.Println(i)
	}

	return i
}

//intMap的乘法
func (ii IntMap) Multi(j uint) IntMap {
	length := len(ii)
	i := make(IntMap, length, length)
	//copy(i,ii)

	for k := 0; k < length; k++ {
		var tmp uint
		tmp = uint(ii[k]) * j
		for m := k; tmp > 0; m++ {
			if len(i) < m+1 {
				i = append(i, make(IntMap, 5, 5)...)
			}
			tmp = tmp + uint(i[m])
			i[m] = uint8(tmp % 10)
			tmp = tmp / 10
		}
	}

	return i
}
