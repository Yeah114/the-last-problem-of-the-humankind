package nbt_parser_item

import (
	"bytes"
	"fmt"

	"github.com/Happy2018new/the-last-problem-of-the-humankind/core/minecraft/protocol"
	"github.com/Happy2018new/the-last-problem-of-the-humankind/mapping"
	nbt_parser_general "github.com/Happy2018new/the-last-problem-of-the-humankind/nbt_parser/general"
	"github.com/mitchellh/mapstructure"
)

// ShieldNBT ..
type ShieldNBT struct {
	HaveBase bool
	Base     int32
	Patterns []nbt_parser_general.BannerPattern
}

// 盾牌
type Shield struct {
	DefaultItem
	NBT ShieldNBT
}

// parse ..
func (s *Shield) parse(tag map[string]any) error {
	var isOminousShield bool

	s.DefaultItem.Enhance.ItemComponent.LockInInventory = false
	s.DefaultItem.Enhance.ItemComponent.LockInSlot = false
	s.DefaultItem.Block = ItemBlockData{}

	if len(tag) == 0 {
		return nil
	}

	patterns, _ := tag["Patterns"].([]any)
	for _, value := range patterns {
		var pattern nbt_parser_general.BannerPattern

		val, ok := value.(map[string]any)
		if !ok {
			continue
		}

		err := mapstructure.Decode(&val, &pattern)
		if err != nil {
			return fmt.Errorf("parse: %v", err)
		}

		if mapping.BannerPatternUnsupported[pattern.Pattern] {
			continue
		}
		if pattern.Pattern == mapping.BannerPatternOminous {
			isOminousShield = true
			s.NBT.Patterns = []nbt_parser_general.BannerPattern{
				pattern,
			}
			break
		}

		s.NBT.Patterns = append(s.NBT.Patterns, pattern)
	}

	if isOminousShield {
		s.Basic.Metadata = 0
		s.Enhance.ItemComponent.KeepOnDeath = false
	}

	s.NBT.Base, s.NBT.HaveBase = tag["Base"].(int32)
	return nil
}

func (s *Shield) ParseNormal(nbtMap map[string]any) error {
	tag, _ := nbtMap["tag"].(map[string]any)
	err := s.parse(tag)
	if err != nil {
		return fmt.Errorf("ParseNormal: %v", err)
	}
	return nil
}

func (s *Shield) ParseNetwork(item protocol.ItemStack, itemName string) error {
	err := s.parse(item.NBTData)
	if err != nil {
		return fmt.Errorf("ParseNetwork: %v", err)
	}
	return nil
}

func (s Shield) IsComplex() bool {
	if len(s.NBT.Patterns) > 0 {
		return true
	}
	if s.NBT.HaveBase {
		return true
	}
	return false
}

func (s Shield) complexFieldsOnly() []byte {
	buf := bytes.NewBuffer(nil)
	w := protocol.NewWriter(buf, 0)

	protocol.SliceUint16Length(w, &s.NBT.Patterns)
	w.Bool(&s.NBT.HaveBase)
	w.Varint32(&s.NBT.Base)

	return buf.Bytes()
}

func (s Shield) NBTStableBytes() []byte {
	return append(s.DefaultItem.NBTStableBytes(), s.complexFieldsOnly()...)
}

func (s *Shield) TypeStableBytes() []byte {
	return append(s.DefaultItem.TypeStableBytes(), s.complexFieldsOnly()...)
}

func (s *Shield) FullStableBytes() []byte {
	return append(s.TypeStableBytes(), s.Basic.Count)
}
