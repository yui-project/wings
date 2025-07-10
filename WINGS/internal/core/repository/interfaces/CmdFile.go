package repository_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ICmdFileRepository interface {
	LoadCmdFileIndexes(config models.TlmCmdFileConfig) ([]models.CmdFileIndex, error)
	LoadCmdFile(config models.TlmCmdFileConfig, index models.CmdFileIndex, cmdDb []models.Cmd) (models.CmdFile, error)
	GetCmdRow(config models.TlmCmdFileConfig, index models.CmdFileIndex, row int) ([]string, error)
	LoadCmdRow(config models.TlmCmdFileConfig, index models.CmdFileIndex, cmdDb []models.Cmd, row int, line string) (models.CmdFileLine, error)
}
