package main

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/ara-framework/nova-proxy/parser"
	"github.com/gookit/color"
)

func main() {
	color.Info.Println("Running Nova Static")
	staticFolder := os.Getenv("STATIC_FOLDER")
	filepath.Walk(staticFolder, func(path string, info os.FileInfo, err error) error {
		fileName := info.Name()
		if strings.HasSuffix(fileName, ".html") && !info.IsDir() {
			e, err := ioutil.ReadFile(path)

			if err != nil {
				log.Fatal(path + " can not be read \n")
			}

			color.Info.Tips(path)

			newHTML := parser.ModifyBody(string(e))

			ioutil.WriteFile(path, []byte(newHTML), 0644)
		}
		return nil
	})

	color.Info.Println("Files processed successfully!!")
}
