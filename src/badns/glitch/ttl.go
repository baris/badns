package glitch

import (
	"github.com/miekg/dns"
)

type Ttl struct {
	TTL int
}

func (g Ttl) Do(request, response *dns.Msg) {
	for _, answer := range response.Answer {
		header := answer.Header()
		header.Ttl = uint32(g.TTL)
	}
}
