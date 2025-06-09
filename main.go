package main

import (
	"fmt"
	"math/rand"
)

//package dns_block

type header struct {
	id      uint16
	qr      bool
	opcode  uint16 // 4-bit
	aa      bool
	tc      bool
	rd      bool
	ra      bool
	z       uint16 // 3-bit
	rcode   uint16 // 4-bit
	qdcount uint16
	ancount uint16
	nscount uint16
	arcount uint16
}

type question struct {
	qname  string
	qtype  uint16
	qclass uint16
}

type resource struct {
	name     uint64
	typ      uint16
	class    uint16
	ttl      uint32
	rdlength uint16
	rdata    uint32
}

type message struct {
	head header
	ques question
	ans  resource
	auth resource
	add  resource
}

func generateHeader(id uint16) header {
	return header{
		id:      id,
		qr:      false,
		opcode:  0,
		aa:      false,
		tc:      false,
		rd:      true,
		ra:      false,
		z:       0,
		rcode:   0,
		qdcount: 1,
		ancount: ,
		nscount: ,
		arcount:
	}
}

func generateQuestion(domain string) question {
	return question{
		qname: domain,
		qtype: ,
		qclass:
	}
}

func generateMessage(id uint16, domain string) message {
	msg := message{
		head: generateHeader(id),
		ques: generateQuestion(domain),
	}
}

func messageToBytes(m message) []byte {

}

func main() {
	fmt.Print("Hello, world!")

	id := rand.Uint32()
	domain := "dns.google.com"
	msg := generateMessage(uint16(id), domain)
	msg_bytes := messageToBytes(msg)
}
