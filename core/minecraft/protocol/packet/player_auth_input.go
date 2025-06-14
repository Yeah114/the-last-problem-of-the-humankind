package packet

import (
	"github.com/Happy2018new/the-last-problem-of-the-humankind/core/minecraft/protocol"

	"github.com/go-gl/mathgl/mgl32"
)

const (
	InputFlagAscend = 1 << iota
	InputFlagDescend
	InputFlagNorthJump
	InputFlagJumpDown
	InputFlagSprintDown
	InputFlagChangeHeight
	InputFlagJumping
	InputFlagAutoJumpingInWater
	InputFlagSneaking
	InputFlagSneakDown
	InputFlagUp
	InputFlagDown
	InputFlagLeft
	InputFlagRight
	InputFlagUpLeft
	InputFlagUpRight
	InputFlagWantUp
	InputFlagWantDown
	InputFlagWantDownSlow
	InputFlagWantUpSlow
	InputFlagSprinting
	InputFlagAscendBlock
	InputFlagDescendBlock
	InputFlagSneakToggleDown
	InputFlagPersistSneak
	InputFlagStartSprinting
	InputFlagStopSprinting
	InputFlagStartSneaking
	InputFlagStopSneaking
	InputFlagStartSwimming
	InputFlagStopSwimming
	InputFlagStartJumping
	InputFlagStartGliding
	InputFlagStopGliding
	InputFlagPerformItemInteraction
	InputFlagPerformBlockActions
	InputFlagPerformItemStackRequest
	InputFlagHandledTeleport
	InputFlagEmoting
	InputFlagMissedSwing
	InputFlagStartCrawling
	InputFlagStopCrawling
	InputFlagStartFlying
	InputFlagStopFlying
	InputFlagClientAckServerData
)

const (
	InputModeMouse = iota + 1
	InputModeTouch
	InputModeGamePad
	InputModeMotionController
)

const (
	PlayModeNormal = iota
	PlayModeTeaser
	PlayModeScreen
	PlayModeViewer
	PlayModeReality
	PlayModePlacement
	PlayModeLivingRoom
	PlayModeExitLevel
	PlayModeExitLevelLivingRoom
	PlayModeNumModes
)

const (
	InteractionModelTouch = iota
	InteractionModelCrosshair
	InteractionModelClassic
)

// PlayerAuthInput is sent by the client to allow for server authoritative movement. It is used to synchronise
// the player input with the position server-side.
// The client sends this packet when the ServerAuthoritativeMovementMode field in the StartGame packet is set
// to true, instead of the MovePlayer packet. The client will send this packet once every tick.
type PlayerAuthInput struct {
	// Pitch and Yaw hold the rotation that the player reports it has.
	Pitch, Yaw float32
	// Position holds the position that the player reports it has.
	Position mgl32.Vec3
	// MoveVector is a Vec2 that specifies the direction in which the player moved, as a combination of X/Z
	// values which are created using the WASD/controller stick state.
	MoveVector mgl32.Vec2
	// HeadYaw is the horizontal rotation of the head that the player reports it has.
	HeadYaw float32
	// InputData is a combination of bit flags that together specify the way the player moved last tick. It
	// is a combination of the flags above.
	InputData uint64
	// InputMode specifies the way that the client inputs data to the screen. It is one of the constants that
	// may be found above.
	InputMode uint32
	// PlayMode specifies the way that the player is playing. The values it holds, which are rather random,
	// may be found above.
	PlayMode uint32
	// GazeDirection is the direction in which the player is gazing, when the PlayMode is PlayModeReality: In
	// other words, when the player is playing in virtual reality.
	GazeDirection mgl32.Vec3
	// Tick is the server tick at which the packet was sent. It is used in relation to
	// CorrectPlayerMovePrediction.
	Tick uint64
	// Delta was the delta between the old and the new position. There isn't any practical use for this field
	// as it can be calculated by the server itself.
	Delta mgl32.Vec3
	// ItemInteractionData is the transaction data if the InputData includes an item interaction.
	ItemInteractionData protocol.UseItemTransactionData
	// ItemStackRequest is sent by the client to change an item in their inventory.
	ItemStackRequest protocol.ItemStackRequest
	// BlockActions is a slice of block actions that the client has interacted with.
	BlockActions []protocol.PlayerBlockAction
	// AnalogueMoveVector is a Vec2 that specifies the direction in which the player moved, as a combination of X/Z
	// values which are created using an analogue input.
	AnalogueMoveVector mgl32.Vec2

	/*
		PhoenixBuilder specific changes.
		Author: Liliya233

		InteractionModel is a constant representing the interaction model the player is using. It is one of the
		constants that may be found above.

		For netease, the data type of this field is uint32,
		but on standard minecraft, this is int32.
	*/
	InteractionModel uint32
	// InteractionModel int32

	// PhoenixBuilder specific fields.
	// Author: Liliya233
	//
	// The following fields are NetEase specific.
	Unknown1 float32
	Unknown2 float32
	Unknown3 bool
	Unknown4 bool
	Unknown5 bool
}

// ID ...
func (pk *PlayerAuthInput) ID() uint32 {
	return IDPlayerAuthInput
}

func (pk *PlayerAuthInput) Marshal(io protocol.IO) {
	io.Float32(&pk.Pitch)
	io.Float32(&pk.Yaw)
	io.Vec3(&pk.Position)
	io.Vec2(&pk.MoveVector)
	io.Float32(&pk.HeadYaw)
	io.Varuint64(&pk.InputData)
	io.Varuint32(&pk.InputMode)
	io.Varuint32(&pk.PlayMode)

	// PhoenixBuilder specific changes.
	// Author: Liliya233
	io.Varuint32(&pk.InteractionModel)
	// io.Varint32(&pk.InteractionModel)

	if pk.PlayMode == PlayModeReality {
		io.Vec3(&pk.GazeDirection)
	}
	io.Varuint64(&pk.Tick)
	io.Vec3(&pk.Delta)

	// PhoenixBuilder specific changes.
	// Author: Liliya233
	//
	// NetEase
	io.Bool(&pk.Unknown3)

	if pk.InputData&InputFlagPerformItemInteraction != 0 {
		io.PlayerInventoryAction(&pk.ItemInteractionData)
	}

	if pk.InputData&InputFlagPerformItemStackRequest != 0 {
		protocol.Single(io, &pk.ItemStackRequest)
	}

	if pk.InputData&InputFlagPerformBlockActions != 0 {
		protocol.SliceVarint32Length(io, &pk.BlockActions)
	}

	io.Vec2(&pk.AnalogueMoveVector)

	// PhoenixBuilder specific changes.
	// Author: Liliya233
	//
	// NetEase
	{
		io.Bool(&pk.Unknown5)
		io.Float32(&pk.Unknown1)
		io.Float32(&pk.Unknown2)
		io.Bool(&pk.Unknown4)
	}
}
