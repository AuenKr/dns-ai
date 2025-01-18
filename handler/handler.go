package handler

import (
	llm "dns-server/utils"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/miekg/dns"
)

var count uint64 = 0
var mutex *sync.Mutex = &sync.Mutex{}

type DNSHandler struct{}

func (h *DNSHandler) ServeDNS(w dns.ResponseWriter, r *dns.Msg) {
	mutex.Lock()
	count++
	mutex.Unlock()

	msg := new(dns.Msg)
	msg.SetReply(r)

	for _, q := range r.Question {
		// Only processing TXT record
		if q.Qtype != dns.TypeTXT {
			msg.Rcode = dns.RcodeNotImplemented
			msg.Answer = append(msg.Answer, &dns.TXT{
				Hdr: dns.RR_Header{Name: q.Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 60},
				Txt: []string{
					"Only support TXT query",
				},
			})
			w.WriteMsg(msg)
			return
		}

		if q.Name == "stats." {
			msg.Answer = append(msg.Answer, &dns.TXT{
				Hdr: dns.RR_Header{Name: q.Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 0},
				Txt: []string{
					"no req " + strconv.Itoa(int(count)),
				},
			})
			w.WriteMsg(msg)
		}

		var allQuestion []string
		// Formatting question
		for _, each := range r.Question {
			var formatQuestion string
			for _, parts := range strings.Split(each.Name, ".") {
				formatQuestion += parts + " "
			}
			allQuestion = append(allQuestion, formatQuestion)
		}

		// Getting response of first question only
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

		msg.Answer = append(msg.Answer, &dns.TXT{
			Hdr: dns.RR_Header{Name: q.Name, Rrtype: dns.TypeTXT, Class: dns.ClassINET, Ttl: 5 * 60},
			Txt: []string{
				result,
			},
		})

		w.WriteMsg(msg)
	}
}
