package japanese

import (
	"fmt"
	"strconv"
	"strings"
)

// 西暦を年月日に分ける
func getWesternDate(westernStrDate string) (int, int, int, error) {

	tmp := strings.Split(westernStrDate, "年")
	year, err := strconv.ParseInt(tmp[0], 10, 32)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("年が数字ではないかもしれません。year[%s]", tmp[0])
	}
	remain := tmp[1]
	tmp = strings.Split(remain, "月")
	month, err := strconv.ParseInt(tmp[0], 10, 32)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("月が数字ではないかもしれません。month[%s]", tmp[0])
	}
	day, err := strconv.ParseInt(strings.Replace(tmp[1], "日", "", 1), 10, 32)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("日が数字ではないかもしれません。day[%s]", tmp[1])
	}

	return int(year), int(month), int(day), nil
}

func ResolveFromWestern(strWesternDate string) (Wareki, error) {

	year, month, day, err := getWesternDate(strWesternDate)

	if err != nil {
		return nil, err
	}

	warekiList := []Wareki{
		NewReiwaFromWestern(year, month, day),
		NewHeiseiFromWestern(year, month, day),
		NewShowaFromWestern(year, month, day),
		NewTaishoFromWestern(year, month, day),
		NewMeijiFromWestern(year, month, day),
	}

	for _, wareki := range warekiList {
		if wareki.Resolve() {
			return wareki, nil
		}
	}

	return nil, fmt.Errorf("対応外の日付です[%s]", strWesternDate)
}
