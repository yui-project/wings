package models

type CmdFile struct {
	Index   CmdFileIndex
	Content []CmdFileLine
}

type CmdFileIndex struct {
	FileId           int
	Name             string
	FilePath         string
	CmdFileInfoIndex int
}

type CmdFileLine struct {
	Type   string
	Method string
	// dynamic ?
	Body          *string
	InlineComment string
	StopFlag      bool
	SyntaxError   bool
	ErrorMessage  string
}

type CmdFileLineLog struct {
	Status  CmdFileLineStatus
	Request CmdFileLine
}

type CmdFileLineStatus struct {
	Success bool
	Error   bool
}
