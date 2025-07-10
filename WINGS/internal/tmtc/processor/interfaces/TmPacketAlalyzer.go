package tmtc_processor_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ITmPacketAnalyzer interface {
	AnalyzeTmPacket(data models.TmPacketData, prevTelemetry []models.TlmPacket) (bool, error)
	GetCmdWindow() (byte error)
	GetRetransmitFlag() bool
	RemoveOperation(opid string) error
}
