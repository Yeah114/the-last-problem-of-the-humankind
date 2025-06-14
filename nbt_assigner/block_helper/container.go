package block_helper

import (
	"github.com/Happy2018new/the-last-problem-of-the-humankind/utils"
)

// ContainerBlockOpenInfo 描述了要
// 打开一个容器所必须知道的最少的信息
type ContainerBlockOpenInfo struct {
	// 这个容器的名称
	Name string
	// 这个容器的方块状态
	States map[string]any
	// ConsiderOpenDirection 指示打开目标容器
	// 是否需要考虑其打开方向上的障碍物方块，
	// 这似乎只对箱子和潜影盒有效
	ConsiderOpenDirection bool
	// 当 ConsiderOpenDirection 为真时，
	// 应在 Facing 填写该容器的朝向，
	// 否则可以置为默认的零值
	Facing uint8
}

// ContainerBlockHelper 描述了一个容器，
// 并记载了它应当如何被打开
type ContainerBlockHelper struct {
	// OpenInfo 提供的信息足以使该容器能被打开
	OpenInfo ContainerBlockOpenInfo
	// IsEmpty 指示该容器是否已经全空
	IsEmpty bool
}

func (c ContainerBlockHelper) BlockName() string {
	return c.OpenInfo.Name
}

func (c ContainerBlockHelper) BlockStates() map[string]any {
	return c.OpenInfo.States
}

func (c ContainerBlockHelper) BlockStatesString() string {
	return utils.MarshalBlockStates(c.OpenInfo.States)
}

// ShouldCleanNearBlock 指示打开该容器前是否需要清除
// 其相邻的方块。offset 指示这个相邻方块的位置。这目前
// 只对箱子和潜影盒有用
func (c ContainerBlockHelper) ShouldCleanNearBlock() (offset [3]int32, needClean bool) {
	if !c.OpenInfo.ConsiderOpenDirection {
		return [3]int32{}, false
	}

	switch c.OpenInfo.Facing {
	case 0:
		return [3]int32{0, -1, 0}, true
	case 1:
		return [3]int32{0, 1, 0}, true
	case 2:
		return [3]int32{0, 0, -1}, true
	case 3:
		return [3]int32{0, 0, 1}, true
	case 4:
		return [3]int32{-1, 0, 0}, true
	case 5:
		return [3]int32{1, 0, 0}, true
	}

	return
}
