package tmtc_transfer_interfaces

import "github.com/ut-issl/wings/internal/core/models"

type ITmtcPacketService interface {
	Send(data models.TcPacketData) error
}
