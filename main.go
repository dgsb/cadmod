package cadmod

import (
	"github.com/caddyserver/caddy/v2"
)

type CadMod struct{}

func (CadMod) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID: "http.handlers.cadmod",
		New: func() caddy.Module {
			return CadMod{}
		},
	}
}

func init() {
	caddy.RegisterModule(CadMod{})
}
