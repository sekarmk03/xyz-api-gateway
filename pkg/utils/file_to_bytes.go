package utils

import (
	"bytes"
	"io"
	"mime/multipart"
)

func FileToBytes(file *multipart.FileHeader) ([]byte, error) {
	openedFile, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer openedFile.Close()

	var buf bytes.Buffer
	_, err = io.Copy(&buf, openedFile) // Replacing ioutil.ReadAll with io.Copy
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
