package packet

import (
	"github.com/Happy2018new/the-last-problem-of-the-humankind/core/minecraft/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

// PlayerInput is sent by the client to the server when the player is moving but the server does not allow it
// to update its movement using the MovePlayer packet. It includes situations where the player is riding an
// entity like a boat. If this is the case, the packet is sent roughly every tick.
type PlayerInput struct {
	// Movement is the movement vector of the input. It should be thought of in Pocket Edition controls, where
	// specific arrows (or a combination of two, resulting in a diagonal arrow) decide the direction of
	// movement. The movement vector typically has a length of 1: Either it has movement on one axis, or it
	// has a combination, resulting in sqrt(2)/2 for both axes.
	Movement mgl32.Vec2
	// Jumping indicates if the player was pressing the jump button during the input. It does not define if
	// the player was actually in the air or not.
	Jumping bool
	// Sneaking indicates if the player was sneaking during the input. Note that this may also be checked by
	// keeping the sneaking state updated using the PlayerAction packet.
	Sneaking bool
}

// ID ...
func (*PlayerInput) ID() uint32 {
	return IDPlayerInput
}

func (pk *PlayerInput) Marshal(io protocol.IO) {
	io.Vec2(&pk.Movement)
	io.Bool(&pk.Jumping)
	io.Bool(&pk.Sneaking)
}
