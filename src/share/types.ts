export type Page = 'meta' | 'entry' | 'category'

export type Route = {
  url: string
  title?: string
  page: Page

  // when page=meta
  // - home: about page on desktop and show menu on mobile
  // - about
  // - contact
  key: string
}

export type Term = {
  display: string
  key: string
  primary?: boolean
}

export type ContentEntry = {
  // raw: string
  html: string
}

export type Link = {
  display: string
  url: string
}

export type Entry = {
  display?: string
  key?: string
  primaryEntry?: Entry

  terms: Term[]
  categories: Category[]
  content: ContentEntry
  footnotes?: string[]
  relatedArticles?: Link[]
  relatedTerms?: Term[]
}

export type Category = {
  display: string
  key: string
  entries?: Entry[]
}
