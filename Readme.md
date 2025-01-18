DNS Query

Goal: Create DNS Server with ai capability

example: 

dns Query: `dig @localhost -p 3000 +short TXT what.is.dns` // Any question

output: `The Domain Name System translates domain names to IP addresses.`

Inspire from:
Github: `https://github.com/knadh/dns.toys`

My Goal: Instead of manually calculating particular query, add LLM is bw to answer that query.

Pros: Can answer almost any query.

Cons: It will be very slow.