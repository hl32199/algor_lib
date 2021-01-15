package algor_lib

type FeeCalc struct {
	inited     bool
	configList []configItem
}

func GetFeeCalculator(configList []configItem) *FeeCalc {
	for _, item := range configList {

	}
}

func (f *FeeCalc) Calc(count int) float64 {

}

type configItem struct {
	UpperCount int
	Fee        float64
	QuickMinus float64
}
