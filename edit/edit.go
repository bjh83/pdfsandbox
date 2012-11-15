package edit

import(
	"bufio"
	"io"
	"bytes"
	"regexp"
	"strconv"
	"compress/flate"
	"github.com/bjh83/pdfstrip/decode"
)

func replaceText(reader *bufio.Reader, writer io.Writer, newText string) error {
	endEx, _ := regexp.Compile("endobj")
	buffer := new(bytes.Buffer)
	tempWriter, err := flate.NewWriter(buffer, 5)
	if err != nil {
		return err
	}
	_, err = tempWriter.Write([]byte(newText))
	err = tempWriter.Close()
	if err != nil {
		return err
	}
	lengthString := strconv.FormatInt(int64(buffer.Len() + 2), 10)
	writer.Write([]byte("<</Length " + lengthString + "/Filter/FlateDecode>>stream\n"))
	_, err = writer.Write([]byte{120, 156})
	_, err = buffer.WriteTo(writer)
	_, err = writer.Write([]byte{'\n'})
	var line string
	for line, err = reader.ReadString('\n'); err == nil && !endEx.MatchString(line); line, err = reader.ReadString('\n') {}
	_, err = writer.Write([]byte("endstream\n"))
	_, err = writer.Write([]byte(line))
	return err
}

func WriteChanges(toRead io.Reader, writer io.Writer, fileData *decode.FileData) error {
	dataMap := fileData.GetMap()
	reader := bufio.NewReader(toRead)
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
			newText, present := dataMap[int(id)]
			if present {
				err = replaceText(reader, writer, newText)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

