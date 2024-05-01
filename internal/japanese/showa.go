package japanese

import (
	"fmt"
)

type Showa struct {
	gengo string
	year  int
	month int
	day   int
}

func NewShowaFromWestern(year int, month int, day int) Showa {
	return Showa{gengo: "昭和", year: int(year), month: int(month), day: int(day)}
}

func (r Showa) Resolve() bool {
	// 平成は1926年12月25日から1989年1月7日(昭和64年)まで

	if r.GetYear() < 1 {
		return false
	}
	if r.GetYear() > 64 {
		return false
	}

	if r.GetYear() == 1 && r.GetMonth() < 12 {
		return false
	}
	if r.GetYear() == 1 && r.GetMonth() == 12 && r.GetDay() < 25 {
		return false
	}
	if r.GetYear() == 64 && r.GetMonth() > 1 {
		return false
	}
	if r.GetYear() == 64 && r.GetMonth() == 1 && r.GetDay() > 7 {
		return false
	}

	return true
}

func (r Showa) GetGengo() string {
	return r.gengo
}

func (r Showa) GetYear() int {
	return r.year - 1925
}

func (r Showa) GetMonth() int {
	return r.month
}
func (r Showa) GetDay() int {
	return r.day
}

func (r Showa) ToString() string {
	if r.GetYear() == 1 {
		return fmt.Sprintf("昭和元年%d月%d日", r.month, r.day)
	}
	return fmt.Sprintf("昭和%d年%d月%d日", r.GetYear(), r.GetMonth(), r.GetDay())
}
