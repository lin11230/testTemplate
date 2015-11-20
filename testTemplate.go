package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

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
	cmd := "wkhtmltopdf"
	args := []string{"--print-media-type", "--page-size", "A4", "-T", "0mm", "-B", "0mm", "-L", "0mm", "-R", "0mm", "--dpi", "600", "test_report.html", "test_report.pdf"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully Generates the report of PDF format")

}
