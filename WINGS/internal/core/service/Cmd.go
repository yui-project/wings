package service

import (
	"github.com/ut-issl/wings/internal/core/models"
	service_interfaces "github.com/ut-issl/wings/internal/core/service/interfaces"
)

type CmdService struct {
}

func NewCmdService() service_interfaces.ICmdService {
	return &CmdService{}
}

func (s *CmdService) GetAllCmds(opid string) ([]models.Cmd, error) {
	// return bummy data for demonstration purposes
	cmds := []models.Cmd{
		{
			Component:      "OBC",
			ExecType:       models.RT,
			ExecTimeInt:    0,
			ExecTimeDouble: 0.0,
			ExecTimeStr:    "0",
			Name:           "TEST_CMD",
			Code:           "0x01",
			Target:         "Main",
			Params: []models.CmdParam{
				{
					Name:        "param1",
					Type:        "int",
					Value:       "123",
					Unit:        "",
					Description: "Test parameter",
				},
			},
			IsDanger:     false,
			IsViaMobc:    false,
			IsRestricted: false,
			Description:  "Test command for operation " + opid,
		},
	}
	return cmds, nil
}
