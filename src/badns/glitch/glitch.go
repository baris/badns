package glitch

import (
	"github.com/miekg/dns"
)

type Glitch interface {
	Do(request, response *dns.Msg)
}
