package main

import (
	"atcoder/template"
	"os"
	"path"
	"path/filepath"
	"time"
)

func main() {
	dir := time.Now().Format("20060102")

	dir = path.Join("cmd", dir)

	for _, dir := range []string{
		path.Join(dir, "A"),
		path.Join(dir, "B"),
		path.Join(dir, "C"),
		path.Join(dir, "D"),
		path.Join(dir, "E"),
	} {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(err)
		}
	}

	for _, file := range []string{
		filepath.Join(dir, "A", "A.go"),
		filepath.Join(dir, "B", "B.go"),
		filepath.Join(dir, "C", "C.go"),
		filepath.Join(dir, "D", "D.go"),
		filepath.Join(dir, "E", "E.go"),
	} {
		_, err := os.Stat(file)
		if err == nil {
			continue
		}

		if !os.IsNotExist(err) {
			panic(err)
		}

		err = os.WriteFile(file, template.Template, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
