package fhj_test

import (
	"fmt"
	"log"

	"github.com/jqs7/go-fhj"
)

func Example() {
	client := fhj.New("https://api.zhconvert.org", "")
	_, result, err := client.Convert(fhj.ConverterTaiwan, "什么鬼", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(result.Text)
	// Output: 什麼鬼
}
