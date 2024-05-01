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
	return Taisho{gengo: "大正", year: year - 1911, month: month, day: day}
}

func NewTaishoFromWareki(gengo string, year int, month int, day int) Taisho {
	return Taisho{gengo: gengo, year: year, month: month, day: day}
}

func (r Taisho) Resolve() bool {
	// 大正は1912年7月30日から1926年12月24日(大正15年)まで

	if r.GetGengo() != "大正" {
		return false
	}
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
	return r.year
}

func (r Taisho) GetMonth() int {
	return r.month
}
func (r Taisho) GetDay() int {
	return r.day
}

func (r Taisho) ToString() string {
	if r.GetYear() == 1 {
		return fmt.Sprintf("%s元年%d月%d日", r.GetGengo(), r.month, r.day)
	}
	return fmt.Sprintf("%s%d年%d月%d日", r.GetGengo(), r.GetYear(), r.GetMonth(), r.GetDay())
}

func (r Taisho) ToWesternString() string {
	return fmt.Sprintf("%d年%d月%d日", (r.GetYear() + 1911), r.GetMonth(), r.GetDay())
}
