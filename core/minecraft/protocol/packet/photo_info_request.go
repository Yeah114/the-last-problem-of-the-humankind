package packet

import (
	"github.com/Happy2018new/the-last-problem-of-the-humankind/core/minecraft/protocol"
)

// PhotoInfoRequest is sent by the client to request photo information from the server. This packet was deprecated in 1.19.80.
type PhotoInfoRequest struct {
	// PhotoID is the ID of the photo.
	PhotoID int64
}

// ID ...
func (*PhotoInfoRequest) ID() uint32 {
	return IDPhotoInfoRequest
}

func (pk *PhotoInfoRequest) Marshal(io protocol.IO) {
	io.Varint64(&pk.PhotoID)
}
