package main

type ParsedCell struct {
	Terms           []Term        `json:"terms"` // and connected terms
	Categories      []Term        `json:"categories"`
	Content         ParsedContent `json:"content"`
	Footnotes       []string      `json:"footnotes,omitempty"`
	RelatedTerms    []Term        `json:"relatedTerms,omitempty"`
	RelatedArticles []Link        `json:"relatedArticles,omitempty"`

	SubCategory string `json:"-"`
}

type Term struct {
	Display string `json:"display"` // "Low season"
	Key     string `json:"key"`     // "low-season"
}

type Link struct {
	Display string `json:"display"`
	Url     string `json:"url"`
}

type ParsedContent struct {
	Raw string `json:"raw,omitempty"`

	// <p>...<a href="/w/room">...</a>...</p>
	HTML string `json:"html"`
}
