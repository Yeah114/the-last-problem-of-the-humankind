package minecraft

import (
	"github.com/Happy2018new/the-last-problem-of-the-humankind/core/py_rpc/mod_event/client_to_server/minecraft/ai_command"
	mei "github.com/Happy2018new/the-last-problem-of-the-humankind/core/py_rpc/mod_event/interface"
)

// 魔法指令
type AICommand struct{ mei.Module }

// Return the module name of a
func (a *AICommand) ModuleName() string {
	return "aiCommand"
}

// Return a pool/map that contains all the event of a
func (a *AICommand) EventPool() map[string]mei.Event {
	return map[string]mei.Event{
		"ExecuteCommandEvent": &ai_command.ExecuteCommandEvent{},
	}
}
