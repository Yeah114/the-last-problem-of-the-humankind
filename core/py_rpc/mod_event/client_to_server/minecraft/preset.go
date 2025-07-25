package minecraft

import (
	"github.com/Happy2018new/the-last-problem-of-the-humankind/core/py_rpc/mod_event/client_to_server/minecraft/preset"
	mei "github.com/Happy2018new/the-last-problem-of-the-humankind/core/py_rpc/mod_event/interface"
)

type Preset struct{ mei.Module }

// Return the module name of p
func (p *Preset) ModuleName() string {
	return "preset"
}

// Return a pool/map that contains all the event of p
func (p *Preset) EventPool() map[string]mei.Event {
	return map[string]mei.Event{
		"GetLoadedInstances": &preset.GetLoadedInstances{},
	}
}
