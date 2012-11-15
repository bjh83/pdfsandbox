package main

import(
	"os"
	"fmt"
	"github.com/bjh83/pdfstrip/decode"
	"pdfplay/edit"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide two file names...")
		return
	}
	fileIn, fileErr := os.Open(os.Args[1])
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}
	fileOut, fileErr := os.Create(os.Args[2])
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}
	fileData, fileErr := decode.Decode(fileIn)
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}
	fileData.Blocks[0].Text = fileData.Blocks[0].Text[:30]
	for index := 1; index < len(fileData.Blocks); index++ {
		fileData.Blocks[index].Text = "\n"
	}
	_, fileErr = fileIn.Seek(0, 0)
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}
	fileErr = edit.WriteChanges(fileIn, fileOut, fileData)
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}
}

