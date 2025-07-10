package service_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ITlmService interface {
	GetLatestTlm(opid string, refTlmTime string) (models.LatestTlm, error)
	GetTlmHistory(opid string) ([]models.TlmHistory, int, error)
	GetPacketsWithData(opid string) ([]string, error)
	GetRecordPacketsWithData(opid string) ([]string, error)
	ConfigureTlmDB(operation models.Operation, config models.TlmCmdFileConfig) error
	// stream ?
	GetTlmLogStream(opid string) error
	// stream ?
	GetRecordTlmLogStream(opid string, packetNames []string) error
}
