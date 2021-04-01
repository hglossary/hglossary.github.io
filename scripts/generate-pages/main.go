package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
)

func main() {
	data := mustRead("src/_tmp/data.js")
	indexHtml := mustRead("docs/index.html")

	var cells []*ParsedCell
	data = bytes.TrimPrefix(data, []byte("export default"))
	must(json.Unmarshal(data, &cells))

	for _, cell := range cells {
		for _, term := range cell.Terms {
			filename := "docs/w/" + term.Key + ".html"
			must(ioutil.WriteFile(filename, indexHtml, 0644))
		}
	}
}

func mustRead(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	must(err)
	return data
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
