package models

type Layout struct {
	TlmView TlmView
	Id      InternalError
	Name    string
}

type TlmView struct {
	AllIndexes []TlmViewIndex
	Blocks     []ViewBlockInfo
	// object ?
	Content string
}

type TlmViewIndex struct {
	Id                  string
	Name                string
	CompoName           string
	FilePath            string
	Type                string
	SelectedTelemetries []string
	DataType            string
	DataLength          string
	YlabelMin           string
	YlabelMax           string
}

type ViewBlockInfo struct {
	Tabs      []TlmViewIndex
	ActiveTab int
}
