package tmtc_manager_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ITmPacketManager interface {
	RemoveOperation(opid string) error
	SetTelemetryDb(opid string, telemetryDb []models.TlmPacket) error
	GetTelemetryDb(opid string) ([]models.TlmPacket, error)
	GetLatestTelemetry(opid string) ([]models.TlmPacket, error)
	RegisterTelemetry(data models.TmPacketData) error
}
