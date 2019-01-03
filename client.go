package fhj

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/levigross/grequests"
)

type Client struct {
	endPoint string
	apiKey   string
}

func New(endPoint, apiKey string) *Client {
	return &Client{
		endPoint: strings.TrimRight(endPoint, "/"),
		apiKey:   apiKey,
	}
}

func (c Client) ServiceInfo() (*RawResp, *StatusData, error) {
	respJSON, err := c.req(http.MethodGet, "/service-info", nil)
	if err != nil {
		return respJSON, nil, err
	}
	status := &StatusData{}
	if err := json.Unmarshal(respJSON.Data, status); err != nil {
		return respJSON, nil, err
	}
	return respJSON, status, nil
}

const (
	ConverterSimplified      = "Simplified"
	ConverterTraditional     = "Traditional"
	ConverterChina           = "China"
	ConverterHongkong        = "Hongkong"
	ConverterTaiwan          = "Taiwan"
	ConverterPinyin          = "Pinyin"
	ConverterMars            = "Mars"
	ConverterWikiSimplified  = "WikiSimplified"
	ConverterWikiTraditional = "WikiTraditional"
)

// Convert extra options: https://docs.zhconvert.org/api/convert/#%E5%AD%97%E5%B9%95%E6%A8%A3%E5%BC%8F
func (c Client) Convert(converter, text string, extraOptions map[string]string) (*RawResp, *ConvertData, error) {
	data := map[string]string{
		"text":      text,
		"converter": converter,
	}
	for k, v := range extraOptions {
		data[k] = v
	}
	respJSON, err := c.req(http.MethodPost, "/convert", &grequests.RequestOptions{
		Data: data,
	})
	if err != nil {
		return respJSON, nil, err
	}
	respData := &ConvertData{}
	if err := json.Unmarshal(respJSON.Data, respData); err != nil {
		return respJSON, nil, err
	}
	return respJSON, respData, nil
}

func (c Client) req(verb, url string, opt *grequests.RequestOptions) (*RawResp, error) {
	if opt == nil {
		opt = &grequests.RequestOptions{}
	}
	if opt.Params == nil {
		opt.Params = make(map[string]string)
	}
	opt.Params["apiKey"] = c.apiKey
	resp, err := grequests.Req(verb, c.endPoint+url, opt)
	if err != nil {
		return nil, err
	}
	if !resp.Ok {
		return nil, errors.New(http.StatusText(resp.StatusCode))
	}
	respJSON := &RawResp{}
	if err := resp.JSON(respJSON); err != nil {
		return nil, err
	}
	if respJSON.Code != 0 {
		return respJSON, errors.New(respJSON.Msg)
	}
	return respJSON, nil
}
