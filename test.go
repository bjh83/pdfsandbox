package main

import(
	"os"
	"fmt"
	"github.com/bjh83/pdfstrip/decode"
	"pdfsandbox/edit"
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
	_, fileErr = fileIn.Seek(0, 0)
	if fileErr != nil {
		fmt.Println(fileErr.Error())
		return
	}
	block, fileErr := decode.GetXRef(fileIn)
	fmt.Println(block.ID)
	data := []byte(block.Text)
	for index := 5; index < len(data); index += 6 {
		fmt.Printf("%x  %x%x%x  %x%x\n", data[index - 5], data[index - 4], data[index - 3], data[index - 2], data[index - 1], data[index - 0])
	}
	/*
	fileData.Blocks[0].Text = fileData.Blocks[0].Text[:30]
	for index := 1; index < len(fileData.Blocks); index++ {
		fileData.Blocks[index].Text = "\n"
	}
	*/
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

