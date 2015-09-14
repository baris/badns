package glitch

import (
	"github.com/miekg/dns"
)

type LimitAnswers struct {
	Limit int
}

func (g LimitAnswers) Do(request, response *dns.Msg) {
	answers := make([]dns.RR, 0)
	for i, answer := range response.Answer {
		if i >= g.Limit {
			break
		}
		answers = append(answers, dns.Copy(answer))
	}
	response.Answer = answers
}
