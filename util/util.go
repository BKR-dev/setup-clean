package util

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
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

// setup directory structure
func MakeDirSkeleton() {
	// setup dirs
	configDirs =[]string{".argocd", ".kube", ".bonprix"}

	for _, d := range configDirs {
		if !dirExists(d) {
			// make dirs
		}
	}

	workDir := "Bonprix"
	home := os.Getenv("HOME")
	baseDir := path.Join(home, workDir)
	// create dir
	err := os.Mkdir(baseDir, 0755)
	if errors.Is(err, &fs.PathError{}) {
		fmt.Fprintf(os.Stderr, "Could not create dir: %s %s", baseDir, err)
	}

}

func dirExists(name string) bool {
		
}

