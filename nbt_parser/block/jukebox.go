package nbt_parser_block

import (
	"bytes"
	"fmt"

	"github.com/Happy2018new/the-last-problem-of-the-humankind/core/minecraft/protocol"
	nbt_parser_interface "github.com/Happy2018new/the-last-problem-of-the-humankind/nbt_parser/interface"
)

// JukeBoxNBT ..
type JukeBoxNBT struct {
	CustomName string
	HaveDisc   bool
	Disc       nbt_parser_interface.Item
}

// 唱片机
type JukeBox struct {
	DefaultBlock
	NBT JukeBoxNBT
}

func (j JukeBox) NeedSpecialHandle() bool {
	if len(j.NBT.CustomName) > 0 {
		return true
	}
	if j.NBT.HaveDisc {
		return true
	}
	return false
}

func (JukeBox) NeedCheckCompletely() bool {
	return true
}

func (j *JukeBox) Parse(nbtMap map[string]any) error {
	j.NBT.CustomName, _ = nbtMap["CustomName"].(string)
	discMap, ok := nbtMap["RecordItem"].(map[string]any)
	if ok {
		disc, err := nbt_parser_interface.ParseItemNormal(discMap)
		if err != nil {
			return fmt.Errorf("Parse: %v", err)
		}
		j.NBT.HaveDisc = true
		j.NBT.Disc = disc
	}
	return nil
}

func (j JukeBox) NBTStableBytes() []byte {
	buf := bytes.NewBuffer(nil)
	w := protocol.NewWriter(buf, 0)

	w.String(&j.NBT.CustomName)
	w.Bool(&j.NBT.HaveDisc)
	if j.NBT.HaveDisc {
		bookStableBytes := j.NBT.Disc.TypeStableBytes()
		w.ByteSlice(&bookStableBytes)
	}

	return buf.Bytes()
}

func (j *JukeBox) FullStableBytes() []byte {
	return append(j.DefaultBlock.FullStableBytes(), j.NBTStableBytes()...)
}
