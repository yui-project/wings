package models

import "time"

type CmdLog struct {
	SentAt         time.Time
	ExecType       CmdExecType
	ExecTimeInt    uint
	ExecTimeDouble float64
	ExecTimeStr    string
	CmdName        string
	Param1         string
	Param2         string
	Param3         string
	Param4         string
	Param5         string
	Param6         string
	OperationId    string
	Operation      Operation
}
