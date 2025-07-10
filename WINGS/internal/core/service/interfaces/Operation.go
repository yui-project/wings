package service_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type IOperationService interface {
	GetCurrentOperations() ([]models.Operation, error)
	GetOperationById(opid string) (models.Operation, error)
	GetOperationHistory(page int, size int, search string) ([]models.Operation, int, error)
	StartOperation(operation models.Operation) (models.Operation, error)
	CancelOperation(opid string) error
	StopOperation(opid string) error
	UpdateOperationHistory(operation models.Operation) error
	DeleteOperationHistory(opid string) error
	GetTlmCmdConfig(opid string) ([]models.TlmCmdConfigurationInfo, error)
}
