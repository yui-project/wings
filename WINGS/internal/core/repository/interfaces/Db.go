package repository_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type IDbRepository interface {
	LoadAllFiles(config models.TlmCmdFileConfig) (interface{}, error)
}
