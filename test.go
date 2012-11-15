package main

import(
	"os"
	"fmt"
	"io/ioutil"
	"github.com/bjh83/pdfstrip/decode"
	"github.com/bjh83/pdfstrip/deformat"
	"pdfplay/edit"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide two files...")
		return
	}
	fileIn, fileErr := os.Open(os.Args[1])
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}
	fileData, fileErr := decode.Decode(fileIn)
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}
	fileData.Blocks[0].Text = fileData.Blocks[0].Text[0:10]
	for index := 1; index < len(fileData.Blocks); index++ {
		fileData.Blocks[index] = "\n"
	}
	edit.WriteChanges(fileIn, fileData)
