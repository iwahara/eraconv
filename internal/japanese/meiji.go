package japanese

import (
	"fmt"
)

type Meiji struct {
	gengo string
	year  int
	month int
	day   int
}

func NewMeijiFromWestern(year int, month int, day int) Meiji {
	return Meiji{gengo: "明治", year: year - 1867, month: month, day: day}
}

func NewMeijiFromWareki(gengo string, year int, month int, day int) Meiji {
	return Meiji{gengo: gengo, year: year, month: month, day: day}
}

func (r Meiji) Resolve() bool {
	// 明治は1873年1月1日から1912年7月29日(明治45年)までだが
	// 明治６年以前は太陰暦で単純変換できないので、いったん対象を明治６年からにしている

	if r.GetGengo() != "明治" {
		return false
	}
	if r.GetYear() < 6 {
		return false
	}
	if r.GetYear() > 45 {
		return false
	}

	if r.GetYear() == 45 && r.GetMonth() > 7 {
		return false
	}
	if r.GetYear() == 45 && r.GetMonth() == 7 && r.GetDay() > 29 {
		return false
	}

	return true
}

func (r Meiji) GetGengo() string {
	return r.gengo
}

func (r Meiji) GetYear() int {
	return r.year
}

func (r Meiji) GetMonth() int {
	return r.month
}
func (r Meiji) GetDay() int {
	return r.day
}

func (r Meiji) ToString() string {
	if r.GetYear() == 1 {
		return fmt.Sprintf("%s元年%d月%d日", r.GetGengo(), r.month, r.day)
	}
	return fmt.Sprintf("%s%d年%d月%d日", r.GetGengo(), r.GetYear(), r.GetMonth(), r.GetDay())
}

func (r Meiji) ToWesternString() string {
	return fmt.Sprintf("%d年%d月%d日", (r.GetYear() + 1867), r.GetMonth(), r.GetDay())
}
