package base62

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

var (
	baseStr    string
	baseStrLen uint64
)

// 必须调用
func MustInit(bs string) {
	if len(bs) == 0 {
		panic("need base string")
	}
	baseStr = bs
	baseStrLen = uint64(len(bs))
}

// Encode 将十进制整数转换为 62 进制字符串。
func Encode(num uint64) string {
	if num == 0 {
		return "0"
	}

	var buf [11]byte
	index := len(buf)

	for num > 0 {
		index--
		buf[index] = baseStr[num%62]
		num /= 62
	}

	return string(buf[index:])
}

// Decode 将 62 进制字符串转换为十进制整数。
func Decode(value string) (uint64, error) {
	if value == "" {
		return 0, errors.New("base62 value cannot be empty")
	}

	var result uint64

	for i := 0; i < len(value); i++ {
		index := strings.IndexByte(baseStr, value[i])
		if index < 0 {
			return 0, fmt.Errorf(
				"invalid base62 character %q at index %d",
				value[i],
				i,
			)
		}

		digit := uint64(index)

		// 防止 result*62+digit 超出 uint64 范围。
		if result > (math.MaxUint64-digit)/62 {
			return 0, errors.New("base62 value overflows uint64")
		}

		result = result*62 + digit
	}

	return result, nil
}
