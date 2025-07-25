package packet

import (
	"github.com/Happy2018new/the-last-problem-of-the-humankind/core/minecraft/nbt"
	"github.com/Happy2018new/the-last-problem-of-the-humankind/core/minecraft/protocol"
)

// BlockActorData is sent by the server to update data of a block entity client-side, for example the data of
// a chest.
type BlockActorData struct {
	// Position is the position of the block that holds the block entity. If no block entity is at this
	// position, the packet is ignored by the client.
	Position protocol.BlockPos
	// NBTData is the new data of the block that will be encoded to NBT and applied client-side, so that the
	// client can see the block update. The NBTData should contain all properties of the block, not just
	// properties that were changed.
	NBTData map[string]any
}

// ID ...
func (*BlockActorData) ID() uint32 {
	return IDBlockActorData
}

func (pk *BlockActorData) Marshal(io protocol.IO) {
	io.UBlockPos(&pk.Position)
	io.NBT(&pk.NBTData, nbt.NetworkLittleEndian)
}
