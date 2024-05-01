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
	return Showa{gengo: "昭和", year: year - 1925, month: month, day: day}
}

func NewShowaFromWareki(gengo string, year int, month int, day int) Showa {
	return Showa{gengo: gengo, year: year, month: month, day: day}
}

func (r Showa) Resolve() bool {
	// 平成は1926年12月25日から1989年1月7日(昭和64年)まで

	if r.GetGengo() != "昭和" {
		return false
	}
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
	return r.year
}

func (r Showa) GetMonth() int {
	return r.month
}
func (r Showa) GetDay() int {
	return r.day
}

func (r Showa) ToString() string {
	if r.GetYear() == 1 {
		return fmt.Sprintf("%s元年%d月%d日", r.GetGengo(), r.month, r.day)
	}
	return fmt.Sprintf("%s%d年%d月%d日", r.GetGengo(), r.GetYear(), r.GetMonth(), r.GetDay())
}

func (r Showa) ToWesternString() string {
	return fmt.Sprintf("%d年%d月%d日", (r.GetYear() + 1925), r.GetMonth(), r.GetDay())
}
