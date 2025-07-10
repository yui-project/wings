package tmtc_manager_interfaces

import (
	"github.com/ut-issl/wings/internal/core/models"
	tmtc_processor_interfaces "github.com/ut-issl/wings/internal/tmtc/processor/interfaces"
)

type ITmtcHandlerFactory interface {
	GetTmPacketAnalyzer(opid string) (tmtc_processor_interfaces.ITmPacketAnalyzer, error)
	GetTcPacketGenerator(opid string) (tmtc_processor_interfaces.ITcPacketGenerator, error)
	GetTmtcPackerService(opid string) (tmtc_processor_interfaces.ITcPacketGenerator, error)
	AddOperation(opid string, component models.Component, target models.TmtcTarget) error
	RemoveOperation(opid string) error
}
