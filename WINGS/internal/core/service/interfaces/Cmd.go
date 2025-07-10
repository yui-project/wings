package service_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ICmdService interface {
	GetAllCmds(opid string) ([]models.Cmd, error)
	// SendCmd(opid string, Cmd models.Cmd, CommanderId string) error
	// SendTypeACmd(opid string, Cmd models.Cmd, CommanderId string, cmdWindow int) error
	// InitializeTypeAstatus(opid string) error
	// SendRawCmd(opid string, packet []byte) error
	// AddCmdFileLineLog(opid string, CmdFileLineLog models.CmdFileLineLog, CommanderId string) error
	// GetCmdFileIndexes(opid string) ([]models.CmdFileIndex, error)
	// GetCmdFile(opid string, cmdFileInfoIndex int, fileId int) ([]models.CmdFile, error)
	// ConfigureCmdDB(operation models.Operation, config models.TlmCmdFileConfig) error
	// ConfigureCmdFile(opration models.Operation, config models.TlmCmdFileConfig) error
	// ConfigureCmdFileLog(operation models.Operation) error
	// ReconfigureCmdFile(opid string) error
	// RemoveCmdFileIndexes(opid string) error
	// // stream ?
	// GetCmdLogStream(opid string) error
	// GetCmdLogHistory(opid string) ([]models.CmdFileLineLog, error)
	// // stream ?
	// GetCmdFileLogSteam(opid string) error
	// GetCmdRow(opid string, cmdFileInfoIndex int, fileId int, row int) (string, error)
	// LoadCmdRow(opid string, cmdFileInfoIndex int, fileId int, row int, line string) (models.CmdFileLine, error)
}
