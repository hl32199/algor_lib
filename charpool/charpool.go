package charpool

import (
	"fmt"
)

func NewCharPool(pool string) *CharPool {
	return &CharPool{pool: []rune(pool)}
}

type CharPool struct {
	pool []rune
}

func (c *CharPool) AddChars(chars string) {
	charRunes := []rune(chars)
	charMap := make(map[rune]struct{}, len(c.pool)+len(charRunes))
	newUniq := make([]rune, 0, len(charRunes))

	for _, i := range c.pool {
		charMap[i] = struct{}{}
	}

	for _, i := range charRunes {
		if _, ok := charMap[i]; !ok {
			newUniq = append(newUniq, i)
			charMap[i] = struct{}{}
		}
	}

	fmt.Println(string(newUniq))

	newPool := make([]rune, 0, len(charMap))
	for i := range charMap {
		newPool = append(newPool, i)
	}
	fmt.Println(string(newPool))
}
