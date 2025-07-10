package tmtc_processor_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ITcPacketGenerator interface {
	GetTcPacketData(opid string, cmd models.Cmd, cmdType byte, cmdWindow byte, tlmCmdConfigInfo []models.TlmCmdConfigurationInfo) (models.TcPacketData, error)
}
