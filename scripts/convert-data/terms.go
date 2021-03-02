package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
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
	kConnectedTerms
	kSubCategory
	kCategory
	kContent
	kComment
	kRelatedArticles
)

var headerRowDesc = map[string]*RowDesc{
	"thuat ngu":            {Key: kTerm},
	"thuat ngu dong nghia": {Key: kConnectedTerms},
	"danh muc con":         {Key: kSubCategory},
	"danh muc":             {Key: kCategory},
	"noi dung":             {Key: kContent},
	"bai viet lien quan":   {Key: kRelatedArticles},
}

type Term struct {
	Display string `json:"display"` // "Low season"
	Key     string `json:"key"`     // "low-season"
}

type Link struct {
	Display string `json:"display"`
	Url     string `json:"url"`
}

type ParsedCell struct {
	Terms           []Term        `json:"terms"` // and connected terms
	Categories      []Term        `json:"categories"`
	Content         ParsedContent `json:"content"`
	Footnotes       []string      `json:"footnotes,omitempty"`
	RelatedTerms    []Term        `json:"relatedTerms,omitempty"`
	RelatedArticles []Link        `json:"relatedArticles,omitempty"`

	SubCategory string `json:"-"`
}

type ParsedContent struct {
	Raw string `json:"raw,omitempty"`

	// <p>...<a href="/w/room">...</a>...</p>
	HTML string `json:"html"`
}

// Node represent
type Node struct {
	Tag      string
	Attrs    []string // class="foo bar"
	Text     string
	Children []*Node
}

type Nodes []*Node

func (nodes *Nodes) append(node *Node) {
	if node == nil {
		return
	}
	node.Text = strings.TrimSpace(node.Text)
	if node.Text != "" {
		*nodes = append(*nodes, node)
	}
}

func textNode(tag, text string, attrs ...string) *Node {
	return &Node{Tag: tag, Text: text, Attrs: attrs}
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
			val := strings.ReplaceAll(NormalizeUnaccent(strings.TrimSpace(cell.Value)), "-", " ")
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

func splitStr(s string) (res []string) {
	ss := strings.Split(s, ";")
	for i := range ss {
		sss := strings.Split(ss[i], ",")
		for _, item := range sss {
			item = strings.TrimSpace(item)
			if item != "" {
				res = append(res, item)
			}
		}
	}
	return res
}

func convertKey(s string) string {
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

func parseRelatedArticles(s string) (result []Link) {
	lines := reLineBreak.Split(s, -1)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := reExtrLink.FindStringSubmatch(line)
		if parts == nil {
			fmt.Println("ERROR external link:", line)
			continue
		}
		result = append(result, Link{Display: parts[1], Url: parts[2]})
	}
	return result
}

var reListItem = regexp.MustCompile(`^([0-9]{1,2}\. )|([a-zđêôơư]\) )`)
var reLineBreak = regexp.MustCompile(`(\r?\n)+`)
var reFootnote = regexp.MustCompile(`\[\[@([^[\]]+)]]`)
var reIntrLink = regexp.MustCompile(`\[\[([^[\]]+)]]`)
var reExtrLink = regexp.MustCompile(`\[([^[\]]+)]\(([^()]+)\)`)

func parseContent(cell *ParsedCell, input string) ParsedContent {
	nodes := parseOutterQuote(input)
	processNodes(nodes, parseParagraph)

	footnoteNumber := 0
	processNodes(nodes, func(s string) Nodes {
		_nodes, footnotes := parseLinks(s, &footnoteNumber)
		cell.Footnotes = footnotes
		return _nodes
	})

	var out strings.Builder
	buildHtmlFromNodes(&out, nodes)
	return ParsedContent{
		// Raw:  input
		HTML: out.String(),
	}
}

func processNodes(nodes Nodes, fn func(string) Nodes) {
	for _, node := range nodes {
		if len(node.Children) != 0 {
			processNodes(node.Children, fn)
		} else if node.Text != "" {
			node.Children = fn(node.Text)
		}
	}
}

func buildHtmlFromNodes(out *strings.Builder, nodes Nodes) {
	for _, node := range nodes {
		if node.Tag == "a" {
			out.WriteString(" ")
		}
		if node.Tag != "" {
			out.WriteString(`<`)
			out.WriteString(node.Tag)
			for _, attr := range node.Attrs {
				out.WriteString(` `)
				out.WriteString(attr)
			}
			out.WriteString(`>`)
		}
		if len(node.Children) != 0 {
			buildHtmlFromNodes(out, node.Children)
		} else {
			out.WriteString(node.Text)
		}
		if node.Tag != "" {
			out.WriteString(`</`)
			out.WriteString(node.Tag)
			out.WriteString(`>`)
		}
		if node.Tag == "a" {
			out.WriteString(" ")
		}
	}
}

func parseOutterQuote(s string) (nodes Nodes) {
	s = strings.TrimSpace(s)
	for s != "" {
		start, end := findOutterQuote(s)
		if start == -1 {
			nodes.append(textNode("", s))
			return
		}
		if end == -1 {
			end = len(s)
		}
		txt := strings.Trim(s[start:end], "\" \r\n\t")
		nodes.append(textNode("div", txt, `class="quote"`))
		s = strings.TrimSpace(s[end:])
	}
	return nodes
}

func findOutterQuote(s string) (start, end int) {
	// start quote
	idx := strings.Index(s, `"`)
	if idx < 0 {
		return -1, -1
	}
	if idx == 0 || s[idx-1] == '\n' {
		start = idx
	} else {
		return -1, -1
	}

	// end quote
	idx = strings.Index(s, `"`)
	if idx < 0 {
		end = -1
	} else if idx == len(s)-1 || s[idx+1] == '\r' || s[idx+1] == '\n' {
		end = idx + 1
	} else {
		end = -1
	}
	return
}

func parseQuote(s string) (nodes Nodes) {
	parts := strings.Split(s, `"`)
	for i, part := range parts {
		if i%2 == 0 {
			nodes.append(textNode("", part))
		} else {
			nodes.append(textNode("div", part, `class="quote"`))
		}
	}
	return nodes
}

func parseParagraph(s string) (nodes Nodes) {
	parts := reLineBreak.Split(s, -1)
	for _, part := range parts {
		node := &Node{Tag: "p", Text: part}
		if reListItem.MatchString(part) {
			node.Attrs = append(node.Attrs, `class="list-item"`)
		}
		nodes.append(node)
	}
	return nodes
}

func processRegex(
	s string, rs []*regexp.Regexp,
	fn func(string), funcs ...func(string, ...string),
) {
	if len(rs) != len(funcs) {
		panic("func and regex do not match")
	}
	for s != "" {
		mi, minA, minB, mIdx := 0, len(s), len(s), []int(nil)
		for i, re := range rs {
			idx := re.FindStringSubmatchIndex(s)
			if idx == nil {
				continue
			}
			if idx[0] < minA {
				mi, minA, minB, mIdx = i, idx[0], idx[1], idx
			}
		}
		if minA == len(s) {
			break
		}
		if minA > 0 {
			ss := strings.TrimSpace(s[:minA])
			if ss != "" {
				fn(ss)
			}
		}
		var parts []string
		for i := 0; i < len(mIdx); i += 2 {
			a, b := mIdx[i], mIdx[i+1]
			parts = append(parts, s[a:b])
		}
		funcs[mi](s[minA:minB], parts...)
		s = s[minB:]
	}
	ss := strings.TrimSpace(s)
	if ss != "" {
		fn(ss)
	}
}

func parseLinks(input string, footnoteNumber *int) (nodes Nodes, footnotes []string) {
	processRegex(input, []*regexp.Regexp{reFootnote, reIntrLink, reExtrLink},
		func(s string) {
			nodes.append(&Node{Text: s})
		},
		func(s string, ss ...string) {
			*footnoteNumber++
			txt, numberTxt := ss[1], fmt.Sprintf("[%d]", *footnoteNumber)
			href := fmt.Sprintf(`href="#note-%d"`, *footnoteNumber)
			footnotes = append(footnotes, txt)
			nodes.append(textNode("a", numberTxt, href, "data-flink"))
		},
		func(s string, ss ...string) {
			txt, key := ss[1], convertKey(ss[1])
			href := `href="/w/` + key + `"`
			link := `data-ilink="` + key + `"`
			nodes.append(textNode("a", txt, href, link))
		},
		func(s string, ss ...string) {
			txt, link := ss[1], ss[2]
			href := `href=` + strconv.Quote(link)
			nodes.append(textNode("a", txt, href, "data-xlink"))
		},
	)
	return nodes, footnotes
}

func loadCellData(mapDesc map[HeaderKey]*RowDesc, rows Rows) (result []*ParsedCell) {
	for _, row := range rows {
		cells := Cells(row.Cells)

		term := parseTerm(getCellValue(cells, mapDesc[kTerm].Index))
		if term.Key == "" {
			continue
		}
		terms := []Term{term}
		terms = append(terms, parseTerms(splitStr(getCellValue(cells, mapDesc[kConnectedTerms].Index)))...)

		var parsedCell ParsedCell
		parsedCell.Terms = terms
		parsedCell.RelatedArticles = parseRelatedArticles(getCellValue(cells, mapDesc[kRelatedArticles].Index))
		parsedCell.Categories = parseTerms(splitStr(getCellValue(cells, mapDesc[kCategory].Index)))
		parsedCell.Content = parseContent(&parsedCell, getCellValue(cells, mapDesc[kContent].Index))
		parsedCell.SubCategory = convertKey(getCellValue(cells, mapDesc[kSubCategory].Index))

		result = append(result, &parsedCell)
	}
	return result
}

func calcRelatedTerms(cells []*ParsedCell) {
	mapSubcategory := make(map[string][]*ParsedCell)
	for _, cell := range cells {
		sc := cell.SubCategory
		if sc == "" {
			continue
		}
		mapSubcategory[sc] = append(mapSubcategory[sc], cell)
	}
	for _, slice := range mapSubcategory {
		sort.Slice(slice, func(i, j int) bool {
			return slice[i].Terms[0].Key < slice[j].Terms[0].Key
		})
	}
	for _, cell := range cells {
		sc := cell.SubCategory
		if sc == "" {
			continue
		}
		var relatedTerms []Term
		for _, _cell := range mapSubcategory[sc] {
			if _cell.Terms[0].Key != cell.Terms[0].Key {
				relatedTerms = append(relatedTerms, _cell.Terms[0])
			}
		}
		cell.RelatedTerms = relatedTerms
	}
}
