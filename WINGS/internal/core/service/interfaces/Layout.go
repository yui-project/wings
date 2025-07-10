package service_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ILayoutService interface {
	GetAllLayout(opid string) ([]models.Layout, error)
	ConfigureLayout(operation models.Operation, config models.TlmCmdFileConfig) error
	SaveLayout(opid string, name string, lytStr string) error
	RemoveLayout(opid string, layoutId int) error
	DeleteLayout(opid string, layoutId int) error
	RemoveLayouts(opid string) error
}
