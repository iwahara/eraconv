package validate

import (
	"testing"
)

func TestValidateArg_令和(t *testing.T) {
	got := ValidateArg("令和5年5月22日")
	if got != nil {
		t.Errorf("おかしいで")
	}
}

func TestValidateArg_昭和(t *testing.T) {
	got := ValidateArg("昭和5年5月22日")
	if got != nil {
		t.Errorf("おかしいで")
	}
}

func TestValidateArg_平成(t *testing.T) {
	got := ValidateArg("平成5年5月22日")
	if got != nil {
		t.Errorf("おかしいで")
	}
}

func TestValidateArg_大正(t *testing.T) {
	got := ValidateArg("大正5年5月22日")
	if got != nil {
		t.Errorf("おかしいで")
	}
}

func TestValidateArg_明治(t *testing.T) {
	got := ValidateArg("明治5年5月22日")
	if got != nil {
		t.Errorf("おかしいで")
	}
}

func TestValidateArg_西暦(t *testing.T) {
	got := ValidateArg("2020年5月22日")
	if got != nil {
		t.Errorf("おかしいで")
	}
}

func TestValidateArg_不正(t *testing.T) {
	got := ValidateArg("aaa5年5月22日")
	if got == nil {
		t.Errorf("通ってしまったのでおかしい")
	}

}
