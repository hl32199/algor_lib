package main

import (
	"fmt"
	"math/rand"
	"time"
)

var fundPool int64

type temaBet struct {
	number int
	amount int64
}

type pinmaBet struct {
	shengxiao int
	amount    int64
}

type bets struct {
	temaBets  []temaBet
	pinmaBets []pinmaBet
}

func main() {
	for i := 0; i < 1000; i++ {
		play(genBets())
	}
	fmt.Printf("奖池余：%d", fundPool)
}

func genBets() bets {
	rand.Seed(time.Now().UnixNano())

	temabets := make([]temaBet, 0, 1000)
	for i := 0; i < 1000; i++ {
		temabets = append(temabets, temaBet{
			number: rand.Intn(49) + 1,
			amount: 10,
		})
	}

	//pinmabets := make([]pinmaBet, 0, 100)
	//for i := 0; i < 100; i++ {
	//	pinmabets = append(pinmabets, pinmaBet{
	//		shengxiao: rand.Intn(12),
	//		amount:    200,
	//	})
	//}

	return bets{
		temaBets: temabets,
		//pinmaBets: pinmabets,
	}
}

func play(bets bets) {
	rand.Seed(time.Now().UnixNano())

	numberPool := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29,
		30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49}

	teIdx := rand.Intn(49)
	tema := numberPool[teIdx]
	numberPool = append(numberPool[:teIdx], numberPool[teIdx+1:]...)
	//fmt.Printf("数字池：%+v,特码：%d\r\n", numberPool, tema)

	shengxiaoMap := make(map[int]struct{}, 6)
	for i := 0; i < 6; i++ {
		shIdx := rand.Intn(len(numberPool))
		shengxiaoma := numberPool[shIdx]
		numberPool = append(numberPool[:shIdx], numberPool[shIdx+1:]...)
		//fmt.Printf("数字池：%+v,生肖码：%d\r\n", numberPool, shengxiaoma)
		shengxiaoMap[shengxiaoma%12] = struct{}{}
	}
	fmt.Printf("生肖码：%+v\r\n", shengxiaoMap)

	for _, tb := range bets.temaBets {
		if tb.number == tema {
			fundPool -= tb.amount * 41
			//fmt.Printf("特码：%d,中了，奖池 -%d,余%d\r\n", tb.number, tb.amount*42, fundPool)
		} else {
			fundPool += tb.amount
			//fmt.Printf("特码：%d,没中，奖池 +%d,余%d\r\n", tb.number, tb.amount, fundPool)
		}
	}

	for _, pb := range bets.pinmaBets {
		if _, ok := shengxiaoMap[pb.shengxiao]; ok {
			if pb.shengxiao == 0 {
				fundPool -= int64(float64(pb.amount) * 0.8)
				//fmt.Printf("品码：%d,中了，奖池 -%d,余%d\r\n", pb.shengxiao, int64(float64(pb.amount)*0.8), fundPool)
			} else {
				fundPool -= pb.amount
				//fmt.Printf("品码：%d,中了，奖池 -%d,余%d\r\n", pb.shengxiao, pb.amount, fundPool)
			}
		} else {
			fundPool += pb.amount
			//fmt.Printf("品码：%d,没中，奖池 +%d,余%d\r\n", pb.shengxiao, pb.amount, fundPool)
		}
	}
}
