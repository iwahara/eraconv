package cmd

func ExamplePrint_明治対象外_西暦から和暦() {

	params := []string{"1872年12月1日"}
	run(params)

	// Output:
	//
}

func ExamplePrint_明治開始_西暦から和暦() {

	params := []string{"1873年1月1日"}
	run(params)

	// Output:
	// 明治6年1月1日
}

func ExamplePrint_明治中_西暦から和暦() {

	params := []string{"1910年1月1日"}
	run(params)

	// Output:
	// 明治43年1月1日
}

func ExamplePrint_明治終了_西暦から和暦() {

	params := []string{"1912年7月29日"}
	run(params)

	// Output:
	// 明治45年7月29日
}

func ExamplePrint_大正開始_西暦から和暦() {

	params := []string{"1912年7月30日"}
	run(params)

	// Output:
	// 大正元年7月30日
}

func ExamplePrint_大正中_西暦から和暦() {

	params := []string{"1920年1月1日"}
	run(params)

	// Output:
	// 大正9年1月1日
}

func ExamplePrint_大正終了_西暦から和暦() {

	params := []string{"1926年12月24日"}
	run(params)

	// Output:
	// 大正15年12月24日
}

func ExamplePrint_昭和開始_西暦から和暦() {

	params := []string{"1926年12月25日"}
	run(params)

	// Output:
	// 昭和元年12月25日
}

func ExamplePrint_昭和中_西暦から和暦() {

	params := []string{"1985年1月1日"}
	run(params)

	// Output:
	// 昭和60年1月1日
}

func ExamplePrint_昭和終了_西暦から和暦() {

	params := []string{"1989年1月7日"}
	run(params)

	// Output:
	// 昭和64年1月7日
}

func ExamplePrint_平成開始_西暦から和暦() {

	params := []string{"1989年1月8日"}
	run(params)

	// Output:
	// 平成元年1月8日
}

func ExamplePrint_平成中_西暦から和暦() {

	params := []string{"2000年4月30日"}
	run(params)

	// Output:
	// 平成12年4月30日
}

func ExamplePrint_平成終了_西暦から和暦() {

	params := []string{"2019年4月30日"}
	run(params)

	// Output:
	// 平成31年4月30日
}

func ExamplePrint_令和開始_西暦から和暦() {

	params := []string{"2019年5月1日"}
	run(params)

	// Output:
	// 令和元年5月1日
}

func ExamplePrint_令和中_西暦から和暦() {

	params := []string{"2022年1月1日"}
	run(params)

	// Output:
	// 令和4年1月1日
}

// --- ここから和暦から西暦のテスト

func ExamplePrint_明治対象外_和暦から西暦() {

	params := []string{"明治5年12月1日"}
	run(params)

	// Output:
	//
}

func ExamplePrint_明治開始_和暦から西暦() {

	params := []string{"明治6年1月1日"}
	run(params)

	// Output:
	// 1873年1月1日
}

func ExamplePrint_明治中_和暦から西暦() {

	params := []string{"明治43年1月1日"}
	run(params)

	// Output:
	// 1910年1月1日
}

func ExamplePrint_明治終了_和暦から西暦() {

	params := []string{"明治45年7月29日"}
	run(params)

	// Output:
	// 1912年7月29日
}

func ExamplePrint_大正開始_和暦から西暦() {

	params := []string{"大正元年7月30日"}
	run(params)

	// Output:
	// 1912年7月30日
}

func ExamplePrint_大正中_和暦から西暦() {

	params := []string{"大正9年1月1日"}
	run(params)

	// Output:
	// 1920年1月1日
}

func ExamplePrint_大正終了_和暦から西暦() {

	params := []string{"大正15年12月24日"}
	run(params)

	// Output:
	// 1926年12月24日
}

func ExamplePrint_昭和開始_和暦から西暦() {

	params := []string{"昭和元年12月25日"}
	run(params)

	// Output:
	// 1926年12月25日
}

func ExamplePrint_昭和中_和暦から西暦() {

	params := []string{"昭和60年1月1日"}
	run(params)

	// Output:
	// 1985年1月1日
}

func ExamplePrint_昭和終了_和暦から西暦() {

	params := []string{"昭和64年1月7日"}
	run(params)

	// Output:
	// 1989年1月7日
}

func ExamplePrint_平成開始_和暦から西暦() {

	params := []string{"平成元年1月8日"}
	run(params)

	// Output:
	// 1989年1月8日
}

func ExamplePrint_平成中_和暦から西暦() {

	params := []string{"平成12年4月30日"}
	run(params)

	// Output:
	// 2000年4月30日
}

func ExamplePrint_平成終了_和暦から西暦() {

	params := []string{"平成31年4月30日"}
	run(params)

	// Output:
	// 2019年4月30日
}

func ExamplePrint_令和開始_和暦から西暦() {

	params := []string{"令和元年5月1日"}
	run(params)

	// Output:
	// 2019年5月1日
}

func ExamplePrint_令和中_和暦から西暦() {

	params := []string{"令和4年1月1日"}
	run(params)

	// Output:
	// 2022年1月1日
}
