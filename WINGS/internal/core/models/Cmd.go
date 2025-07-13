package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

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
type Cmds []Cmd

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

func (c *Cmds) ValidateAndFormat(cmd *Cmd) error {
	exist := false

	cmdFromDb := Cmd{}
	for _, dbCmd := range *c {
		if cmd.Name == dbCmd.Name && cmd.Component == dbCmd.Component {
			exist = true
			cmdFromDb = dbCmd
			break
		}
	}
	if !exist {
		return fmt.Errorf("command not found in CmdDB: %s", cmd.Name)
	}

	// Parameter validation
	if (len(cmdFromDb.Params) != 0 &&
		strings.ToLower(cmdFromDb.Params[len(cmdFromDb.Params)-1].Type) != "raw" &&
		len(cmd.Params) != len(cmdFromDb.Params)) ||
		(len(cmdFromDb.Params) == 0 && len(cmd.Params) != 0) {
		return fmt.Errorf("command parameters mismatch: expected %d, got %d", len(cmdFromDb.Params), len(cmd.Params))
	}

	for i := range cmd.Params {
		switch strings.ToLower(cmd.Params[i].Type) {
		case "int8_t", "int8":
			v, err := strconv.ParseInt(cmd.Params[i].Value, 16, 8)
			if err != nil {
				return fmt.Errorf("wrong type of parameters: %s", cmd.Params[i].Value)
			}
			cmd.Params[i].Value = fmt.Sprintf("0x%X", v)
		case "uint8_t", "uint8":
			v, err := strconv.ParseUint(cmd.Params[i].Value, 16, 8)
			if err != nil {
				return fmt.Errorf("wrong type of parameters: %s", cmd.Params[i].Value)
			}
			cmd.Params[i].Value = fmt.Sprintf("0x%X", v)
		case "int16_t", "int16":
			v, err := strconv.ParseInt(cmd.Params[i].Value, 16, 16)
			if err != nil {
				return fmt.Errorf("wrong type of parameters: %s", cmd.Params[i].Value)
			}
			cmd.Params[i].Value = fmt.Sprintf("0x%X", v)
		case "uint16_t", "uint16":
			v, err := strconv.ParseUint(cmd.Params[i].Value, 16, 16)
			if err != nil {
				return fmt.Errorf("wrong type of parameters: %s", cmd.Params[i].Value)
			}
			cmd.Params[i].Value = fmt.Sprintf("0x%X", v)
		case "int32_t", "int32":
			v, err := strconv.ParseInt(cmd.Params[i].Value, 16, 32)
			if err != nil {
				return fmt.Errorf("wrong type of parameters: %s", cmd.Params[i].Value)
			}
			cmd.Params[i].Value = fmt.Sprintf("0x%X", v)
		case "uint32_t", "uint32":
			v, err := strconv.ParseUint(cmd.Params[i].Value, 16, 32)
			if err != nil {
				return fmt.Errorf("wrong type of parameters: %s", cmd.Params[i].Value)
			}
			cmd.Params[i].Value = fmt.Sprintf("0x%X", v)
		case "float":
			_, err := strconv.ParseFloat(cmd.Params[i].Value, 32)
			if err != nil {
				return fmt.Errorf("wrong type of parameters: %s", cmd.Params[i].Value)
			}
		case "double":
			_, err := strconv.ParseFloat(cmd.Params[i].Value, 64)
			if err != nil {
				return fmt.Errorf("wrong type of parameters: %s", cmd.Params[i].Value)
			}
		case "raw":
			_, err := strconv.ParseUint(cmd.Params[i].Value, 16, 8*len(cmd.Params[i].Value)/2)
			if err != nil {
				return fmt.Errorf("wrong type of parameters: %s", cmd.Params[i].Value)
			}
		default:
			return fmt.Errorf("undefined type: %s (check CMD_DB)", cmd.Params[i].Type)
		}

	}

	return nil
}

func (c *Cmds) CmdToLog(cmd Cmd, opid string) *CmdLog {
	log := &CmdLog{
		SentAt:         time.Now(),
		ExecType:       cmd.ExecType,
		ExecTimeInt:    cmd.ExecTimeInt,
		ExecTimeDouble: cmd.ExecTimeDouble,
		ExecTimeStr:    cmd.ExecTimeStr,
		CmdName:        cmd.Name,
		OperationId:    opid,
	}

	if len(cmd.Params) > 0 {
		log.Param1 = &cmd.Params[0].Value
	}
	if len(cmd.Params) > 1 {
		log.Param2 = &cmd.Params[1].Value
	}
	if len(cmd.Params) > 2 {
		log.Param3 = &cmd.Params[2].Value
	}
	if len(cmd.Params) > 3 {
		log.Param4 = &cmd.Params[3].Value
	}
	if len(cmd.Params) > 4 {
		log.Param5 = &cmd.Params[4].Value
	}
	if len(cmd.Params) > 5 {
		log.Param6 = &cmd.Params[5].Value
	}

	return log
}
