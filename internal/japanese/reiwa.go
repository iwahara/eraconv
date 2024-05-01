package japanese

import (
	"fmt"
)

type Reiwa struct {
	gengo string
	year  int
	month int
	day   int
}

func NewReiwaFromWestern(year int, month int, day int) Reiwa {
	return Reiwa{gengo: "令和", year: year - 2018, month: month, day: day}
}

func NewReiwaFromWareki(gengo string, year int, month int, day int) Reiwa {
	return Reiwa{gengo: gengo, year: year, month: month, day: day}
}

func (r Reiwa) Resolve() bool {
	// 令和は2019/05/01から

	if r.GetGengo() != "令和" {
		return false
	}

	if r.GetYear() < 1 {
		return false
	}
	if r.GetYear() == 1 && r.GetMonth() < 5 {
		return false
	}

	return true
}

func (r Reiwa) GetGengo() string {
	return r.gengo
}

func (r Reiwa) GetYear() int {
	return r.year
}

func (r Reiwa) GetMonth() int {
	return r.month
}
func (r Reiwa) GetDay() int {
	return r.day
}

func (r Reiwa) ToString() string {
	if r.GetYear() == 1 {
		return fmt.Sprintf("%s元年%d月%d日", r.GetGengo(), r.month, r.day)
	}
	return fmt.Sprintf("%s%d年%d月%d日", r.GetGengo(), r.GetYear(), r.GetMonth(), r.GetDay())
}

func (r Reiwa) ToWesternString() string {
	return fmt.Sprintf("%d年%d月%d日", (r.GetYear() + 2018), r.GetMonth(), r.GetDay())
}
