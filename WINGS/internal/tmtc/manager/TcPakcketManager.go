package tmtc_manager

import (
	"github.com/ut-issl/wings/internal/core/models"
	tmtc_manager_interfaces "github.com/ut-issl/wings/internal/tmtc/manager/interfaces"
)

type TcPacketManager struct {
}

func NewTcPacketManager() tmtc_manager_interfaces.ITcPacketManager {
	return &TcPacketManager{}
}

func (m *TcPacketManager) GetCmdDb(opid string) (models.Cmds, error) {
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
