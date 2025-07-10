package tmtc_manager_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ITcPacketManager interface {
	SetCMdDb(opid string, cmdDb []models.Cmd) error
	RemoveOperation(opid string) error
	GetCmdDb(opid string) ([]models.Cmd, error)
	RegisterCmd(opid string, cmd models.Cmd, cmdWindow byte, tlmCmdConfigInfo []models.TlmCmdConfigurationInfo) error
}
