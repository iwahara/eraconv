package validate

import (
	"errors"
	"regexp"
)

func ValidateArg(arg string) error {

	regex := regexp.MustCompile(`^(令和|平成|昭和|大正|明治|\d{4})`)

	if regex.MatchString(arg) {
		return nil
	}
	return errors.New("引数の先頭は「令和」「平成」「昭和」「大正」「明治」「数字4桁」である必要があります。")
}
