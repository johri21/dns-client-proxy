package main

import (
	"dns"
	"log"
	"time"
)

type server struct {
	Address string
	Net     string
}

const CLOUDFLARE_ADDR = "1.1.1.1:853"

func (srv *server) startServer() {

	go func() {
		server := &dns.Server{Addr: srv.Address, Net: srv.Net, Handler: dns.HandlerFunc(dnsHandler)}
		log.Printf("Starting the server Net: %s on address %s", srv.Net, srv.Address)
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to start %s listener %s\n", srv.Net, err.Error())
		}
	}()
}


func dnsHandler(w dns.ResponseWriter, m *dns.Msg) {
	c := dns.Client{Net: "tcp-tls", Timeout: time.Second * 5, DialTimeout: time.Second * 5}
	r, t, err := c.Exchange(m, CLOUDFLARE_ADDR)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("Resolving %s took %v time ", m.Question[0].Name, t)
	w.WriteMsg(r)
}
