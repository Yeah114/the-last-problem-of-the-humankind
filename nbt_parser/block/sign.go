package nbt_parser_block

import (
	"bytes"
	"fmt"

	"github.com/Happy2018new/the-last-problem-of-the-humankind/core/minecraft/protocol"
	"github.com/Happy2018new/the-last-problem-of-the-humankind/mapping"
	"github.com/Happy2018new/the-last-problem-of-the-humankind/utils"
	"github.com/mitchellh/mapstructure"
)

// SignText ..
type SignText struct {
	IgnoreLighting byte   `mapstructure:"IgnoreLighting"`
	SignTextColor  int32  `mapstructure:"SignTextColor"`
	Text           string `mapstructure:"Text"`
}

// SignNBT ..
type SignNBT struct {
	IsWaxed   byte     `mapstructure:"IsWaxed"`
	FrontText SignText `mapstructure:"FrontText"`
	BackText  SignText `mapstructure:"BackText"`
}

// 告示牌
type Sign struct {
	DefaultBlock
	NBT SignNBT
}

func (s *Sign) NeedSpecialHandle() bool {
	if s.NBT.IsWaxed == 1 {
		return true
	}

	texts := []SignText{s.NBT.FrontText, s.NBT.BackText}
	for _, value := range texts {
		if len(value.Text) > 0 {
			return true
		}
	}

	return false
}

func (s Sign) NeedCheckCompletely() bool {
	return true
}

func (s *Sign) Parse(nbtMap map[string]any) error {
	var result SignNBT
	var legacy SignText

	if _, ok := nbtMap["IsWaxed"]; ok {
		err := mapstructure.Decode(&nbtMap, &result)
		if err != nil {
			return fmt.Errorf("Parse: %v", err)
		}
		s.NBT = result
	} else {
		err := mapstructure.Decode(&nbtMap, &legacy)
		if err != nil {
			return fmt.Errorf("Parse: %v", err)
		}
		s.NBT.FrontText = legacy
		s.NBT.BackText = SignText{
			IgnoreLighting: 0,
			SignTextColor:  utils.EncodeVarRGBA(0, 0, 0, 255),
			Text:           "",
		}
	}

	rgb, _ := utils.DecodeVarRGBA(s.NBT.FrontText.SignTextColor)
	bestColor := utils.SearchForBestColor(rgb, mapping.DefaultDyeColor)
	s.NBT.FrontText.SignTextColor = utils.EncodeVarRGBA(bestColor[0], bestColor[1], bestColor[2], 255)

	rgb, _ = utils.DecodeVarRGBA(s.NBT.BackText.SignTextColor)
	bestColor = utils.SearchForBestColor(rgb, mapping.DefaultDyeColor)
	s.NBT.BackText.SignTextColor = utils.EncodeVarRGBA(bestColor[0], bestColor[1], bestColor[2], 255)

	if len(s.NBT.FrontText.Text) == 0 {
		s.NBT.FrontText.IgnoreLighting = 0
		s.NBT.FrontText.SignTextColor = utils.EncodeVarRGBA(0, 0, 0, 255)
	}
	if len(s.NBT.BackText.Text) == 0 {
		s.NBT.BackText.IgnoreLighting = 0
		s.NBT.BackText.SignTextColor = utils.EncodeVarRGBA(0, 0, 0, 255)
	}

	return nil
}

func (s Sign) NBTStableBytes() []byte {
	buf := bytes.NewBuffer(nil)
	w := protocol.NewWriter(buf, 0)

	texts := []SignText{s.NBT.FrontText, s.NBT.BackText}
	for _, value := range texts {
		w.String(&value.Text)
		w.Int32(&value.SignTextColor)
		w.Uint8(&value.IgnoreLighting)
	}
	w.Uint8(&s.NBT.IsWaxed)

	return buf.Bytes()
}

func (s *Sign) FullStableBytes() []byte {
	return append(s.DefaultBlock.FullStableBytes(), s.NBTStableBytes()...)
}
