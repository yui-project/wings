package tmtc_manager_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ITcPacketManager interface {
	GetCmdDb(opid string) (models.Cmds, error)
	// SetCMdDb(opid string, cmdDb []models.Cmd) error
	// RemoveOperation(opid string) error
	// RegisterCmd(opid string, cmd models.Cmd, cmdWindow byte, tlmCmdConfigInfo []models.TlmCmdConfigurationInfo) error
}
