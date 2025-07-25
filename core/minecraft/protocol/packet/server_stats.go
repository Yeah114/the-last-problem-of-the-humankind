package packet

import "github.com/Happy2018new/the-last-problem-of-the-humankind/core/minecraft/protocol"

// ServerStats is a packet sent from the server to the client to update the client on server statistics. It is purely
// used for telemetry.
type ServerStats struct {
	// ServerTime ...
	ServerTime float32
	// NetworkTime ...
	NetworkTime float32
}

// ID ...
func (pk *ServerStats) ID() uint32 {
	return IDServerStats
}

func (pk *ServerStats) Marshal(io protocol.IO) {
	io.Float32(&pk.ServerTime)
	io.Float32(&pk.NetworkTime)
}
