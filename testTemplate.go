package main

import (
	"fmt"
	"log"
	"os"

	"github.com/fatih/structs"
	"github.com/hoisie/mustache"
)

var (
	templateHtml *mustache.Template
	tmpDir       = "./"
)

type Data struct {
	Name string
}

type PeakItem struct {
	MoM      string
	PeakItem int
}

func main() {
	templateHtml, err := mustache.ParseFile("./report.html")
	if err != nil {
		log.Fatal("Parse Html failed, ", err)
	}

	var tmpfile = tmpDir + "test_report.html"
	tmpHtmlFile, err := os.Create(tmpfile)
	if err != nil {
		log.Fatal("Create tmp file failed, ", err)
	}
	var data Data
	data.Name = "!!! XXXX !!!!"
	mapping := structs.Map(data)
	var str = templateHtml.Render(mapping)
	filelength, err := tmpHtmlFile.WriteString(str)
	fmt.Printf("wrote %d bytes\n", filelength)
	tmpHtmlFile.Sync()

}
