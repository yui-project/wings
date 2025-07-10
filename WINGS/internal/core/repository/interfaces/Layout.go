package repository_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ILayoutRepository interface {
	LoadAllFiles(config models.TlmCmdFileConfig) (interface{}, error)
	SaveLayout(config models.TlmCmdFileConfig, name string, lytStr string) error
	RemaneLayout(config models.TlmCmdFileConfig, name string, oldName string) error
	DeleteLayout(config models.TlmCmdFileConfig, name string) error
}
