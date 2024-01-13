package cmd

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

type Translater interface {
	Translate(text string, sourceLang, targetLang language.Tag) (string, error)
}

func Execute(translater Translater) {
	if err := rootCmd(translater).Execute(); err != nil {
		fatal(err)
	}
}

func rootCmd(translater Translater) *cobra.Command {
	var (
		sourceLangStr string
		targetLangStr string
	)

	rootCmd := &cobra.Command{
		Use:   "translater [text]",
		Short: "Text tranlater",
		Long:  `Text tranlater`,
		Run: func(cmd *cobra.Command, args []string) {
			sourceLang, err := language.Parse(sourceLangStr)
			if err != nil {
				fatal(err)
			}
			targetLang, err := language.Parse(targetLangStr)
			if err != nil {
				fatal(err)
			}

			var translateText string
			if len(args) == 0 {
				stdin, err := io.ReadAll(cmd.InOrStdin())
				if err != nil {
					fatal(err)
				}
				translateText = string(stdin)
			} else {
				translateText = strings.Join(args, " ")
			}

			result, err := translater.Translate(translateText, sourceLang, targetLang)
			if err != nil {
				fatal(err)
			}
			fmt.Println(result)
		},
	}

	rootCmd.Flags().StringVarP(&sourceLangStr, "source-lang", "s", "en", "Language of the text to be translated")
	rootCmd.Flags().StringVarP(&targetLangStr, "target-lang", "t", "ru", "The language into which the text should be translated")

	return rootCmd
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
