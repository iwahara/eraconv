package japanese

import (
	"fmt"
)

type Heisei struct {
	gengo string
	year  int
	month int
	day   int
}

func NewHeisei(year int, month int, day int) Heisei {
	return Heisei{gengo: "平成", year: int(year), month: int(month), day: int(day)}
}

func (r Heisei) Resolve() bool {
	// 平成は1989年1月8日から2019年4月30日(平成31年)まで

	if r.GetYear() < 1 {
		return false
	}
	if r.GetYear() > 31 {
		return false
	}
	if r.GetYear() == 1 && r.GetMonth() == 1 && r.GetDay() < 8 {
		return false
	}
	if r.GetYear() == 31 && r.GetMonth() >= 5 {
		return false
	}

	return true
}

func (r Heisei) GetGengo() string {
	return r.gengo
}

func (r Heisei) GetYear() int {
	return r.year - 1988
}

func (r Heisei) GetMonth() int {
	return r.month
}
func (r Heisei) GetDay() int {
	return r.day
}

func (r Heisei) ToString() string {
	if r.GetYear() == 1 {
		return fmt.Sprintf("平成元年%d月%d日", r.month, r.day)
	}
	return fmt.Sprintf("平成%d年%d月%d日", r.GetYear(), r.GetMonth(), r.GetDay())
}
