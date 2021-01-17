package algor_lib

import "errors"

type FeeCalc struct {
	inited     bool
	ConfigList []*ConfigItem
}

func (f *FeeCalc) Init(configList []ConfigItem) error {
	var lastItem ConfigItem
	for _, item := range configList {
		if item.UpperCount == 0 {
			lastItem = item
			continue
		}

		newItem := item
		for k, config := range f.ConfigList {
			if config.UpperCount == newItem.UpperCount {
				//重复的配置覆盖前面的
				*f.ConfigList[k] = newItem
				break
			} else if config.UpperCount > newItem.UpperCount {
				newList := append(f.ConfigList[:k], &newItem)
				f.ConfigList = append(newList, f.ConfigList[k:]...)
				break
			}
		}
		f.ConfigList = append(f.ConfigList, &newItem)
	}

	if lastItem == (ConfigItem{}) {
		return errors.New("lack of last gear config")
	}
	f.ConfigList = append(f.ConfigList, &lastItem)

	//计算速算扣除数
	var tmpTotalAmount float64
	var lowerCount int
	for _, config := range f.ConfigList {
		config.QuickPlus = tmpTotalAmount - (float64(lowerCount) * config.Fee)
		if config.UpperCount > 0 {
			tmpTotalAmount += config.Fee * float64(config.UpperCount-lowerCount)
			lowerCount = config.UpperCount
		}
	}
	f.inited = true
	return nil
}

func GetFeeCalculator(configList []ConfigItem) (*FeeCalc, error) {
	calculator := new(FeeCalc)
	err := calculator.Init(configList)
	if err != nil {
		return nil, errors.New("invalid config list")
	}

	return calculator, nil
}

func (f *FeeCalc) Calc(count int) (float64, error) {
	if f.inited == false {
		return 0, errors.New("FeeCalculator not inited")
	}

	if count == 0 {
		return float64(0), nil
	}

	for _, config := range f.ConfigList {
		if count <= config.UpperCount || config.UpperCount == 0 {
			return (config.Fee*float64(count) + config.QuickPlus), nil
		}
	}
	return 0, errors.New("not find match gear")
}

type ConfigItem struct {
	UpperCount int
	Fee        float64
	QuickPlus  float64
}
