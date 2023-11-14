package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	// TODO: Update to use arg or if no arg, cwd
	cwd, _ := os.Getwd()
	tree := make(map[string]string)
	err := walkTree(cwd, tree)
	if err != nil {
		fmt.Println(err)
	}

	for true {
		newTree := make(map[string]string)
		err := walkTree(cwd, newTree)
		if err != nil {
			fmt.Println(err)
		}
		for dir, hash := range newTree {
			if tree[dir] != hash {
				fmt.Println("Changed!")
			}
		}
		tree = newTree
		// TODO: Upate to a CLI value with default
		time.Sleep(5 * time.Second)
	}
}

func walkTree(dir string, hashes map[string]string) error {
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() && path != "" {
				err = addHash(hashes, path)
				if err != nil {
					return err
				}
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

func addHash(hashes map[string]string, path string) error {
	hasher := sha256.New()
	s, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	hasher.Write(s)
	hashes[path] = hex.EncodeToString(hasher.Sum(nil))
	return nil
}
