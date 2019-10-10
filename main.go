package main

import (
    "github.com/miekg/dns"
    "net"
    "log"
    "fmt"
    "time"
)

type Resolver struct {
    ok bool
	timeout time.Duration
}


func recursiveDNS(question string, qType uint16) {

    config, _ := dns.ClientConfigFromFile("/etc/resolv.conf")
    c := new(dns.Client)

    m := new(dns.Msg)
    m.SetQuestion(dns.Fqdn(question), dns.TypeA)
    m.RecursionDesired = true

    r, _, err := c.Exchange(m, net.JoinHostPort(config.Servers[0], config.Port))

    if r == nil {
        log.Fatalf("*** error: %s\n", err.Error())
    }

    if dns.RcodeNameError == r.Rcode {
        log.Fatalf(" *** error: NXDOMAIN")
    }

    fmt.Println(r)
}


func main() {
    
    // apex := "www.renanbastos.com.br"
    apex := "ssh.sbh.eng.br"
    // apex := "sbh.eng.br"
    // apex := "sdasda.op" 

    recursiveDNS(apex, dns.TypeA)
    

    // if r.Rcode != dns.RcodeSuccess {
    //         log.Fatalf("invalid answer name after %s query for %s\n", r.Rcode,apex)
    // }
    // // Stuff must be in the answer section
    // if len(r.Answer) < 1 {
    //     fmt.Println(r.Answer)
    // }
    // for _, a := range r.Answer {
    //         fmt.Printf("%v\n", a)
    // }
}

