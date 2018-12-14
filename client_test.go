package fhj_test

import (
	"strconv"
	"testing"

	"github.com/jqs7/go-fhj"
)

var client = fhj.New("https://api.zhconvert.org", "")

func TestClient_ServiceInfo(t *testing.T) {
	_, status, err := client.ServiceInfo()
	if err != nil {
		t.Fatal(err)
	}
	if len(status.Converters) == 0 {
		t.Fatal("converters is empty")
	}
}

func TestClient_Convert(t *testing.T) {
	_, data, err := client.Convert(fhj.ConverterHongkong, "Hello 什么鬼 こんばんは", map[string]string{
		"diffEnable":               strconv.FormatBool(true),
		"jpTextConversionStrategy": "protectOnlySameOrigin",
		"jpTextStyles":             "OPJP,EDJP,*noAutoJpTextStyles",
	})
	if err != nil {
		t.Fatal(err)
	}
	if data.Text != "Hello 甚麼鬼 こんばんは" {
		t.Fatal("convert result error")
	}
	if len(data.JPTextStyles) == 0 {
		t.Fatal("may be something goes wrong")
	}
}
