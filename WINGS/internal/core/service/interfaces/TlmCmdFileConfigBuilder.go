package service_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ITlmCmdFileConfigBuilder interface {
	Build(opid string) (models.TlmCmdFileConfig, error)
}
