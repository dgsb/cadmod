package cadmod

import (
	"fmt"
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
	/*"github.com/davecgh/go-spew/spew"*/)

var (
	_ caddyfile.Unmarshaler = CadMod{}
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

func (c CadMod) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	for d.Next() {
	}
	return nil
}

func (c CadMod) ServeHTTP(wr http.ResponseWriter, req *http.Request, h caddyhttp.Handler) error {
	if req.TLS != nil {
		/*spew.Dump(req.TLS.PeerCertificates)*/
		for _, cert := range req.TLS.PeerCertificates {
			for _, dnsName := range cert.DNSNames {
				fmt.Println("dns name: ", dnsName)
			}
			for _, email := range cert.EmailAddresses {
				fmt.Println("email: ", email)
			}
		}
	}
	req.Header.Add("X-fake-added-header", "I'm fake !!")
	return h.ServeHTTP(wr, req)
}

func init() {
	caddy.RegisterModule(CadMod{})
	httpcaddyfile.RegisterHandlerDirective("cadmod", func(confhdl httpcaddyfile.Helper) (
		caddyhttp.MiddlewareHandler, error,
	) {
		c := CadMod{}
		c.UnmarshalCaddyfile(confhdl.Dispenser)
		return c, nil
	})
}
