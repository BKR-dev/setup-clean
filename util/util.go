package util

import (
	"fmt"
	"os"
	"os/exec"
)

// open useful links in borwser
// HAS to be http(s)://www._.com
// (bookmarks, oracleclou, jira boards...)
func OpenLinksInBrowser() {
	links := []string{
		"https://www.google.com",
		"https://www.gobyexample.com",
		"https://www.github.com"}

	for i := 0; i < len(links); i++ {
		err := exec.Command("open", links[i]).Start()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening URL: %s", links[i])
		}
	}
}

// opens a single link
func OpenSingleLinkInBrowser(url string) {
	err := exec.Command("open", url).Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening URL: %s", url)
	}
}
