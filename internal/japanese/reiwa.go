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
	return Reiwa{gengo: "令和", year: int(year), month: int(month), day: int(day)}
}

func (r Reiwa) Resolve() bool {
	// 令和は2019/05/01から

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
	return r.year - 2018
}

func (r Reiwa) GetMonth() int {
	return r.month
}
func (r Reiwa) GetDay() int {
	return r.day
}

func (r Reiwa) ToString() string {
	if r.GetYear() == 1 {
		return fmt.Sprintf("令和元年%d月%d日", r.month, r.day)
	}
	return fmt.Sprintf("令和%d年%d月%d日", r.GetYear(), r.GetMonth(), r.GetDay())
}
