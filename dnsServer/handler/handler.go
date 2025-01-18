package handler

import (
	llm "dns-server/utils"
	"fmt"
	"strings"

	"github.com/miekg/dns"
)

type DNSHandler struct{}

func (h *DNSHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	msg := new(dns.Msg)
	msg.SetReply(r)

	for _, q := range r.Question {
		// Only processing TXT record
		if q.Qtype != dns.TypeTXT {
			msg.Rcode = dns.RcodeNotImplemented
		} else {
			var allQuestion []string
			// Formatting question
			for _, each := range r.Question {
				var formatQuestion string
				for _, parts := range strings.Split(each.Name, ".") {
					formatQuestion += parts + " "
				}
				allQuestion = append(allQuestion, formatQuestion)
			}

			fmt.Println("All question ", allQuestion)

			// Getting response
			result, err := llm.GenerateContent(allQuestion[0])

			if err != nil {
				fmt.Println("error occur ", err)
				msg.Rcode = dns.RcodeServerFailure
				msg.Answer = append(msg.Answer, &dns.TXT{
					Hdr: dns.RR_Header{Name: q.Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 60},
					Txt: []string{
						"Fail to generate response from LLM",
					},
				})
				w.WriteMsg(msg)
				return
			}

			fmt.Print("result ", result)

			msg.Answer = append(msg.Answer, &dns.TXT{
				Hdr: dns.RR_Header{Name: q.Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 60},
				Txt: []string{
					result,
				},
			})
		}
	}

	w.WriteMsg(msg)
}
