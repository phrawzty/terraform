package aws

import (
	"bytes"
	"testing"
)

func TestReadZipAsBase64(t *testing.T) {
	expected := []byte("VGhpcyBmaWxlIGlzIHRvIHRlc3QgdGhlIGJhc2U2NCBlbmNvZGluZyBmb3IgTGFtYmRhIGNvZGUgdXBsb2FkCg==")
	filename := "test-fixtures/test-lambda.txt"
	zipbytes, err := readFileAsBase64(filename)
	if err != nil {
		t.Errorf("Unable to read %s", filename)
	}
	if !bytes.Equal(expected, zipbytes) {
		t.Errorf("Expected %s but found %s", expected, zipbytes)
	}

	zipbytes, err = readFileAsBase64("this-file-should-not-exist")
	if err == nil {
		t.Error("This should return an error if we can't read the file")
	}
}
