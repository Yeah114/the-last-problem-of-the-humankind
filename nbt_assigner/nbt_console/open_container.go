package nbt_console

import (
	"fmt"

	"github.com/Happy2018new/the-last-problem-of-the-humankind/core/minecraft/protocol"
	"github.com/Happy2018new/the-last-problem-of-the-humankind/game_control/game_interface"
	"github.com/Happy2018new/the-last-problem-of-the-humankind/nbt_assigner/block_helper"
)

// OpenContainerByIndex 打开 index 所指示的操作台方块。
// 被打开的目标方块必须是容器、铁砧或织布机。
// index 可用于 BlockByIndex 或 BlockPosByIndex
func (c *Console) OpenContainerByIndex(index int) (success bool, err error) {
	var container block_helper.ContainerBlockHelper
	var isContainer bool
	api := c.api

	block := c.BlockByIndex(index)
	_, ok := (*block).(block_helper.AnvilBlockHelper)
	if !ok {
		_, ok = (*block).(block_helper.LoomBlockHelper)
	}
	if !ok {
		container, isContainer = (*block).(block_helper.ContainerBlockHelper)
		if !isContainer {
			return false, fmt.Errorf("OpenContainerByIndex: Block %T is not a container; *block = %#v", *block, *block)
		}
	}

	if isContainer {
		offset, shouldClean := container.ShouldCleanNearBlock()
		if shouldClean {
			nearBlock := c.NearBlockByIndex(index, offset)
			if _, ok := (*nearBlock).(block_helper.Air); !ok {
				err := api.SetBlock().SetBlock(
					c.NearBlockPosByIndex(index, offset),
					"minecraft:air", "[]",
				)
				if err != nil {
					return false, fmt.Errorf("OpenContainerByIndex: %v", err)
				}
				*c.NearBlockByIndex(index, offset) = block_helper.Air{}
			}
		}
	}

	blockPos := c.BlockPosByIndex(index)
	if err = c.CanReachOrMove(blockPos); err != nil {
		return false, fmt.Errorf("OpenContainerByIndex: %v", err)
	}

	request := game_interface.UseItemOnBlocks{
		HotbarSlotID: c.currentHotBar,
		BlockPos:     blockPos,
		BlockName:    (*block).BlockName(),
		BlockStates:  (*block).BlockStates(),
	}
	success, err = api.ContainerOpenAndClose().OpenContainer(request, false)
	if err != nil {
		return false, fmt.Errorf("OpenContainerByIndex: %v", err)
	}

	return success, nil
}

// OpenContainerByOffset 打开 offset 所指示的操作台方块。
// 被打开的目标方块必须是容器。应当说明的是，铁砧也是容器。
// offset 可用于 BlockByOffset 或 BlockPosByOffset
func (c *Console) OpenContainerByOffset(offset protocol.BlockPos) (success bool, err error) {
	success, err = c.OpenContainerByIndex(helperBlockMappingInv[offset])
	if err != nil {
		return false, fmt.Errorf("OpenContainerByOffset: %v", err)
	}
	return
}
