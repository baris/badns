package glitch

import (
	"github.com/miekg/dns"
)

type NoAnswer struct {
}

func (g NoAnswer) Do(request, response *dns.Msg) {
	response.Answer = make([]dns.RR, 0)
}
