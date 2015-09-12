package main

import (
	"badns/glitch"

	"strings"

	"github.com/miekg/dns"
)

const (
	ORIGIN string = "$ORIGIN badns.badns."
	SOA    string = "@ SOA badns.badns. admin.badns.badns 999988776600 1800 900 0604800 604800"
	SERVER string = "8.8.8.8:53"
)

func query(r *dns.Msg) *dns.Msg {
	c := new(dns.Client)
	m, _, _ := c.Exchange(r, SERVER)
	return m
}

type BaDNSHandler struct {
	glitches   []glitch.Glitch
	hostPrefix string
}

func (badns *BaDNSHandler) AddGlitch(glitch glitch.Glitch) {
	badns.glitches = append(badns.glitches, glitch)
}

func (badns BaDNSHandler) ServeDNS(w dns.ResponseWriter, request *dns.Msg) {
	log.Debug("Handling DNS query for ", request.Question)

	response := query(request)

	if hostPrefixMatches(request, badns.hostPrefix) {
		for _, glitch := range badns.glitches {
			glitch.Do(request, response)
		}
	}

	w.WriteMsg(response)
}

func hostPrefixMatches(request *dns.Msg, prefix string) bool {
	if prefix == "" {
		return true
	}
	for _, q := range request.Question {
		if strings.HasPrefix(q.Name, prefix) {
			return true
		}
	}
	return false
}
