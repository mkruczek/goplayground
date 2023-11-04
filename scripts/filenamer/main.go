package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

const regexToRemove = "^[A-Za-z]_\\d{4}-\\d{2}-\\d{2}_"

func main() {
	dir := "..."

	re := regexp.MustCompile(regexToRemove)

	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if re.MatchString(file.Name()) {
			newName := re.ReplaceAllString(file.Name(), "")
			oldPath := filepath.Join(dir, file.Name())
			newPath := filepath.Join(dir, newName)

			err := os.Rename(oldPath, newPath)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("name changed: %s -> %s\n", file.Name(), newName)
		}
	}
}
