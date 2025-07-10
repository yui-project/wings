package tmtc_transfer_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ITcPacketQueue interface {
	Add(opid string) error
	Remove(opid string) error
	Enqueue(data models.TcPacketData) error
	Dequeue(opid string) (models.TcPacketData, error)
	PacketQueueExists(opid string) bool
}
