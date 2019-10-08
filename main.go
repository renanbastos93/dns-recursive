package main

import (
	"fmt"

	"github.com/domainr/dnsr"
)

type answer struct {
	Question string
	Type     string
	Value    string
}

type iDNSRecursive struct {
	ok      bool
	answers []answer
}

var iDNSr iDNSRecursive
var iDNSAnswer answer

func iDNSrecursive(question string, qtype string) {
	r := dnsr.New(1000)
	resolved := r.Resolve(question, qtype)
	for _, rr := range resolved {
		if rr.Type == "A" || rr.Type == "AAAA" {
			iDNSr.ok = true
			iDNSAnswer.Question = question
			iDNSAnswer.Type = rr.Type
			iDNSAnswer.Value = rr.Value
			iDNSr.answers = append(iDNSr.answers, iDNSAnswer)
			break
		}
	}
}

func main() {
	iDNSrecursive("www.renanbastos.com.br", "A")
	iDNSrecursive("www.google.com", "A")
	fmt.Println(iDNSr) // {true [{www.renanbastos.com.br A 191.6.198.85} {www.google.com A 172.217.29.164}]}
}
