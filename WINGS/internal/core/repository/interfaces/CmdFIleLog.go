package repository_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ICmdFileLogRepository interface {
	AddHistory(opid string, cmdFileLineLog models.CmdFileLineLog, commanderId string) error
	InitializeLogFiles(opid string) error
	GetLogFileStream(opid string) error
	GetCmdLogHistory(opid string) ([]models.CmdFileLineLog, error)
}
