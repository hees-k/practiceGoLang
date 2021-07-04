package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func removeEmptyFolder(path string) bool {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}

	if len(files) > 1 {
		return false
	} else if len(files) == 1 {
		// mac
		if !files[0].IsDir() && files[0].Name() != ".DS_Store" {
			return false
		}
	}

	err = os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
	return true
}

func travel(root string) ([]string, error) {
	var files []string
	err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if entry.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}

func main() {
	root := "./test_dir"
	if len(os.Args) >= 2 {
		root = os.Args[1]
	} else {
		program := os.Args[0]
		fmt.Printf("usage: %s ./to_delete_directory\n", program)
		return
	}

	files, err := travel(root)
	if err != nil {
		panic(err)
		return
	}

	files = append([]string{root}, files...)

	for i := len(files) - 1; i > 0; i-- {
		var path = files[i]
		if removeEmptyFolder(path) {
			fmt.Println(path, "is removed")
		} else {
			fmt.Println(path, "is not empty")
		}
	}
}
