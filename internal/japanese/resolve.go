package japanese

import (
	"fmt"
	"regexp"
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

func getWarekiDate(strWarekiDate string) (string, int, int, int, error) {
	reEra := regexp.MustCompile(`^(令和|平成|昭和|大正|明治)`)
	reYear := regexp.MustCompile(`(\d+|元)年`)
	reMonth := regexp.MustCompile(`(\d+)月`)
	reDay := regexp.MustCompile(`(\d+)日`)

	eraMatch := reEra.FindStringSubmatch(strWarekiDate)
	yearMatch := reYear.FindStringSubmatch(strWarekiDate)
	monthMatch := reMonth.FindStringSubmatch(strWarekiDate)
	dayMatch := reDay.FindStringSubmatch(strWarekiDate)

	if len(eraMatch) < 2 {
		return "", 0, 0, 0, fmt.Errorf("元号が見つかりませんでした。入力された文字列[%s]", strWarekiDate)
	}

	if len(yearMatch) < 2 {
		return "", 0, 0, 0, fmt.Errorf("年が見つかりませんでした。入力された文字列[%s]", strWarekiDate)
	}

	if len(monthMatch) < 2 {
		return "", 0, 0, 0, fmt.Errorf("月が見つかりませんでした。入力された文字列[%s]", strWarekiDate)
	}

	if len(dayMatch) < 2 {
		return "", 0, 0, 0, fmt.Errorf("日が見つかりませんでした。入力された文字列[%s]", strWarekiDate)
	}

	year, err := strconv.ParseInt(yearMatch[1], 10, 32)
	if err != nil {
		if yearMatch[1] != "元" {
			return "", 0, 0, 0, fmt.Errorf("年が数字ではないかもしれません。year[%s]", yearMatch[1])
		}
		year = 1
	}

	month, err := strconv.ParseInt(monthMatch[1], 10, 32)
	if err != nil {
		return "", 0, 0, 0, fmt.Errorf("月が数字ではないかもしれません。month[%s]", monthMatch[1])
	}

	day, err := strconv.ParseInt(dayMatch[1], 10, 32)
	if err != nil {
		return "", 0, 0, 0, fmt.Errorf("日が数字ではないかもしれません。day[%s]", dayMatch[1])
	}

	return eraMatch[1], int(year), int(month), int(day), nil
}

func ResolveFromWareki(strWarekiDate string) (Wareki, error) {

	gengo, year, month, day, err := getWarekiDate(strWarekiDate)

	if err != nil {
		return nil, err
	}

	warekiList := []Wareki{
		NewReiwaFromWareki(gengo, year, month, day),
		NewHeiseiFromWareki(gengo, year, month, day),
		NewShowaFromWareki(gengo, year, month, day),
		NewTaishoFromWareki(gengo, year, month, day),
		NewMeijiFromWareki(gengo, year, month, day),
	}

	for _, wareki := range warekiList {
		if wareki.Resolve() {
			return wareki, nil
		}
	}

	return nil, fmt.Errorf("対応外の日付です[%s]", strWarekiDate)
}
