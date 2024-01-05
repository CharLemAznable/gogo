package ext_test

import (
	"bufio"
	"errors"
	"github.com/CharLemAznable/gogo/ext"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestDump(t *testing.T) {
	req, _ := http.NewRequest("Get", "/", nil)
	requestBody, _ := ext.DumpRequestBody(req)
	if requestBody != nil {
		t.Errorf("Expected requestBody is nil, but got '%s'", string(requestBody))
	}
	requestBody, _ = ext.DumpRequestBody(req)
	if requestBody != nil {
		t.Errorf("Re-Dump Expected requestBody is nil, but got '%s'", string(requestBody))
	}
	req.Body = errReadBody{}
	_, err := ext.DumpRequestBody(req)
	if err == nil {
		t.Error("Expected Dump requestBody error, but not")
	}

	rsp, _ := http.ReadResponse(bufio.NewReader(
		strings.NewReader("HTTP/1.1 200 OK\r\n\r\nOK")), req)
	responseBody, _ := ext.DumpResponseBody(rsp)
	if "OK" != string(responseBody) {
		t.Errorf("Expected responseBody is 'OK', but got '%s'", string(responseBody))
	}
	responseBody, _ = ext.DumpResponseBody(rsp)
	if "OK" != string(responseBody) {
		t.Errorf("Re-Dump Expected responseBody is 'OK', but got '%s'", string(responseBody))
	}
	rsp.Body = errCloseBody{}
	_, err = ext.DumpResponseBody(rsp)
	if err == nil {
		t.Error("Expected Dump responseBody error, but not")
	}
}

type errReadBody struct{}

func (errReadBody) Read([]byte) (int, error)         { return 0, errors.New("error") }
func (errReadBody) Close() error                     { return nil }
func (errReadBody) WriteTo(io.Writer) (int64, error) { return 0, nil }

type errCloseBody struct{}

func (errCloseBody) Read([]byte) (int, error)         { return 0, io.EOF }
func (errCloseBody) Close() error                     { return errors.New("error") }
func (errCloseBody) WriteTo(io.Writer) (int64, error) { return 0, nil }
