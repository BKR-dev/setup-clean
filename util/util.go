package util

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path"
)

var (
	configDirs = []string{
		".argocd",
		".kube",
		".bonprix",
	}
	workDir = "Bonprix"
	home    = os.Getenv("HOME")
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
	for _, v := range configDirs {

		if !dirExists(v) {
			fmt.Println("creating: ", v)
			err := os.Mkdir(path.Join(home, v), 0755)
			if errors.Is(err, &fs.PathError{}) {
				fmt.Fprintf(os.Stderr, "Could not create dir: %s %s", v, err)
			}

		}
	}

	baseDir := path.Join(home, workDir)
	fmt.Println("creating: ", workDir)
	err := os.Mkdir(baseDir, 0755)
	if errors.Is(err, &fs.PathError{}) {
		fmt.Fprintf(os.Stderr, "Could not create dir: %s %s", baseDir, err)
	}
}

func dirExists(dirName string) bool {
	_, err := os.ReadDir(dirName)
	if err != nil {
		return true
	} else {
		return false
	}
}

func CleanUpAllDirs() {
	for _, v := range configDirs {
		fmt.Println("removing dir: ", v)
		os.RemoveAll(path.Join(home, v))
	}
	fmt.Println("removing dir: ", workDir)
	os.RemoveAll(path.Join(home, workDir))
}
