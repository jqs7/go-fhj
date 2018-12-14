package fhj

import (
	"encoding/json"
)

type RawResp struct {
	Code      int
	Data      json.RawMessage
	Msg       string
	Revisions Revisions
	ExecTime  float64
}

type Revisions struct {
	Build string
	Msg   string
	Time  int
}

type StatusData struct {
	Converters          map[string]ConverterDesc
	Modules             map[string]ModuleDesc
	ConverterCategories map[string]string
	ModuleCategories    map[string]string
	TextFormats         map[string]string
	DiffTemplates       map[string]DiffTemplateDesc
	AllowEmptyApiKey    bool
	MaxPostBodyBytes    int
}

type ConvertData struct {
	Converter    string
	Text         string
	Diff         string
	TextFormat   string
	UsedModules  []string
	JPTextStyles []string
}

type ConverterDesc struct {
	Name string
	Desc string
	Cat  string
}

type ModuleDesc struct {
	Name     string
	Desc     string
	Cat      string
	IsManual bool
}

type DiffTemplateDesc struct {
	Desc string
}
