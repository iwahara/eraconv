/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/iwahara/eraconv/internal/japanese"
	"github.com/iwahara/eraconv/internal/validate"
	"github.com/spf13/cobra"
)

func outputToday() error {
	western := time.Now().Format("2006年1月2日")
	messageWestern := fmt.Sprintf("西暦：%s\n", western)
	fmt.Print(messageWestern)

	japanese, err := japanese.ResolveFromWestern(western)
	if err != nil {
		return err
	}
	fmt.Printf("和暦：%s", japanese.ToString())
	fmt.Println()
	return nil
}

func customArgValidation(cmd *cobra.Command, args []string) error {

	if len(args) < 1 {
		return nil
	}

	input := args[0]
	if err := validate.ValidateArg(input); err != nil {
		return err
	}
	return nil
}
func isWareki(arg string) bool {

	regex := regexp.MustCompile(`^(令和|平成|昭和|大正|明治)`)

	if regex.MatchString(arg) {
		return true
	}
	return false
}

func outputDate(arg string) error {
	if isWareki(arg) {
		//和暦から西暦へ変換する
		wareki, err := japanese.ResolveFromWareki(arg)

		if err != nil {
			return err
		}

		fmt.Println(wareki.ToWesternString())

	} else {
		//西暦から和暦へ変換する（バリデーションで不正なものはすべて弾いているので）
		wareki, err := japanese.ResolveFromWestern(arg)

		if err != nil {
			return err
		}

		fmt.Println(wareki.ToString())
	}
	return nil
}

func run(args []string) error {
	if len(args) < 1 {
		outputToday()
		return nil
	}
	arg := args[0]
	err := outputDate(arg)
	if err != nil {
		return err
	}
	return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "eraconv",
	Short: "西暦と和暦を相互に変換します。",
	Long: `西暦と和暦を相互に変換します。
例えば、'eraconv 令和4年5月10日'と入力すると、'2022年5月10日'が出力されます。
逆に、'eraconv 2022年5月10日' と入力すると、出力として'令和4年5月10日'と出力されます。
引数を省略すると今日の日付を西暦と和暦で表示します。
`,
	Args: customArgValidation,
	RunE: func(cmd *cobra.Command, args []string) error {
		return run(args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.eraconv.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
