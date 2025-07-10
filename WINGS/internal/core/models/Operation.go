package models

import "time"

type TmtcTarget int

const (
	TmtcIf TmtcTarget = iota
	Infostellar
)

type Operation struct {
	Id              string
	PathNumber      string
	Comment         string
	CreatedAt       time.Time
	IsRunning       bool
	IsTmtcConnected bool
	FileLocation    TlmCmdFileLocation
	TmtcTarget      TmtcTarget
	ComponentId     string
	Component       Component
}
