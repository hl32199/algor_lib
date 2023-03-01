package algor_lib

import (
	"fmt"
	"math/rand"
	"time"
)

const CardNumber = 52

func min(a, b uint64) uint64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b uint64) uint64 {
	if a > b {
		return a
	} else {
		return b
	}
}

type stg struct {
	name string
	f    stgf
}

type stgf func(great, little uint32, p *player, j *Jiazi) (callAmount uint64)

var stg1 = stg{
	name: "点数达6就全要",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		if great-little >= 6 {
			return j.poolAmount
		}
		return 0
	},
}
var stg2 = stg{
	name: "点数达7就全要",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		if great-little >= 7 {
			return j.poolAmount
		}
		return 0
	},
}
var stg3 = stg{
	name: "点数达8就全要",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		if great-little >= 8 {
			return j.poolAmount
		}
		return 0
	},
}
var stg4 = stg{
	name: "点数达9就全要",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		if great-little >= 9 {
			return j.poolAmount
		}
		return 0
	},
}
var stg5 = stg{
	name: "点数达10就全要",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		if great-little >= 10 {
			return j.poolAmount
		}
		return 0
	},
}
var stg6 = stg{
	name: "点数达8就全要,留底",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		if great-little >= 8 {
			return j.poolAmount - 1
		}
		return 0
	},
}
var stg7 = stg{
	name: "点数8要1/3,9要1/2,更大全要",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		switch great - little {
		case 8:
			return j.poolAmount / 3
		case 9:
			return j.poolAmount / 2
		case 10, 11, 12:
			return j.poolAmount
		default:
			return 0
		}
	},
}
var stg8 = stg{
	name: "点数8要一半但不少于2,更大全要",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		switch great - little {
		case 8:
			return max(2, j.poolAmount/2)
		case 9, 10, 11, 12:
			return j.poolAmount
		default:
			return 0
		}
	},
}
var stg9 = stg{
	name: "7要2,8要4,9要6,10要一半但最多10,更大全要",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		switch great - little {
		case 7:
			return 2
		case 8:
			return 4
		case 9:
			return 6
		case 10:
			return max(10, j.poolAmount/2)
		case 11, 12:
			return j.poolAmount
		default:
			return 0
		}
	},
}
var stg10 = stg{
	name: "6要2,7要4,8要6,9要一半但最多10,更大全要",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		switch great - little {
		case 6:
			return 2
		case 7:
			return 4
		case 8:
			return 6
		case 9:
			return max(10, j.poolAmount/2)
		case 10, 11, 12:
			return j.poolAmount
		default:
			return 0
		}
	},
}
var stg11 = stg{
	name: "点数达8全要，最大40",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		if great-little >= 8 {
			return 40
		}
		return 0
	},
}
var stg12 = stg{
	name: "点数达8全要，最大40,赢超80则最大拿赢的一半",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		if great-little >= 8 {
			return max(40, uint64(p.walletInOut)/2)
		}
		return 0
	},
}
var stg13 = stg{
	name: "点数达8全要，赢不超20只要10，赢不超40按赢数拿，赢超40拿40，赢超80那赢数一半",
	f: func(great, little uint32, p *player, j *Jiazi) (callAmount uint64) {
		if great-little >= 8 {
			if p.walletInOut < 5 {
				return 2
			}
			if p.walletInOut <= 40 {
				return uint64(p.walletInOut)
			}
			return max(40, uint64(p.walletInOut)/2)
		}
		return 0
	},
}

type player struct {
	name        string
	walletInOut int64
	stg         stg
	next        *player
}

type Jiazi struct {
	playerHead  *player
	clearCards  []uint32
	dirtyCards  []uint32
	dealIdx     uint32
	poolAmount  uint64
	printDetail bool
	r           *rand.Rand
}

func (j *Jiazi) initCards() {
	if j.r == nil {
		j.r = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	j.clearCards = make([]uint32, CardNumber)
	j.dealIdx = CardNumber
	j.dirtyCards = make([]uint32, 0, CardNumber)
	for i := uint32(1); i <= 13; i++ {
		j.dirtyCards = append(j.dirtyCards, i, i, i, i)
	}
	j.shuffle(6)
}

func (j *Jiazi) shuffle(times uint32) {
	copy(j.clearCards, j.clearCards[j.dealIdx:])
	j.clearCards = j.clearCards[:CardNumber-j.dealIdx]
	j.dealIdx = 0

	for i := uint32(0); i < times; i++ {
		j.shuffleRound()
		j.shuffleCut()
	}
	j.clearCards = append(j.clearCards, j.dirtyCards...)
	j.dirtyCards = make([]uint32, 0, CardNumber)
}

func (j *Jiazi) shuffleRound() {
	n := len(j.dirtyCards)
	part1 := make([]uint32, n/2)
	part2 := make([]uint32, n-(n/2))
	copy(part1, j.dirtyCards[:n/2])
	copy(part2, j.dirtyCards[n/2:])
	j.dirtyCards = nil

	var i1, i2, nxt int
	for i1 < len(part1) || i2 < len(part2) {
		if i1 < len(part1) {
			nxt = i1 + j.r.Intn(5) + 1
			if nxt > len(part1) {
				nxt = len(part1)
			}
			j.dirtyCards = append(j.dirtyCards, part1[i1:nxt]...)
			i1 = nxt
		}

		if i2 < len(part2) {
			nxt = i2 + j.r.Intn(5) + 1
			if nxt > len(part2) {
				nxt = len(part2)
			}
			j.dirtyCards = append(j.dirtyCards, part2[i2:nxt]...)
			i2 = nxt
		}
	}
}

func (j *Jiazi) shuffleCut() {
	cutIdx := 10 + j.r.Intn(33)
	part1 := make([]uint32, 0, cutIdx)
	part1 = append(part1, j.dirtyCards[:cutIdx]...)
	copy(j.dirtyCards, j.dirtyCards[cutIdx:])
	copy(j.dirtyCards[len(j.dirtyCards)-cutIdx:], part1)
}

func (j *Jiazi) play(round uint32) {
	if j.countPlayer() < 2 {
		fmt.Println("少于2人玩不了")
		return
	}
	j.initCards()

	p := j.playerHead
	for ; round > 0; round-- {
		j.playerPlay(p)
		if p.next == nil {
			p = j.playerHead
		} else {
			p = p.next
		}
	}

	//对玩家输赢排序后输出
	fmt.Println("====================================================")
	cnt := int(j.countPlayer())
	rank := make([]*player, 0, cnt)

	p = j.playerHead
	for p != nil {
		rank = append(rank, p)
		p = p.next
	}

	for i := 0; i < cnt-1; i++ {
		for j := 0; j < cnt-i-1; j++ {
			if rank[j].walletInOut < rank[j+1].walletInOut {
				tmp := rank[j]
				rank[j] = rank[j+1]
				rank[j+1] = tmp
			}
		}
	}

	for _, pyr := range rank {
		fmt.Printf("玩家%s %d;策略：%s\r\n", pyr.name, pyr.walletInOut, pyr.stg.name)
	}

	fmt.Printf("资金池：%d\r\n", j.poolAmount)
}

func (j *Jiazi) countPlayer() uint32 {
	var pyrcnt uint32
	pyr := j.playerHead
	for {
		if pyr == nil {
			break
		}
		pyr = pyr.next
		pyrcnt++
	}
	return pyrcnt
}

func (j *Jiazi) AddPlayer(pyr *player) {
	if j.playerHead == nil {
		j.playerHead = pyr
		return
	}

	tailplayer := j.playerHead
	for {
		if tailplayer.next == nil {
			tailplayer.next = pyr
			break
		}
		tailplayer = tailplayer.next
	}
}

func (j *Jiazi) playerPlay(pyr *player) {
	//牌不够了，重新洗牌
	if CardNumber-j.dealIdx < 3 {
		j.shuffle(3)
	}

	//下底
	if j.poolAmount == 0 {
		p := j.playerHead
		for {
			p.walletInOut--
			j.poolAmount++
			if p.next == nil {
				break
			}
			p = p.next
		}
		if j.printDetail {
			fmt.Printf("下底，资金池：%d\r\n", j.poolAmount)
		}
	}

	//起牌
	g, l := j.get2Cards()
	callAmount := pyr.stg.f(g, l, pyr, j)
	callAmount = min(callAmount, j.poolAmount)

	if callAmount > 0 {
		card3 := j.dealCard()
		if card3 < g && card3 > l {
			pyr.walletInOut += int64(callAmount)
			j.poolAmount -= callAmount
			if j.printDetail {
				fmt.Printf("玩家%s赢%d，总输赢：%d，资金池：%d\r\n", pyr.name, callAmount, pyr.walletInOut, j.poolAmount)
			}
		} else {
			pyr.walletInOut -= int64(callAmount)
			j.poolAmount += callAmount
			if j.printDetail {
				fmt.Printf("玩家%s输%d，总输赢：%d，资金池：%d\r\n", pyr.name, callAmount, pyr.walletInOut, j.poolAmount)
			}
		}
		j.dirtyCards = append(j.dirtyCards, card3)
	} else {
		if j.printDetail {
			fmt.Printf("玩家%s放弃\r\n", pyr.name)
		}
	}
	j.dirtyCards = append(j.dirtyCards, g, l)
}

func (j *Jiazi) get2Cards() (great, little uint32) {
	card1 := j.dealCard()
	card2 := j.dealCard()
	if card1 >= card2 {
		return card1, card2
	}
	return card2, card1
}

func (j *Jiazi) dealCard() uint32 {
	card := j.clearCards[j.dealIdx]
	j.dealIdx++
	return card
}
