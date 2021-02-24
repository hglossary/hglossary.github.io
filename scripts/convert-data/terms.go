package main

import (
	"fmt"
	"strings"
)

type RowDesc struct {
	Name  string
	Key   HeaderKey
	Index int
}

type HeaderKey int

const (
	kTerm HeaderKey = iota + 1
	kConnected
	kRelated
	kCategory
	kContent
	kComment
)

var headerRowDesc = map[string]*RowDesc{
	"term":           {Key: kTerm},
	"connected term": {Key: kConnected},
	"related term":   {Key: kRelated},
	"category":       {Key: kCategory},
	"content":        {Key: kContent},
	"comment":        {Key: kComment},
}

type Term struct {
	Display string `json:"display"` // "Low season"
	Key     string `json:"key"`     // "low-season"
}

type ParsedCell struct {
	Terms        []Term        `json:"terms"`
	RelatedTerms []Term        `json:"relatedTerms"`
	Categories   []Term        `json:"categories"`
	Content      ParsedContent `json:"content"`
	Comment      string        `json:"comment"`
}

type ParsedContent struct {
	Raw string `json:"raw"`

	LinkedTerms map[string]string `json:"linkedTerms"`

	// <p>...<a href="/t/room">...</a>...</p>
	HTML string `json:"html"`
}

type Output struct {
	Index map[string]ParsedCell `json:"index"`
}

func init() {
	for name := range headerRowDesc {
		headerRowDesc[name].Name = name
	}
}

func loadHeaderRow(rows Rows) (at int, result map[HeaderKey]*RowDesc) {
	for i := 0; i < 3; i++ { // only load in first 3 row
		result = make(map[HeaderKey]*RowDesc)
		row := rows.Index(i)
		if row == nil {
			panic("can not find header row")
		}
		cells := Cells(row.Cells)
		for j := 0; j < len(cells); j++ {
			cell := cells.Index(j)
			val := strings.ToLower(strings.TrimSpace(cell.Value))
			desc := headerRowDesc[val]
			if desc == nil {
				continue
			}
			if desc.Key == 0 {
				continue
			}
			newDesc := *desc
			newDesc.Index = j
			result[desc.Key] = &newDesc
		}
		if len(result) == len(headerRowDesc) {
			return i, result
		}
	}
	for _, desc := range result {
		fmt.Printf("found row: %v\n", desc.Name)
	}
	panic("can not find header row")
}

func getCellValue(cells Cells, key int) string {
	cell := cells.Index(key)
	if cell == nil {
		return ""
	}
	return strings.TrimSpace(cell.Value)
}

func splitStr(s string) []string {
	ss := strings.Split(s, ";")
	for i := range ss {
		ss[i] = strings.TrimSpace(ss[i])
	}
	return ss
}

func convertKey(s string) string {
	s = strings.ToLower(s)
	return NormalizeUnaccent(s)
}

func parseTerm(s string) Term {
	return Term{
		Display: s,
		Key:     convertKey(s),
	}
}

func parseTerms(ss []string) (rs []Term) {
	for _, s := range ss {
		term := parseTerm(s)
		if term.Key != "" {
			rs = append(rs, term)
		}
	}
	return rs
}

func parseContent(s string) ParsedContent {
	return ParsedContent{
		Raw:  s,
		HTML: `<p>` + s + `</p>`,
	}
}

func loadCellData(mapDesc map[HeaderKey]*RowDesc, rows Rows) (result []ParsedCell) {
	for _, row := range rows {
		cells := Cells(row.Cells)

		term := parseTerm(getCellValue(cells, mapDesc[kTerm].Index))
		if term.Key == "" {
			continue
		}
		terms := []Term{term}
		terms = append(terms, parseTerms(splitStr(getCellValue(cells, mapDesc[kConnected].Index)))...)

		var parsedCell ParsedCell
		parsedCell.Terms = terms
		parsedCell.RelatedTerms = parseTerms(splitStr(getCellValue(cells, mapDesc[kRelated].Index)))
		parsedCell.Categories = parseTerms(splitStr(getCellValue(cells, mapDesc[kCategory].Index)))
		parsedCell.Content = parseContent(getCellValue(cells, mapDesc[kContent].Index))
		parsedCell.Comment = getCellValue(cells, mapDesc[kComment].Index)

		result = append(result, parsedCell)
	}
	return result
}
