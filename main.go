package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	matches   []string
	lock      = sync.Mutex{}
	waitgroup = sync.WaitGroup{}
)

func reportErrorAndExit(e error) {
	log.Fatal(e)
	os.Exit(1)
}

func fileSearch(root string, filename string) {
	fmt.Printf("Searchig for %s in %s\n", filename, root)

	files, err := os.ReadDir(root)
	if err != nil {
		reportErrorAndExit(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), filename) {
			fmt.Println()
			lock.Lock()
			matches = append(matches, filepath.Join(root, file.Name()))
			lock.Unlock()
		}

		if file.IsDir() {
			waitgroup.Add(1)
			go fileSearch(filepath.Join(root, file.Name()), filename)
		}
	}
	waitgroup.Done()
}

func main() {
	if args := os.Args; len(args) != 3 {
		e := errors.New("2 string arguments required: [root dir] [filename to search]")
		reportErrorAndExit(e)
	} else {
		root, filename := args[1], args[2]

		waitgroup.Add(1)
		go fileSearch(root, filename)
		waitgroup.Wait()

		if len(matches) == 0 {
			fmt.Printf("❌ no matches found\n")
			os.Exit(0)
		}

		for _, file := range matches {
			fmt.Printf("✅ Matched %s\n", file)
		}
		fmt.Printf("ℹ️ Total matched files: %d\n", len(matches))
	}
}
