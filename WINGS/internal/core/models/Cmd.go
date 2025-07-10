package models

type CmdExecType int

const (
	RT CmdExecType = iota
	TL
	BL
	UTL
)

type Cmd struct {
	Component      string
	ExecType       CmdExecType
	ExecTimeInt    uint
	ExecTimeDouble float64
	ExecTimeStr    string
	Name           string
	Code           string
	Target         string
	Params         []CmdParam
	IsDanger       bool
	IsViaMobc      bool
	IsRestricted   bool
	Description    string
}

type CmdParam struct {
	Name        string
	Type        string
	Value       string
	Unit        string
	Description string
}

type CmdFileLineLogs struct {
	Time      string
	Commander string
	Content   string
	Status    string
}
