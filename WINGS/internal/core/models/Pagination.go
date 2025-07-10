package models

type Pagination struct {
	PageMeta PageMeta
	Data     interface{}
}

type PageMeta struct {
	Page      int
	Size      int
	PageCount int
	Links     PageLink
}

type PageLink struct {
	Self     string
	First    string
	Previous string
	Next     string
	Last     string
}
