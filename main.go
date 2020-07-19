package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/browser"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: gg query")
		return
	}

	queries := strings.Join(os.Args[1:], " ")
	const urlFmt = "https://google.com/m?q=%s"
	searchURL := fmt.Sprintf(urlFmt, queries)
	fmt.Println("Search URL:", searchURL)
	if err := browser.OpenURL(searchURL); err != nil {
		fmt.Println("error:", err)
	}
}
