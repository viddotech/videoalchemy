package pretty

import (
	"fmt"
	"github.com/fatih/color"
)

func NotifyNormalText(message string, keywords ...string) {
	bold := color.New(color.FgCyan).Add(color.Bold).Add(color.Italic).SprintFunc()
	c := color.New(color.FgHiWhite).SprintFunc()
	boldKeywords := make([]interface{}, len(keywords))
	for i, keyword := range keywords {
		boldKeywords[i] = bold(keyword)
	}
	fmt.Printf(c(message)+"\n", boldKeywords...)
}

func NotifySuccessText(message string, keywords ...string) {
	bold := color.New(color.FgGreen).Add(color.Bold).Add(color.Italic).SprintFunc()
	c := color.New(color.FgHiWhite).SprintFunc()
	boldKeywords := make([]interface{}, len(keywords))
	for i, keyword := range keywords {
		boldKeywords[i] = bold(keyword)
	}
	fmt.Printf(c(message)+"\n", boldKeywords...)
}

func NotifyDangerousText(message string, keywords ...string) {
	bold := color.New(color.FgWhite).Add(color.Bold).Add(color.Italic).SprintFunc()
	c := color.New(color.FgHiRed).SprintFunc()
	boldKeywords := make([]interface{}, len(keywords))
	for i, keyword := range keywords {
		boldKeywords[i] = bold(keyword)
	}
	fmt.Printf(c(message)+"\n", boldKeywords...)
}
