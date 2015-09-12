package glitch

import (
	"time"

	"github.com/miekg/dns"
)

type Delay struct {
	Duration time.Duration
}

func (g Delay) Do(request, response *dns.Msg) {
	time.Sleep(g.Duration)
}
