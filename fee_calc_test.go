package algor_lib

import (
	"testing"
)

func TestGetFeeCalculator(t *testing.T) {
	configList := []ConfigItem{
		{UpperCount: 5, Fee: 30.00},
		{UpperCount: 20, Fee: 15.00},
		{UpperCount: 50, Fee: 10.00},
		{UpperCount: 100, Fee: 9.00},
		{UpperCount: 500, Fee: 8.00},
		{UpperCount: 1000, Fee: 7.00},
		{UpperCount: 2000, Fee: 6.00},
		{UpperCount: 3000, Fee: 5.00},
		{UpperCount: 4000, Fee: 4.00},
		{UpperCount: 5000, Fee: 3.00},
		{UpperCount: 6000, Fee: 2.00},
		{UpperCount: 0, Fee: 1.00},
	}

	calculator, err := GetFeeCalculator(configList)
	if err != nil {
		t.Fatalf("get fee calculator error,:%s", err)
	}
	t.Logf("config list:%#v", calculator.ConfigList[0])
	t.Logf("config list:%#v", calculator.ConfigList[1])
	t.Logf("config list:%#v", calculator.ConfigList[11])
	if len(calculator.ConfigList) != 12 {
		t.Fatalf("wrong config list length,lentgh:%d,should:%d", len(calculator.ConfigList), 12)
	}
	config0 := ConfigItem{UpperCount: 5, Fee: 30.00, QuickPlus: 0}
	config1 := ConfigItem{UpperCount: 20, Fee: 15.00, QuickPlus: 75}
	config2 := ConfigItem{UpperCount: 50, Fee: 10.00, QuickPlus: 175}
	config11 := ConfigItem{UpperCount: 0, Fee: 1.00, QuickPlus: 21825}
	if *calculator.ConfigList[0] != config0 {
		t.Fatalf("wrong config item 0,config:%#v,should:%#v", calculator.ConfigList[0], config0)
	}
	if *calculator.ConfigList[1] != config1 {
		t.Fatalf("wrong config item 1,config:%#v,should:%#v", calculator.ConfigList[1], config1)
	}
	if *calculator.ConfigList[2] != config2 {
		t.Fatalf("wrong config item 2,config:%#v,should:%#v", calculator.ConfigList[2], config2)
	}
	if *calculator.ConfigList[11] != config11 {
		t.Fatalf("wrong config item 11,config:%#v,should:%#v", calculator.ConfigList[11], config11)
	}
}

func TestGetFeeCalc(t *testing.T) {
	configList := []ConfigItem{
		{UpperCount: 5, Fee: 30.00},
		{UpperCount: 20, Fee: 15.00},
		{UpperCount: 50, Fee: 10.00},
		{UpperCount: 100, Fee: 9.00},
		{UpperCount: 500, Fee: 8.00},
		{UpperCount: 1000, Fee: 7.00},
		{UpperCount: 2000, Fee: 6.00},
		{UpperCount: 3000, Fee: 5.00},
		{UpperCount: 4000, Fee: 4.00},
		{UpperCount: 5000, Fee: 3.00},
		{UpperCount: 6000, Fee: 2.00},
		{UpperCount: 0, Fee: 1.00},
	}

	calculator, err := GetFeeCalculator(configList)
	if err != nil {
		t.Fatalf("get fee calculator error,:%s", err)
	}

	var fee float64
	fee, _ = calculator.Calc(2)
	if fee != 60 {
		t.Fatalf("calc error,result:%.2f,should:%d", fee, 60)
	}
	fee, _ = calculator.Calc(20)
	if fee != 375 {
		t.Fatalf("calc error,result:%.2f,should:%d", fee, 375)
	}
	fee, _ = calculator.Calc(7500)
	if fee != 29325 {
		t.Fatalf("calc error,result:%.2f,should:%d", fee, 29325)
	}
}
