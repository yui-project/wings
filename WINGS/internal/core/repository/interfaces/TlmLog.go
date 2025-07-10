package repository_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ITlmLog interface {
	AddHistory(opid string, packet models.TlmPacket) error
	GetTelemetryHistory(opid string, telemetryDb []models.TlmPacket) ([]models.TlmPacketHistory, error)
	GetPacketsWithData(opid string) ([]string, error)
	GetRecordPacketsWithData(opid string) ([]string, error)
	InitializeLogFiles(opid string, telemetryDb []models.TlmPacket) error
	GetLogFileStream(opid string, packetName string) error
	GetRecordLogFileStream(opid string, packetName string) error
}
