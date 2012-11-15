package edit

import(
	"bufio"
	"io"
	"regexp"
	"compress/flate"
	"github.com/bjh83/pdfstrip/decode"
)

func replaceText(reader *bufio.Reader, writer io.Writer, newText string) error {
	endEx, _ := regexp.Compile("endobj")
	byteBuffer := make([]byte, len(newText))
	buffer := bytes.NewBuffer(byteBuffer)
	tempWriter := flate.NewWriter(buffer, -1)
	_, err := tempWriter.Write([]byte(newText))
	err := tempWriter.Flush()
	if err != nil {
		return err
	}
	byteBuffer = byteBuffer[12:]
	byteBuffer[0] = byte(120)
	byteBuffer[1] = byte(156)
	_, err = writer.Write(buffer)
	var line string
	for line, err = reader.ReadString('\n'); err == nil && !endEx.MatchString(line); line, err = reader.ReadString('\n') {}
	_, err = writer.Write([]byte(line))
	return err
}

func WriteChanges(toRead io.Reader, writer io.Writer, fileData *decode.FileData) error {
	dataMap := fileData.GetMap()
	reader := bufio.NewReader(toRead)
	endEx, _ := regexp.Compile("endobj")
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		_, err = writer.Write([]byte(line))
		if err != nil {
			return err
		}
		hasId, id := decode.GetID(line)
		if hasId {
			newText, present := dataMap[id]
			if present {
				err = replaceText(reader, writer, newText)
				if err != nil {
					return err
				}
			}
		}
	}
}

