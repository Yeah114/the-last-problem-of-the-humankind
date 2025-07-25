package block_helper

import "github.com/Happy2018new/the-last-problem-of-the-humankind/utils"

// AnvilBlockHelper 描述了一个铁砧
type AnvilBlockHelper struct {
	States map[string]any
}

func (AnvilBlockHelper) KnownBlockStates() bool {
	return true
}

func (AnvilBlockHelper) BlockName() string {
	return "minecraft:anvil"
}

func (a AnvilBlockHelper) BlockStates() map[string]any {
	return a.States
}

func (a AnvilBlockHelper) BlockStatesString() string {
	return utils.MarshalBlockStates(a.States)
}
