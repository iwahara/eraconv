package japanese

import (
	"fmt"
)

type Taisho struct {
	gengo string
	year  int
	month int
	day   int
}

func NewTaishoFromWestern(year int, month int, day int) Taisho {
	return Taisho{gengo: "大正", year: int(year), month: int(month), day: int(day)}
}

func (r Taisho) Resolve() bool {
	// 大正は1912年7月30日から1926年12月24日(大正15年)まで

	if r.GetYear() < 1 {
		return false
	}
	if r.GetYear() > 15 {
		return false
	}
	if r.GetYear() == 1 && r.GetMonth() < 7 {
		return false
	}
	if r.GetYear() == 1 && r.GetMonth() == 7 && r.GetDay() < 30 {
		return false
	}
	if r.GetYear() == 15 && r.GetMonth() == 12 && r.GetDay() > 24 {
		return false
	}

	return true
}

func (r Taisho) GetGengo() string {
	return r.gengo
}

func (r Taisho) GetYear() int {
	return r.year - 1911
}

func (r Taisho) GetMonth() int {
	return r.month
}
func (r Taisho) GetDay() int {
	return r.day
}

func (r Taisho) ToString() string {
	if r.GetYear() == 1 {
		return fmt.Sprintf("大正元年%d月%d日", r.month, r.day)
	}
	return fmt.Sprintf("大正%d年%d月%d日", r.GetYear(), r.GetMonth(), r.GetDay())
}
