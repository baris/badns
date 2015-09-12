package glitch

import (
	"strings"

	"github.com/miekg/dns"
)

type ReplaceType struct {
	Type string
}

func (g ReplaceType) Do(request, response *dns.Msg) {
	from_type := "*"
	to_type := ""
	if strings.Index(g.Type, ":") >= 0 {
		fields := strings.Split(g.Type, ":")
		from_type, to_type = fields[0], fields[1]
	} else {
		to_type = g.Type
	}

	for i, answer := range response.Answer {
		answer = dns.Copy(answer)
		header := answer.Header()
		if from_type == "*" || header.Rrtype == dns.StringToType[from_type] {
			header.Rrtype = dns.StringToType[to_type]
		}
		response.Answer[i] = answer
	}
}
